package main

func partition(arr []int, l, r int) (pos int) {
	x := arr[r]
	pos = l
	for i := l; i < r; i++ {
		if arr[i] <= x {
			arr[pos], arr[i] = arr[i], arr[pos]
			pos++
		}
	}
	arr[pos], arr[r] = arr[r], arr[pos]
	return
}

func quickSort(arr []int, i, j int) {
	if j <= i {
		return
	}
	m := partition(arr, i, j)
	quickSort(arr, i, m-1)
	quickSort(arr, m+1, j)
}
