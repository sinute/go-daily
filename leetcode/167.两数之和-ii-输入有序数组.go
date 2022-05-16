package leetcode

/*
 * @lc app=leetcode.cn id=167 lang=golang
 *
 * [167] 两数之和 II - 输入有序数组
 */

// @lc code=start
func twoSum(numbers []int, target int) []int {
	//for i:=0; i<len(numbers);i++  {
	//	for j:=i+1; j<len(numbers);j++  {
	//		if numbers[i] + numbers[j] == target {
	//			return []int{i+1,j+1}
	//		}
	//	}
	//}
	//return []int{0,0}
	i, j := 0, len(numbers)-1
	for i < j {
		sum := numbers[i] + numbers[j]
		if sum == target {
			return []int{i + 1, j + 1}
		} else if sum > target {
			j--
		} else {
			i++
		}
	}
	return []int{0, 0}
}

// @lc code=end
