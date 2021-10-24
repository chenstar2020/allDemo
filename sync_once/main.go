package main

import (
	"fmt"
	"sync"
	"time"
)

func main(){
	once := sync.Once{}
	for i := 0; i < 10; i++{
		go func() {
			once.Do(func() {
				fmt.Println("only print once")
			})
		}()
	}

	time.Sleep(time.Second)
}
