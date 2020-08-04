package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	//for i:=0; i<len(numbers);i++  {
	//	for j:=i+1; j<len(numbers);j++  {
	//		if numbers[i] + numbers[j] == target {
	//			return []int{i+1,j+1}
	//		}
	//	}
	//}
	//return []int{0,0}
	i,j:=0,len(numbers)-1
	for i<j {
		sum := numbers[i] + numbers[j]
		if  sum == target {
			return []int{i+1,j+1}
		} else if sum > target{
			j--
		} else {
			i++
		}
	}
	return []int{0,0}
}



func main() {
	fmt.Println(twoSum([]int{2,7,11,15}, 9))
	fmt.Println(twoSum([]int{-1,0}, -1))
}