package main

import (
	"fmt"
	"regexp"
)

func main() {
	expr := `^\/\d+$`
	var reg *regexp.Regexp
	var err error
	if reg, err = regexp.Compile(expr); err != nil {
		panic(err)
	}

	s := "/123"
	if !reg.MatchString(s) {
		fmt.Printf("s: %s, not match\n", s)
	} else {
		fmt.Printf("s: %s, match\n", s)
	}
}
