package leetcode

/*
 * @lc app=leetcode.cn id=7 lang=golang
 *
 * [7] 整数反转
 */

// @lc code=start
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

// @lc code=end
