package main

//冒泡排序
func bubbleSort(num []int){
	for i := 0; i < len(num);i++ {
		for j := 1; j < len(num) - i; j++ {
			if num[j-1] > num[j] {
				num[j], num[j-1] = num[j-1], num[j]
			}
		}
	}
}
