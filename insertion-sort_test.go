package main

import (
	"fmt"
	"testing"
)

func Test_insertionSort(t *testing.T) {
	type args struct {
		arr []int
		asc bool
	}
	tests := []struct {
		name string
		args args
	}{
		{args: args{arr: []int{31, 41, 59, 26, 41, 58}, asc: true}},
		{args: args{arr: []int{31, 41, 59, 26, 41, 58}, asc: false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			insertionSort(tt.args.arr, tt.args.asc)
			fmt.Println(tt.args.arr)
		})
	}
}
