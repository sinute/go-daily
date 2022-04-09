package main

import (
	"fmt"
	"testing"
)

func Test_heapSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
	}{
		{args: args{arr: []int{31, 41, 59, 26, 41, 58}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			heapSort(tt.args.arr)
			fmt.Println(tt.args.arr)
		})
	}
}
