package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func reply(conn net.Conn) {
	for {
		buf, err := bufio.NewReader(conn).ReadBytes('\n')
		if err != nil {
			fmt.Println(conn.RemoteAddr(), "[ Disconnected ]")
			return
		}
		if strings.TrimRight(string(buf), "\n\r") == "exit" {
			conn.Close()
			fmt.Println(conn.RemoteAddr(), "[ Disconnected ]")
			return
		}
		if _, err = conn.Write(buf); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Print(conn.RemoteAddr(), " - ", string(buf))
	}
}

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:12345")
	defer listener.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
		fmt.Println(conn.RemoteAddr(), "[ Connected ]")
		go reply(conn)
	}
}
