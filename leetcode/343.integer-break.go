package main

import (
	"fmt"
	"math"
)

func integerBreak(n int) int {
	if n == 2 {
		return 1
	} else if n == 3 {
		return 2
	}
	max := 0
	switch n % 3 {
	case 0:
		max = int(math.Pow(float64(3), float64(n / 3)))
	case 1:
		max = int(math.Pow(float64(3), float64(n / 3) - 1)) * 4
	default:
		max = int(math.Pow(float64(3), float64(n / 3))) * 2
	}
	return max
}

func main() {
	fmt.Println(integerBreak(2), 1)
	fmt.Println(integerBreak(3), 2)
	fmt.Println(integerBreak(4), 4)
	fmt.Println(integerBreak(8), 18)
	fmt.Println(integerBreak(10), 36)
}

//7
//3 4						12 <-
//2 2 3						12

//8
//4 4						16
//4 2 2						16
//3 3 2						18 <-
//2 2 2 2					16

//9
//5 4						20
//3 3 3						27  <-
//3 2 2 2					24

//10
//5 5						25
//3 3 4						36 <-
//3 3 2 2					36
//2 2 2 2 2					32

//32
//2		16 * 16												256
//3		11 * 11 * 10										1210
//4		8  * 8  * 8  * 8									4096
//5		6  * 6  * 6  * 7 * 7								10584
//6		5  * 5  * 5  * 5 * 6 * 6 							22500
//7		5  * 5  * 5  * 5 * 4 * 4 * 4						40000
//8		4  * 4  * 4  * 4 * 4 * 4 * 4 * 4					65536
//9		4  * 4  * 4  * 4 * 4 * 3 * 3 * 3 * 3				82944
//10	4  * 4  * 3  * 3 * 3 * 3 * 3 * 3 * 3 * 3			104976
//11	3  * 3  * 3  * 3 * 3 * 3 * 3 * 3 * 3 * 3 * 2		118098 <-



