package main

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
