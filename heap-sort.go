package main

func maxHeapify(arr []int, pos int) {
	length := len(arr)
	largest := pos
	l := 2*pos + 1
	r := l + 1
	if l < length && arr[l] > arr[pos] {
		largest = l
	}
	if r < length && arr[r] > arr[largest] {
		largest = r
	}

	if pos != largest {
		arr[pos], arr[largest] = arr[largest], arr[pos]
		maxHeapify(arr, largest)
	}
}

func buildMaxHeap(arr []int) {
	for i := (len(arr) - 1) / 2; i >= 0; i-- {
		maxHeapify(arr, i)
	}
}

func heapSort(arr []int) {
	buildMaxHeap(arr)
	for i := len(arr) - 1; i >= 1; i-- {
		arr[i], arr[0] = arr[0], arr[i]
		maxHeapify(arr[:i-1], 0)
	}
}
