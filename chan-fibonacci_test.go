package main

import (
	"fmt"
	"testing"
)

func Test_fibonacci(t *testing.T) {
	x := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(<-x)
	}
}
