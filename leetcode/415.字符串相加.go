package leetcode

/*
 * @lc app=leetcode.cn id=415 lang=golang
 *
 * [415] 字符串相加
 */

// @lc code=start
func addStrings(num1 string, num2 string) string {
	l1 := len(num1)
	l2 := len(num2)
	num := make([]byte, 0)
	carry := 0
	v := 0
	for p1, p2 := l1-1, l2-1; p1 >= 0 || p2 >= 0; p1, p2 = p1-1, p2-1 {
		if p1 >= 0 && p2 >= 0 {
			v = int(num1[p1]) + int(num2[p2]) + carry - 48 - 48
		} else if p1 >= 0 && p2 < 0 {
			v = int(num1[p1]) + carry - 48
		} else if p1 < 0 && p2 >= 0 {
			v = int(num2[p2]) + carry - 48
		}
		carry = 0
		if v >= 10 {
			carry = 1
			v -= 10
		}
		num = append([]byte{byte(v + 48)}, num...)
	}
	if carry == 1 {
		num = append([]byte{byte(1 + 48)}, num...)
	}
	return string(num)
}

// @lc code=end
