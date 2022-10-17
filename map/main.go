package main

import (
	"fmt"
)

func main() {
	v := map[int][]int{
		1: {11, 12, 13, 14, 15},
		2: {21, 22, 23, 24, 25},
		3: {31, 32, 33, 34, 35},
	}
	del(v)
	fmt.Printf("%+v\n", v)

	v = map[int][]int{}
	for _, i := range v[0] {
		fmt.Println(i)
	}
}

func del(v map[int][]int) {
	delete(v, 2)
}
