package leetcode

/*
 * @lc app=leetcode.cn id=8 lang=golang
 *
 * [8] 字符串转换整数 (atoi)
 */

// @lc code=start
func myAtoi(str string) int {
	l := len(str)
	r := 0
	start := false
	symbol := 1
	for i := 0; i < l; i++ {
		// 空格跳过
		if !start && str[i] == 32 {
			continue
		}
		// 负数
		if !start && str[i] == 45 {
			symbol = -1
			start = true
			continue
		}
		// 正数
		if !start && str[i] == 43 {
			start = true
			continue
		}
		// 数字
		if str[i] >= 48 && str[i] <= 57 {
			// 计算当前字符数值
			m := symbol * (int(str[i]) - 48)
			start = true
			// 判断越界情况
			if r > 214748364 || (r == 214748364 && m > 7) {
				return 2147483647
			}
			if r < -214748364 || (r == -214748364 && m < -8) {
				return -2147483648
			}
			r = r*10 + m
		} else {
			// 其他字符
			break
		}
	}
	return r
}

// @lc code=end
