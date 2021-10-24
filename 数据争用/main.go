package main

import (
	"fmt"
	"sync/atomic"
)

var flag int64 = 0
var count int64 = 0

func add(){
	for{
		if atomic.CompareAndSwapInt64(&flag, 0, 1){  //判断flag是否为0, 如果为0 置为1
			count++
			atomic.StoreInt64(&flag, 0)      //flag置为0
			return
		}
	}
}


func main(){
	for i := 0; i < 10000; i++{
		go add()
	}

	fmt.Println("count:", count)
}
