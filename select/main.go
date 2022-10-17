package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second)

	var t time.Time
	for {
		select {
		case <-ticker.C:
			time.Sleep(time.Second * 2)
			fmt.Println(t)
		}
		fmt.Println("loop")
	}
}
