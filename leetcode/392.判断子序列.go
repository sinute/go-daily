package leetcode

/*
 * @lc app=leetcode.cn id=392 lang=golang
 *
 * [392] 判断子序列
 */

// @lc code=start
func isSubsequence(s string, t string) bool {
	lenS := len(s)
	lenT := len(t)
	if lenS == 0 {
		return true
	}
	if lenT == 0 || lenS > lenT {
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

// @lc code=end
