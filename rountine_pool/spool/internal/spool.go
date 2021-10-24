/*
 * 共享协程池
 */

package internal

import (
	"fmt"
	"sync"
	"time"
)

type SPool struct {
	jobQueue chan WorkerInterface            //工作队列
	wg 		sync.WaitGroup
	workerNum int                           
	TotalNum  int
	CounterOk int
	CounterFail int
	CounterOut int
	mutexFail sync.RWMutex
	mutexOk   sync.RWMutex
	mutexOut  sync.RWMutex
	TimeStart int64
	TimeOut   int
	Debug     bool
}

func NewSPool(workerNum int, totalNum int, timeout int, debug bool)*SPool{
	p := &SPool{
		jobQueue: make(chan WorkerInterface, workerNum),
		TotalNum: totalNum,
		workerNum: workerNum,
		TimeOut: timeout,
		Debug: debug,
		TimeStart: time.Now().Unix(),
	}

	p.dispatch()

	return p
}

// Commit 提交任务到任务队列
func (p *SPool) Commit(w WorkerInterface) {
	p.jobQueue <- w
}

//等待组 关闭channel
func (p *SPool) Release() {
	close(p.jobQueue)
	p.wg.Wait()

	if p.Debug{
		p.RuntimeLog()
	}
}

//计数器-失败
func (p *SPool) CountFail() {
	if p.Debug{
		p.mutexFail.Lock()
		p.CounterFail++
		p.mutexFail.Unlock()
	}
}

//计数器-成功
func (p *SPool) CountOk() {
	if p.Debug{
		p.mutexOk.Lock()
		p.CounterOk++
		p.mutexOk.Unlock()
	}
}

// CountTimeOut 计数器-超时
func (p *SPool) CountTimeOut() {
	if p.Debug{
		p.mutexOut.Lock()
		p.CounterOut++
		p.mutexOut.Unlock()
	}
}

// RuntimeLog 实时日志
func (p *SPool) RuntimeLog() {
	if p.Debug {
		ttime := MathDecimal(float64(time.Now().Unix() - p.TimeStart))    //运行时间
		trange := MathDecimal(float64(p.TotalNum) / ttime)
		if p.CounterOk > 0 || p.CounterFail > 0 {
			if p.TimeOut > 0 {
				fmt.Println(fmt.Sprintln("runtime:total|fail|timeout:", p.TotalNum, "|", p.CounterFail, "|", p.CounterOut, "", "消耗时间:(", ttime, "s)", "平均:(", trange, "次/s)"))
			} else {
				fmt.Println(fmt.Sprintln("runtime:total|fail:", p.TotalNum, "|", p.CounterFail, "", "消耗时间:(", ttime, "s)", "平均:(", trange, "次/s)"))
			}
		}
	}
}

//分发任务
func (p *SPool) dispatch() {
	p.wg.Add(p.workerNum)

	for i := 0; i < p.workerNum;i++{    //开多个工作协程处理任务队列
		go func(i int) {
			defer func() {
				p.wg.Done()
				if err := recover(); err != nil{
					fmt.Println("worker error", err)
				}
			}()

			fmt.Println("worker:", i)
			for w := range p.jobQueue{
				if p.TimeOut > 0 {
					timeout_ch := make(chan interface{})

					go func(wr WorkerInterface) { p.runTaskTimeout(wr, timeout_ch) }(w)

					for {
						select {
						case <-timeout_ch:
							goto forend
						case <-time.After(time.Duration(p.TimeOut) * time.Second):
							p.CountTimeOut()
							fmt.Println(w.GetTaskID(), "timeout")
							goto forend
						}
					}
				forend:
				}else{
					p.runTask(w)
				}
			}
		}(i)
	}

	//p.wg.Wait()
}

//执行超时任务
func (p *SPool) runTaskTimeout(wr WorkerInterface, timeoutCh chan interface{}){
	defer func() {
		if err := recover(); err != nil{
			fmt.Println("task error", err)
		}
	}()

	err := wr.Task()

	timeoutCh <- "ok"

	if err != nil{
		p.CountFail()
		panic(err)
	}else{
		p.CountOk()
	}
}

func (p *SPool) runTaskTimeWithFunc(wr job, timeoutCh chan interface{}){
	defer func() {
		if err := recover(); err != nil{
			fmt.Println("task error", err)
		}
	}()

	err := wr()

	timeoutCh <- "ok"

	if err != nil{
		p.CountFail()
		panic(err)
	}else{
		p.CountOk()
	}
}

//执行普通任务
func (p *SPool) runTask(wr WorkerInterface){
	err := wr.Task()

	if err != nil{
		p.CountFail()
		panic(err)
	}else{
		p.CountOk()
	}
}

var _ Pool = (*SPool)(nil)