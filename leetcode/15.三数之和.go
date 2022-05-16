package leetcode

import "sort"

/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */

// @lc code=start
func threeSum(nums []int) [][]int {
	l := len(nums)
	if l == 3 {
		if nums[0]+nums[1]+nums[2] == 0 {
			return [][]int{nums}
		} else {
			return [][]int{}
		}
	}
	sort.Ints(nums)
	m := make(map[int]map[int]map[int]bool)
	for i := 0; i < l; i++ {
		for start, end := i+1, l-1; start < end; {
			if nums[i]+nums[start]+nums[end] == 0 {
				if _, ok := m[nums[i]]; !ok {
					m[nums[i]] = make(map[int]map[int]bool)
				}
				if _, ok := m[nums[i]][nums[start]]; !ok {
					m[nums[i]][nums[start]] = make(map[int]bool)
				}
				m[nums[i]][nums[start]][nums[end]] = true
				start++
			} else if nums[i]+nums[start]+nums[end] > 0 {
				end--
			} else {
				start++
			}
		}
	}
	v := make([][]int, 0)
	for i := range m {
		for j := range m[i] {
			for k := range m[i][j] {
				v = append(v, []int{i, j, k})
			}
		}
	}
	return v
}

// @lc code=end
