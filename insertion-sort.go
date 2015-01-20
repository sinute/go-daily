package main

import (
	"fmt"
)

//  插入排序
func insertionSort(arr []int, asc bool) []int {
	length := len(arr)
	if length < 2 {
		return arr
	}
	for i := 1; i < length; i++ {
		x := i
		for j := i - 1; j >= 0; j-- {
			if (arr[x] < arr[j]) == asc {
				arr[x], arr[j] = arr[j], arr[x]
				x -= 1
			} else {
				break
			}
		}

	}
	return arr
}

func main() {
	arr := []int{31, 41, 59, 26, 41, 58}
	fmt.Println(insertionSort(arr, true))
	fmt.Println(insertionSort(arr, false))
}
