package main

import (
	"testing"
	"time"
)

//测试命令: go test -bench="Fib$" -benchtime=50x .
func BenchmarkFib(b *testing.B) {
	time.Sleep(time.Second * 3)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		Fib(30)
	}
}
