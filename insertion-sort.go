package main

import (
	"fmt"
)

//  插入排序
func insertionSort(arr []int, asc bool) {
	length := len(arr)
	if length < 2 {
		return
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
}

func main() {
	arr := []int{31, 41, 59, 26, 41, 58}
	insertionSort(arr, true)
	fmt.Println(arr)
	insertionSort(arr, false)
	fmt.Println(arr)
}
