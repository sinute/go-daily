package leetcode

/*
 * @lc app=leetcode.cn id=9 lang=golang
 *
 * [9] 回文数
 */

// @lc code=start
func isPalindrome(x int) bool {
	if x < 0 {
		// 负数不可能是回文数
		return false
	} else if x < 10 {
		// [0, 10) 一定是回文数
		return true
	} else if x%10 == 0 {
		// 末位是0不可能是回文数
		return false
	}
	y := 0
	for ; y < x; x = x / 10 {
		y = y*10 + x%10
	}
	// 1. 形如 2112 (21 == 21)
	// 2. 形如 12321 (12 == 123 / 10)
	return y == x || (y >= 10 && x == y/10)
}

// @lc code=end
