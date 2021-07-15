package main

import (
	"math/rand"
	"time"
)

func GenerateWithCap(n int)[]int{
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, 0, n)
	for i := 0; i < n; i++ {
		nums = append(nums, rand.Int())
	}
	return nums
}

func Generate(n int)[]int{
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, 0)
	for i := 0; i < n; i++ {
		nums = append(nums, rand.Int())
	}
	return nums
}


func BubbleSort(num []int){
	for i := 0; i < len(num);i++ {
		for j := 1; j < len(num) - i; j++ {
			if num[j - 1] > num[j] {
				num[j], num[j-1] = num[j-1], num[j]
			}
		}
	}
}