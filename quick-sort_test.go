package main

import (
	"reflect"
	"testing"
)

func Test_quickSort(t *testing.T) {
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
			if quickSort(tt.args.arr, 0, len(tt.args.arr)-1); !reflect.DeepEqual(tt.args.arr, tt.want) {
				t.Errorf("quickSort() = %v, want %v", tt.args.arr, tt.want)
			}
		})
	}
}
