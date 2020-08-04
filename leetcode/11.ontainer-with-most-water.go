package main

import (
	"fmt"
	"math"
)

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

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}), 49)
	fmt.Println(maxArea([]int{2,3,4,5,18,17,6}), 17)
	fmt.Println(maxArea([]int{1, 2}), 1)
}
