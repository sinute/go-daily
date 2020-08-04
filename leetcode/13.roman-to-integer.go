package main

import (
	"fmt"
)

func romanToInt(s string) int {
	//I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
	//X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。 
	//C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。
	m := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	v := m[s[len(s)-1]]
	for i := len(s) - 2; i >= 0; i-- {
		if m[s[i]] >= m[s[i+1]] {
			v += m[s[i]]
		} else {
			v -= m[s[i]]
		}
	}
	return v
}

func main() {
	fmt.Println(romanToInt("III"), 3)
	fmt.Println(romanToInt("IV"), 4)
	fmt.Println(romanToInt("IX"), 9)
	fmt.Println(romanToInt("LVIII"), 58)
	fmt.Println(romanToInt("MCMXCIV"), 1994)
}
