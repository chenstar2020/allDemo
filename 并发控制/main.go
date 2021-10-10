package main

import (
	"context"
	"fmt"
	"time"
)

func main(){
	ctx := context.Background()     	//返回一个context接口的简单实现
	before := time.Now()
	preCtx, _ := context.WithTimeout(ctx, time.Millisecond * 500)  //派生出子context
	go func() {
		childCtx, _ := context.WithTimeout(preCtx, 300 * time.Millisecond)
		select {
		case <- childCtx.Done():        //子context退出的三种时机：1.父context退出 2.超时退出 3.主动调用cancel函数退出
			after := time.Now()
			fmt.Println("child during:", after.Sub(before).Milliseconds())
		}
	}()

	select {
	case <- preCtx.Done():
		after := time.Now()
		fmt.Println("pre after:", after.Sub(before).Milliseconds())
	}
}

type Value struct{
	data int
}

func Stream(ctx context.Context, out chan <- Value) error{
	v, err := DoSomething(ctx)
	if err != nil{
		return err
	}
	select {
	case <-ctx.Done():       //协程主动退出  Done函数返回一个读取通道
		return ctx.Err()
	case out <- v:			 //正常写入数据退出

	}
	return nil
}

func DoSomething(ctx context.Context)(Value, error){
	return Value{}, nil
}