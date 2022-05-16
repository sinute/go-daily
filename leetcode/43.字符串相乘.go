package leetcode

/*
 * @lc app=leetcode.cn id=43 lang=golang
 *
 * [43] 字符串相乘
 */

// @lc code=start
func multiply(num1 string, num2 string) string {
	carry := 0
	r := make([]byte, 0)
	l1 := len(num1)
	l2 := len(num2)
	l := l1 + l2
	values := make([][]byte, l)
	for j := l2 - 1; j >= 0; j-- {
		for i := l1 - 1; i >= 0; i-- {
			v := (num1[i]-'0')*(num2[j]-'0') + byte(carry)
			carry = int(v / 10)
			v %= 10
			values[l1-1-i+l2-1-j] = append(values[l1-1-i+l2-1-j], v)
		}
		if carry > 0 {
			values[l1+l2-1-j] = append(values[l1+l2-1-j], byte(carry))
			carry = 0
		}
	}

	carry = 0
	for i := 0; i < l; i++ {
		v := carry
		for j := len(values[i]) - 1; j >= 0; j-- {
			v += int(values[i][j])
		}
		carry = v / 10
		v %= 10
		r = append([]byte{byte(v) + '0'}, r...)
	}

	for len(r) > 1 && r[0] == '0' {
		r = r[1:]
	}
	return string(r)
}

// @lc code=end
