package main

import "fmt"

func isSubsequence(s string, t string) bool {
	lenS := len(s)
	lenT := len(t)
	if lenS == 0 {
		return true
	}
	if  lenT == 0 || lenS > lenT {
		return false
	}
	j := 0
	for i := 0; i < lenT; i++ {
		if s[j] == t[i] {
			j++
		}
		if j == lenS {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(isSubsequence("abc", "ahbgdc"), true)
	fmt.Println(isSubsequence("a", "ahbgdc"), true)
	fmt.Println(isSubsequence("", "ahbgdc"), false)
}
