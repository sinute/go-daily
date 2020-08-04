package main

import "fmt"

func twoSum(nums []int, target int) []int {
	for i:=0; i<len(nums);i++  {
		for j:=i+1; j<len(nums);j++  {
			if nums[i] + nums[j] == target {
				return []int{i,j}
			}
		}
	}
	return []int{0,0}
}

func main() {
	fmt.Println(twoSum([]int{2,7,11,15}, 9))
	fmt.Println(twoSum([]int{2,5,15,7,11}, 9))
	fmt.Println(twoSum([]int{-1,0}, -1))
}
