package leetcode

/*
 * @lc app=leetcode.cn id=4 lang=golang
 *
 * [4] 寻找两个正序数组的中位数
 */

// @lc code=start
// func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
// 	// 长度, 偏移
// 	l1 := len(nums1)
// 	l2 := len(nums2)
// 	if l1 < l2 {
// 		return findMedianSortedArrays(nums2, nums1)
// 	}
// 	pos1 := l1 / 2
// 	pos2 := l2 / 2
// 	hasTwo := false
// 	if (l1+l2)&1 == 1 {
// 		hasTwo = true
// 	}
// 	// 奇数 中位数位置为 l / 2
// 	// 偶数 中位数位置为 l / 2 和 l / 2 - 1
// 	// 二分查找
// 	for true {
// 		if hasTwo {

// 		} else {
// 			if nums1[pos1] >= nums2[pos2-1] && nums1[pos1] <= nums2[pos2] {
// 				//1 2 |3| 4
// 				//  2 |3| 6
// 				return float64(nums1[pos1])
// 			} else if nums1[pos1] >= nums2[pos2-1] && nums1[pos1] > nums2[pos2] {
// 				p := (pos1 + 1) / 2
// 				//1 2 |3| 4
// 				//  2 |2| 6

// 				// 1 |2| 3 4
// 				// 2 |2| 6
// 			} else if nums1[pos1] < nums2[pos2-1] && nums1[pos1] <= nums2[pos2] {

// 			} else if nums1[pos1] < nums2[pos2-1] && nums1[pos1] > nums2[pos2] {

// 			}
// 		}
// 	}
// }

// @lc code=end

// func aa(nums1 []int, pos1 int, nums2 []int, pos2 int) (i, j int) {
// 	if nums1[pos1] >= nums2[pos2-1] && nums1[pos1] <= nums2[pos2] {
// 		return pos1, pos2
// 	} else if nums1[pos1] >= nums2[pos2-1] && nums1[pos1] > nums2[pos2] {
// 		p := (pos1 + 1) / 2
// 		//1 2 |3| 4
// 		//  2 |2| 6

// 		// 1 |2| 3 4
// 		// 2 |2| 6
// 	} else if nums1[pos1] < nums2[pos2-1] && nums1[pos1] <= nums2[pos2] {

// 	} else if nums1[pos1] < nums2[pos2-1] && nums1[pos1] > nums2[pos2] {

// 	}
// }
