/*
 * 非共享协程池
 */

package internal

import (
	"fmt"
	"sync"
	"time"
)

type WPool struct {
	jobQueue chan WorkerInterface
	limitChan chan interface{}
	wg sync.WaitGroup
	TotalNum int
	CounterOk int
	CounterFail int
	CounterOut  int
	mutexFail	sync.Mutex
	mutexOk 	sync.Mutex
	mutexOut    sync.Mutex
	TimeStart   int64
	TimeOut     int
	Debug       bool
}

func (p *WPool) Commit(w WorkerInterface) {
	p.limitChan <- "ok"
	p.jobQueue <- w
}

func (p *WPool) Release() {
	p.wg.Wait()
	close(p.jobQueue)
	close(p.limitChan)

	if p.Debug{
		p.RuntimeLog()
	}
}

func (p *WPool) CountFail() {
	if p.Debug{
		p.mutexOk.Lock()
		defer p.mutexOk.Unlock()

		p.CounterFail++
	}
}

func (p *WPool) CountOk() {
	p.mutexOk.Lock()
	defer p.mutexOk.Unlock()

	p.CounterOk++
}

func (p *WPool) CountTimeOut() {
	if p.Debug{
		p.mutexOk.Lock()
		defer p.mutexOk.Unlock()

		p.CounterOut++
	}
}

func (p *WPool) RuntimeLog() {
	if p.Debug{
		ttime := MathDecimal(float64(time.Now().Unix() - p.TimeStart))
		trange := MathDecimal(float64(p.TotalNum) / ttime)
		if p.CounterOk > 0 || p.CounterFail > 0 {
			if p.TimeOut >0 {
				fmt.Println(fmt.Sprintln("runtime:total|fail|timeout:", p.TotalNum, "|",p.CounterFail,"|", p.CounterOut, "", "消耗时间:(", ttime, "s)", "平均:(", trange, "次/s)"))
			}else{
				fmt.Println(fmt.Sprintln("runtime:total|fail:", p.TotalNum, "|", p.CounterFail, "", "消耗时间:(", ttime, "s)", "平均:(", trange, "次/s)"))
			}
		}
	}
}

func (p *WPool) dispatch() {
	p.TimeStart = time.Now().Unix()

	go func() {
		for w := range p.jobQueue{
			go func(wr WorkerInterface) {
				defer func() {
					p.wg.Done()
					<-p.limitChan
					if err := recover(); err != nil{
						fmt.Println("task run err", err)
					}
				}()

				if p.TimeOut > 0 {
					timeout_ch := make(chan interface{})

					go func() {
						p.runTaskTimeout(wr, timeout_ch)
					}()

					for {
						select {
						case <- timeout_ch:
							return
						case <- time.After(time.Duration(p.TimeOut) * time.Second):
							p.CountTimeOut()
							if p.Debug{
								fmt.Println(wr.GetTaskID(), "timeout")
							}
							return
						}
					}
				}else{
					p.runTask(wr)
				}
			}(w)

		}
	}()
}

func (p *WPool) runTaskTimeout(wr WorkerInterface, timeout_ch chan interface{}){
	defer func() {
		if err := recover(); err != nil{
			fmt.Println("task run err", err)
		}
	}()

	err := wr.Task()

	timeout_ch <- "ok"

	if err == nil{
		p.CountOk()
	}else{
		p.CountFail()
		panic(err)
	}
}

func (p *WPool) runTask(wr WorkerInterface){
	defer func() {
		if err := recover(); err != nil{
			fmt.Println("task run err", err)
		}
	}()

	err := wr.Task()

	if err == nil{
		p.CountOk()
	}else{
		p.CountFail()
		panic(err)
	}
}

func NewWPool(workerNum int, totalNum int, timeout int, debug bool)*WPool{
	p := WPool{
		TotalNum: totalNum,
		jobQueue: make(chan WorkerInterface),
		limitChan: make(chan interface{}, workerNum),
		TimeOut: timeout,
		Debug: debug,
	}
	p.wg.Add(totalNum)
	p.dispatch()
	return &p
}


var _ Pool = (*WPool)(nil)