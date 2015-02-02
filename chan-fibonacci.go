package main

import (
	"fmt"
)

func fibonacci() chan int {
	in := make(chan int, 2)
	in <- 0
	in <- 1
	out := make(chan int)
	go func(in, out chan int) {
		for {
			a := <-in
			in <- <-in + a
			in <- a
			out <- a
		}
	}(in, out)
	return out
}

func main() {
	x := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(<-x)
	}
}
