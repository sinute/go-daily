package main

import (
	"fmt"
	"math"
	"sort"
)

//16. 最接近的三数之和
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	l := len(nums)
	closest := 0
	init := false
	for i := 0; i < l-2; i++ {
		for start, end := i+1, l-1; start < end; {
			current := nums[i] + nums[start] + nums[end]
			distance := current-target
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

func main() {
	//fmt.Println(threeSumClosest([]int{-1, 2, 1, -4}, 1), 2)
	fmt.Println(threeSumClosest([]int{1,1,1,0}, -100), 2)
}
