package main

func bubbleSort(arr []int) {
	length := len(arr)
	if length < 2 {
		return
	}
	for i := 0; i < length; i++ {
		for j := length - 1; j >= i+1; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
}
