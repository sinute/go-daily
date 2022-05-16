package leetcode

import "sort"

/*
 * @lc app=leetcode.cn id=18 lang=golang
 *
 * [18] 四数之和
 */

// @lc code=start
func fourSum(nums []int, target int) [][]int {
	l := len(nums)
	r := make([][]int, 0)
	if l < 4 {
		return r
	}
	sort.Ints(nums)
	tmp := make(map[int]map[int]map[int]map[int]bool)
	for i := 0; i < l-3; i++ {
		for j := i + 1; j < l-2; j++ {
			// 外层循环整体趋势变大
			for start, end := j+1, l-1; start < end; {
				// 内层双指针
				if nums[i]+nums[j]+2*nums[start] > target || nums[i]+nums[j]+2*nums[end] < target {
					break
				}
				if nums[i]+nums[j]+nums[start]+nums[end] > target {
					end--
				} else if nums[i]+nums[j]+nums[start]+nums[end] < target {
					start++
				} else {
					if _, ok := tmp[nums[i]]; !ok {
						tmp[nums[i]] = make(map[int]map[int]map[int]bool)
					}
					if _, ok := tmp[nums[i]][nums[j]]; !ok {
						tmp[nums[i]][nums[j]] = make(map[int]map[int]bool)
					}
					if _, ok := tmp[nums[i]][nums[j]][nums[start]]; !ok {
						tmp[nums[i]][nums[j]][nums[start]] = make(map[int]bool)
					}
					tmp[nums[i]][nums[j]][nums[start]][nums[end]] = true
					start++
				}
			}
		}
	}
	return uniqMap(tmp)
}

func uniqMap(tmp map[int]map[int]map[int]map[int]bool) [][]int {
	r := make([][]int, 0)
	for i := range tmp {
		for j := range tmp[i] {
			for start := range tmp[i][j] {
				for end := range tmp[i][j][start] {
					r = append(r, []int{i, j, start, end})
				}
			}
		}
	}
	return r
}

// @lc code=end
