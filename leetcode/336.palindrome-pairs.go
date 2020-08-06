package main

import (
	"fmt"
)

//336. 回文对
//func palindromePairs(words []string) [][]int {
//	r := make([][]int, 0)
//	l := len(words)
//
//	for i := 0; i < l; i++ {
//		for j := 0; j < l; j++ {
//			if i == j {
//				continue
//			}
//			if isPalindrome(words[i] + words[j]) {
//				r = append(r, []int{i, j})
//			}
//		}
//	}
//	return r
//}

func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func palindromePairs(words []string) [][]int {
	dict := make(map[string]int)
	l := len(words)
	for i := 0; i < l; i++ {
		dict[reverse(words[i])] = i
	}
	r := make([][]int, 0)
	for i := 0; i < l; i++ {
		if words[i] == "" {
			for j := 0; j < l; j++ {
				if i == j {
					continue
				}
				if isPalindrome(words[j]) {
					r = append(r, []int{i, j})
				}
			}
			continue
		}
		for j := 0; j < len(words[i]); j++ {
			left := words[i][0:j]
			right := words[i][j:]
			if index, ok := dict[left]; ok && index != i && isPalindrome(right) {
				// i: abcddd
				// index: cba
				r = append(r, []int{i, index})
			}
			if index, ok := dict[right]; ok && index != i && isPalindrome(left) {
				// i: abcddd
				// index: dddcba
				r = append(r, []int{index, i})
			}
		}
	}
	return r
}

func reverse(s string) string {
	r := []byte(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func main() {
	fmt.Println(palindromePairs([]string{"a", ""}), [][]int{{1, 0}, {0, 1}})
	fmt.Println(palindromePairs([]string{"a", "b", "c", "ab", "ac", "aa"}), [][]int{{3, 0}, {1, 3}, {4, 0}, {2, 4}, {5, 0}, {0, 5}})
	fmt.Println(palindromePairs([]string{"lls", "s"}), [][]int{{1, 0}})
	fmt.Println(palindromePairs([]string{"abcd", "dcba", "lls", "s", "sssll"}), [][]int{{0, 1}, {1, 0}, {3, 2}, {2, 4}})
	fmt.Println(palindromePairs([]string{"a", "abc", "aba", ""}), [][]int{{0, 3}, {3, 0}, {3, 2}, {2, 3}})
	fmt.Println(palindromePairs([]string{"bat", "tab", "cat"}), [][]int{{0, 1}, {1, 0}})
}
