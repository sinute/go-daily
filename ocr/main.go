package main

import (
	"fmt"

	"github.com/otiai10/gosseract/v2"
)

func main() {
	client := gosseract.NewClient()
	defer client.Close()
	client.SetLanguage("chi_sim")
	err := client.SetImage("你好.png")
	if err != nil {
		panic(err)
	}
	text, err := client.Text()
	fmt.Println(text)
	fmt.Println(err)
	// Hello, World!
}
