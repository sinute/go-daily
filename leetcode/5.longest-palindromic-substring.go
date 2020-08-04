package main

import "fmt"

//func longestPalindrome(s string) string {
//	l := len(s)
//	if l <= 1 {
//		return s
//	}
//	start := 0
//	end := 0
//	maxLen := 1
//	maxLenStart := 0
//	m := make(map[int]map[int]bool)
//	for i := 0; i < l; i++ {
//		m[i] = make(map[int]bool)
//		m[i][i] = true
//	}
//	for end < l {
//		start = 0
//		for start < end {
//			if s[end] == s[start] {
//				if start + 1 > end - 1 {
//					m[start][end] = true
//				} else {
//					m[start][end] = m[start+1][end-1]
//				}
//				if m[start][end] && end+1-start > maxLen {
//					maxLen = end + 1 - start
//					maxLenStart = start
//				}
//			} else {
//				m[start][end] = false
//			}
//			start++
//		}
//		end++
//	}
//	return s[maxLenStart : maxLenStart+maxLen]
//}

func longestPalindrome(s string) string {
	start := 0
	l := len(s)
	maxLen := 0
	maxString := ""
	for l - start > maxLen  {
		end := l - 1
		for end - start + 1 > maxLen {
			if isPalindromicSubstring(s, start, end) {
				maxLen = end - start + 1
				maxString = s[start:end + 1]
			}
			end--
		}
		start++
	}
	return maxString
}

func isPalindromicSubstring(s string, start, end int) bool {
	for start < end {
		if s[start] == s[end] {
			start++
			end--
		} else {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(longestPalindrome("babad"), "bab")
	fmt.Println(longestPalindrome("cbbd"), "bb")
	fmt.Println(longestPalindrome("a"), "a")
	fmt.Println(longestPalindrome("ac"), "a")
	fmt.Println(longestPalindrome("dddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddd"), "dddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddd")
	//fmt.Println(isPalindromicSubstring("babad", 1, 3))
}
