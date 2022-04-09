package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func ReadFile() {
	f, _ := os.Open("readme.md")
	defer f.Close()
	r := bufio.NewReader(f)
	buf := make([]byte, 1024)
	for {
		n, err := r.Read(buf)
		if err == io.EOF {
			break
		}
		fmt.Println(string(buf[:n]))
	}
}
