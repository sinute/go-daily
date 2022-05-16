package leetcode

import (
	"math"
	"sort"
)

/*
 * @lc app=leetcode.cn id=16 lang=golang
 *
 * [16] 最接近的三数之和
 */

// @lc code=start
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	l := len(nums)
	closest := 0
	init := false
	for i := 0; i < l-2; i++ {
		for start, end := i+1, l-1; start < end; {
			current := nums[i] + nums[start] + nums[end]
			distance := current - target
			if !init {
				closest = current
				init = true
			}
			if distance == 0 {
				return target
			} else {
				if math.Abs(float64(distance)) < math.Abs(float64(closest-target)) {
					closest = current
				}
				if distance > 0 {
					end--
				} else {
					start++
				}
			}
		}
	}
	return closest
}

// @lc code=end
