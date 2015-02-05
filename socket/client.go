package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	if conn, err := net.Dial("tcp", "127.0.0.1:8888"); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	} else {
		rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
		for {
			if buf, err := bufio.NewReader(os.Stdin).ReadBytes('\n'); err != nil {
				fmt.Println(err)
				os.Exit(-1)
			} else {
				rw.Write(buf)
				rw.Flush()
				switch buf, err = rw.ReadBytes('\n'); err {
				case io.EOF:
					fmt.Println("Server [Closed]")
					os.Exit(0)
				case nil:
					fmt.Print("Server:", string(buf))
				default:
					fmt.Println(err)
					os.Exit(-1)
				}
			}
		}
	}
}
