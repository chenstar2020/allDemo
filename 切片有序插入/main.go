package main

import (
	"fmt"
)

func main(){
	func(){

	}()

	arr := make([]int, 0)
	arr = sortedAppend(arr)
	arr = sortedAppend(arr, 12, 8, 7)
	arr = sortedAppend(arr, 1, 20, 10)
	fmt.Println(arr)
}

func sortedInsert(arr []int, ele int)[]int{
	index := 0
	for _, v := range arr{
		if ele >= v {
			index++
		}else{
			break
		}
	}

	newArr := make([]int, 0)
	newArr = append(newArr, arr[:index]...)
	newArr = append(newArr, ele)
	newArr = append(newArr, arr[index:]...)

	return newArr
}

func sortedAppend(arr []int, ele... int)[]int{
	newArr := make([]int, len(arr))
	copy(newArr, arr)

	for i := 0; i < len(ele);i++{
		newArr = sortedInsert(newArr, ele[i])
	}
	return newArr
}
