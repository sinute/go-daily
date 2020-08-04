package main

import (
	"fmt"
)

//func reverse(x int) int {
//	symbol := 1
//	if x < 0 {
//		symbol = -1
//	}
//	x = symbol * x
//	s := strconv.Itoa(x)
//	lenX := len(s)
//	bytes := make([]byte,lenX)
//	for i, j := 0, lenX - 1; i < lenX || i < j; i++ {
//		bytes[i] = s[j]
//		bytes[j] = s[i]
//		j--
//	}
//	r,_ := strconv.Atoi(string(bytes))
//	r = symbol * r
//	if r <= -1 << 31 || r >= 1 << 31 - 1 {
//		return 0
//	}
//	return r
//}

func reverse(x int) int {
	r := 0
	for x != 0 {
		m := x % 10
		if r > 214748364 || (r == 214748364 && m > 7) {
			return 0
		}
		if r < -214748364 || (r == -214748364 && m < -8) {
			return 0
		}
		r = r*10 + m
		x /= 10
	}
	return r
}

func main() {
	fmt.Println(reverse(-123), -321)
	fmt.Println(reverse(1534236469), 0)
	fmt.Println(reverse(1563847412), 0)
	fmt.Println(reverse(1), 1)
	fmt.Println(reverse(120), 21)
	fmt.Println(reverse(-23), -32)
	fmt.Println(reverse(000), 0)
}
