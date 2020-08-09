package main

import (
	"fmt"
)

var m = [][]string{
	{},
	{},
	{"a", "b", "c"},
	{"d", "e", "f"},
	{"g", "h", "i"},
	{"j", "k", "l"},
	{"m", "n", "o"},
	{"p", "q", "r", "s"},
	{"t", "u", "v"},
	{"w", "x", "y", "z"},
}

//17. 电话号码的字母组合
func letterCombinations(digits string) []string {
	l := len(digits)
	if l == 0 {
		return []string{}
	}
	r := m[digits[0]-48]
	for i := 1; i < l; i++ {
		r = Combinations(r, digits[i])
	}
	return r
}

func Combinations(strs []string, chr byte) []string {
	l := len(strs)
	r := make([]string, 0)
	for i := 0; i < l; i++ {
		for j := 0; j < len(m[chr-48]); j++ {
			r = append(r, strs[i]+m[chr-48][j])
		}
	}
	return r
}

func main() {
	fmt.Println(letterCombinations(""), []string{})
	fmt.Println(letterCombinations("2"), []string{"a", "b", "c"})
	fmt.Println(letterCombinations("23"), []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"})
}
