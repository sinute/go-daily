package leetcode

import "math"

/*
 * @lc app=leetcode.cn id=343 lang=golang
 *
 * [343] 整数拆分
 */

// @lc code=start
func integerBreak(n int) int {
	if n == 2 {
		return 1
	} else if n == 3 {
		return 2
	}
	max := 0
	switch n % 3 {
	case 0:
		max = int(math.Pow(float64(3), float64(n/3)))
	case 1:
		max = int(math.Pow(float64(3), float64(n/3)-1)) * 4
	default:
		max = int(math.Pow(float64(3), float64(n/3))) * 2
	}
	return max
}

// @lc code=end
