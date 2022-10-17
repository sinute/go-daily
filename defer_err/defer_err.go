package main

import (
	"fmt"
)

func main() {
	fmt.Println(deferErrWillNotReturn())
	fmt.Println(deferErrWillReturn())
}

func newErr() error {
	return fmt.Errorf("err")
}

func deferErrWillNotReturn() error {
	var err error
	defer newErr()
	return err
}

func deferErrWillReturn() (err error) {
	defer func() {
		anotherRrr := newErr()
		if anotherRrr != nil {
			err = anotherRrr
		}
	}()
	return err
}
