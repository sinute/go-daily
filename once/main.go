package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	m map[int]int
	o sync.Once
)

func main() {
	m = make(map[int]int)
	//go func() {
	//	for true {
	//		insert(1,1)
	//	}
	//}()
	//go func() {
	//	for true {
	//		insert(1,1)
	//	}
	//}()
	go func() {
		for true {
			insertOnce(1, 1)
		}
	}()
	go func() {
		for true {
			insertOnce(1, 1)
		}
	}()
	time.Sleep(5 * time.Second)
	//go insert(1,2)
	//go insert(1,2)
	//go insertOnce(1)
	//go insertOnce(1)
}

func insert(id int, value int) {
	m[id] = value
	fmt.Printf("insert id(%d), value(%d)\n", id, value)
}

func insertOnce(id int, value int) {
}
