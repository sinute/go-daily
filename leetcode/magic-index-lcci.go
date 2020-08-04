package main

import (
	"fmt"
)

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

// 面试题 08.03. 魔术索引
func main() {
	fmt.Println(findMagicIndex([]int{-294354269, -168926144, -62738268, 3}), 3)
	fmt.Println(findMagicIndex([]int{0, 2, 3, 4, 5}), 0)
	fmt.Println(findMagicIndex([]int{1, 1, 1}), 1)
	fmt.Println(findMagicIndex([]int{0}), 0)
	fmt.Println(findMagicIndex([]int{1}), -1)
	fmt.Println(findMagicIndex([]int{5}), -1)
	fmt.Println(findMagicIndex([]int{-531369933, -469065528, -430059048, -428981853, -319235969, -288076332, -286667432, -282312559, -197049680, -197022263, -174416117, -138027773, -121899023, -111631966, -107567458, -70437707, -52463072, -45519851, -38641451, -15825815, -3835472, -1525043, 22, 566842886, 593757472, 605439236, 619794079, 640069993, 657657758, 718772950, 815849552, 839357142, 936585256, 1006188278, 1042347147, 1057129320, 1062178586, 1069769438}), 22)
	fmt.Println(findMagicIndex([]int{1, 1, 1, 4, 4, 4, 6, 6, 6}), 1)
}
