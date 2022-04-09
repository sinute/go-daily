package main

import (
	"reflect"
	"testing"
)

func Test_bubbleSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{args: args{arr: []int{31, 41, 59, 26, 41, 58}}, want: []int{26, 31, 41, 41, 58, 59}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Run(tt.name, func(t *testing.T) {
				if bubbleSort(tt.args.arr); !reflect.DeepEqual(tt.args.arr, tt.want) {
					t.Errorf("bubbleSort() = %v, want %v", tt.args.arr, tt.want)
				}
			})
		})
	}
}
