package leetcode

import "math"

/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */

// @lc code=start
func maxArea(height []int) int {
	l := len(height)
	max := 0
	for i, j := 0, l-1; j > i; {
		s := int(math.Min(float64(height[i]), float64(height[j])) * float64(j-i))
		if s > max {
			max = s
		}
		if height[i] > height[j] {
			j--
		} else {
			i++
		}

	}
	return max
}

// @lc code=end
