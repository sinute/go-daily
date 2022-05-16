package leetcode

/*
 * @lc app=leetcode.cn id=696 lang=golang
 *
 * [696] 计数二进制子串
 */

// @lc code=start
func countBinarySubstrings(s string) int {
	l := len(s)
	if l <= 1 {
		return 0
	}
	r := 0
	curr := s[0]
	for i := 1; i < l; {
		start, end := i-2, i+1
		if s[i] != curr {
			curr = s[i]
			r++
			for ; start >= 0 && end < l && s[end] == curr && s[start] != curr; start, end = start-1, end+1 {
				r++
			}
		}
		if end < l {
			i = end
		} else {
			break
		}
	}
	return r
}

// @lc code=end
