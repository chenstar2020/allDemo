package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func generate(message string, interval time.Duration)(chan string, chan struct{}){
	mc := make(chan string)
	sc := make(chan struct{})   //用于控制协程退出

	go func() {
		defer func() {
			close(sc)
		}()

		for{
			select{
			case <-sc:
				return
			default:
				time.Sleep(interval)
				mc <- message
			}
		}
	}()

	return mc, sc
}

func stopGenerating(mc chan string, sc chan struct{}){
	sc <- struct{}{}
	close(mc)
}

//多路复用函数
//创建并返回整合消息channel和控制并发的wg
func multiplex(mcs ...chan string)(chan string, *sync.WaitGroup){
	mmc := make(chan string)
	wg := &sync.WaitGroup{}

	for _, mc := range mcs{
		wg.Add(1)

		go func(mc chan string, wg *sync.WaitGroup) {
			defer wg.Done()

			for m := range mc{
				mmc <- m
			}
		}(mc, wg)
	}

	return mmc, wg
}


func main(){
	//循环写入数据
	mc1, sc1 := generate("message from generator 1", 200 *time.Millisecond)
	mc2, sc2 := generate("message from generator 2", 300 * time.Millisecond)

	mmc, wg1 := multiplex(mc1, mc2)

	errs := make(chan error)

	go func() {
		sc := make(chan os.Signal, 1)
		signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s signal received",  <-sc)
	}()

	wg2 := &sync.WaitGroup{}
	wg2.Add(1)
	go func(){
		defer wg2.Done()

		for m := range mmc{
			fmt.Println("打印消息：", m)
		}
	}()

	if err := <-errs; err != nil{
		fmt.Println("退出消息：", err.Error())
	}

	stopGenerating(mc1, sc1)
	stopGenerating(mc2, sc2)
	wg1.Wait()

	close(mmc)
	wg2.Wait()

}
