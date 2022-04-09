package main

import (
	"fmt"
	"testing"
)

func Test_mergeSort(t *testing.T) {
	type args struct {
		arr   []int
		start int
		end   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{args: args{arr: []int{31, 41, 59, 26, 41, 58}, start: 0, end: 6}},
		{args: args{arr: []int{31, 41, 59, 26, 41, 58}, start: 2, end: 6}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println(mergeSort(tt.args.arr, tt.args.start, tt.args.end))
			// if got := mergeSort(tt.args.arr, tt.args.start, tt.args.end); !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("mergeSort() = %v, want %v", got, tt.want)
			// }
		})
	}
}
