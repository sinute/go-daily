package leetcode

func findMagicIndex(nums []int) int {
	return findMagicIndexWithOffset(nums, 0)
}

func findMagicIndexWithOffset(nums []int, offset int) int {
	length := len(nums)
	if length == 0 {
		// 没有元素，返回-1
		return -1
	} else if length == 1 {
		// 只有一个元素，判断第0个元素是否和下标相等
		if nums[0] == offset {
			return nums[0]
		} else {
			return -1
		}
	}
	// 二分查找
	i := length / 2
	// 有序，从左侧开始查是否满足条件
	left := -1
	// 左侧快速跳过值比下标小的段
	if nums[i] >= offset {
		left = findMagicIndexWithOffset(nums[0:i], offset)
	}
	if left != -1 {
		// 如果有，即为结果
		return left
	} else if nums[i] == offset+i {
		// 左侧没有判断中间位置是否满足条件
		return nums[i]
	} else {
		// 最后在右侧位置查找是否可以满足条件
		right := findMagicIndexWithOffset(nums[i:length], i+offset)
		if right != -1 {
			return right
		}
	}
	return -1
}
