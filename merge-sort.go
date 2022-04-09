package main

func merge(l []int, r []int) (arr []int) {
	lLength := len(l)
	rLength := len(r)
	arr = make([]int, 0, lLength+rLength)
	i, j := 0, 0
	for i < lLength && j < rLength {
		if l[i] < r[j] {
			arr = append(arr, l[i])
			i++
		} else {
			arr = append(arr, r[j])
			j++
		}
	}
	if i == lLength {
		arr = append(arr, r[j:]...)
	} else {
		arr = append(arr, l[i:]...)
	}
	return
}

// 合并排序
func mergeSort(arr []int, start, end int) []int {
	length := len(arr)
	if start > end || start < 0 || end > length || length == 1 {
		return arr
	}
	l := arr[start : (start+end)/2]
	r := arr[(start+end)/2 : end]
	return merge(mergeSort(l, 0, len(l)), mergeSort(r, 0, len(r)))
}
