package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	start, end, length, maxLength := 0, 0, 0, 0
	sub := make(map[byte]int)
	n := len(s)
	// end 当前子串尾端
	// pos 字符所在位置
	// start 当前子串开始位置
	for end < n  {
		pos, ok := sub[s[end]]
		if !ok || pos < start {
			// 未出现过 或者 出现的位置在子串开始位置之前，记录当前字符位置
			sub[s[end]] = end
			length++
		} else {
			// 出现重复字符
			// 记录当前最大长度
			if maxLength < length {
				maxLength = length
			}
			// 使用重复字符位置覆盖
			sub[s[end]] = end
			// 开始位置移动到重复的字符后一位
			start = pos + 1
			// 重置当前长度 = 结束位置 - 开始位置 + 1
			length = end - start + 1
		}
		// 本轮结束，末尾往后移一位
		end++
	}

	if maxLength < length {
		maxLength = length
	}
	return maxLength
}

func main() {
	fmt.Println(lengthOfLongestSubstring("dvdf"), 3)
	fmt.Println(lengthOfLongestSubstring("abcabcbb"), 3)
	fmt.Println(lengthOfLongestSubstring("bbb"), 1)
	fmt.Println(lengthOfLongestSubstring("pwwkew"), 3)
	fmt.Println(lengthOfLongestSubstring("a"), 1)
	fmt.Println(lengthOfLongestSubstring("au"), 2)
}
