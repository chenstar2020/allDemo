package main

import (
	"testing"
)

//go test -bench="Generate" .
func BenchmarkGenerate(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Generate(1000000)
	}
}

func BenchmarkGenerateWithCap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GenerateWithCap(1000000)
	}
}

func benchmarkGenerate(i int, b *testing.B){
	for n := 0; n < b.N; n++ {
		Generate(i)
	}
}

func BenchmarkGenerate1000(b *testing.B) {
	benchmarkGenerate(1000, b)
}

func BenchmarkGenerate10000(b *testing.B) {
	benchmarkGenerate(10000, b)
}

func BenchmarkGenerate100000(b *testing.B) {
	benchmarkGenerate(100000, b)
}

func BenchmarkBubbleSort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()            //停止计时
		nums := GenerateWithCap(10000)
		b.StartTimer()           //开始计时
		BubbleSort(nums)
	}
}