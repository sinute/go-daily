package leetcode

import (
	"bytes"
	"strings"
)

/*
 * @lc app=leetcode.cn id=12 lang=golang
 *
 * [12] 整数转罗马数字
 */

// @lc code=start
func intToRoman(num int) string {
	//I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
	//X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。
	//C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。
	m := map[int]string{
		1:    "I",
		4:    "IV",
		5:    "V",
		9:    "IX",
		10:   "X",
		40:   "XL",
		50:   "L",
		90:   "XC",
		100:  "C",
		400:  "CD",
		500:  "D",
		900:  "CM",
		1000: "M",
	}
	v := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	buf := bytes.Buffer{}
	for i := 0; num > 0; i++ {
		t := num / v[i]
		buf.WriteString(strings.Repeat(m[v[i]], t))
		num -= t * v[i]
	}
	return buf.String()
}

// @lc code=end
