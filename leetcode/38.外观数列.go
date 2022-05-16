package leetcode

/*
 * @lc app=leetcode.cn id=38 lang=golang
 *
 * [38] 外观数列
 */

// @lc code=start
func countAndSay(n int) string {
	s := "1"
	for i := 1; i < n; i++ {
		s = cas(s)
	}

	return s
}

func cas(s string) string {
	repeat := 0
	curr := byte(0)
	stack := make([]byte, 0)

	b := []byte(s)

	for i := 0; i < len(b); i++ {
		if curr == byte(0) || curr == b[i] {
			curr = s[i]
			repeat++
			continue
		}
		stack = append(stack, byte(repeat+48), curr)
		repeat = 1
		curr = b[i]
	}

	if repeat > 0 {
		stack = append(stack, byte(repeat+48), curr)
	}

	return string(stack)
}

// @lc code=end
