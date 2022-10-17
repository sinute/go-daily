package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx0 := context.Background()

	println(time.Now().Format("2006-01-02 15:04:05 start"))

	ctx1, cancel1 := context.WithTimeout(ctx0, 3*time.Second)
	defer cancel1()

	go func() {
		for {
			select {
			case <-ctx1.Done():
				println(time.Now().Format("2006-01-02 15:04:05 ctxI done"))
				return
			}
		}
	}()

	ctx2, cancel2 := context.WithTimeout(ctx1, 5*time.Second)
	defer cancel2()

	go func() {
		for {
			select {
			case <-ctx2.Done():
				println(time.Now().Format("2006-01-02 15:04:05 ctxII done"))
				return
			}
		}
	}()

	time.Sleep(10 * time.Second)
}

func main2() {
	fmt.Println("start", time.Now())

	ctx1, _ := context.WithTimeout(context.Background(), 5*time.Second)
	ctx2, cancel2 := context.WithTimeout(context.Background(), 2*time.Second)

	cancel2()

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		<-ctx1.Done()
		fmt.Println("ctx1", time.Now())
		wg.Done()
	}()
	go func() {
		<-ctx2.Done()
		fmt.Println("ctx2", time.Now())
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("end", time.Now())
	return
}
