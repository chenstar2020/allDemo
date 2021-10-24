package internal

import (
	"fmt"
	"strconv"
)

// Pool 协程池接口
type Pool interface {
	Commit(workerInterface WorkerInterface)
	Release()
	CountFail()			  //失败计数
	CountOk()             //成功计数
	CountTimeOut()        //超时计数
	RuntimeLog()		  //实时日志
	dispatch()            //分发任务到工作协程
}

type job func() error

// WorkerInterface 任务接口
type WorkerInterface interface {
	Task() error
	GetTaskID() interface{}
}

func MathDecimal(value float64) float64{
	value, _ = strconv.ParseFloat(fmt.Sprintf("%0.2f", value), 64)
	return value
}