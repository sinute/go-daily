package main

import (
	"log"
	"testing"
)

func catch() {
	log.Println("catch:", recover())
}

func Test_panic(t *testing.T) {
	defer catch()
	defer log.Println(recover())
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()

	panic("i am dead")
}

func Test_panic2(t *testing.T) {
	z := 0
	x := 0
	y := 0
	func() {
		defer func() {
			if recover() != nil {
				z = 999
			}

		}()
		z = x / y

	}()

	log.Println("x / y = ", z)
}
