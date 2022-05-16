package leetcode

/*
 * @lc app=leetcode.cn id=6 lang=golang
 *
 * [6] Z 字形变换
 */

// @lc code=start
func convert(s string, numRows int) string {
	l := len(s)
	if numRows >= l || numRows == 1 {
		return s
	}
	str := make([]byte, 0, l)
	for i := 0; i < numRows; i++ {
		for j := 0; true; j++ {
			y := j*(2*numRows-2) + i
			x := y - 2*i // j * (2 * numRows - 2) - i
			if x >= l {
				break
			}
			if i == 0 {
				// 第一行
				if y > l-1 {
					continue
				}
				str = append(str, s[y])
			} else if i == numRows-1 {
				// 最后一行
				if x < 0 {
					continue
				}
				str = append(str, s[x])
			} else {
				if x < 0 {
					str = append(str, s[y])
				} else if y >= l {
					str = append(str, s[x])
				} else {
					str = append(str, s[x], s[y])
				}
			}
		}
	}
	return string(str)
}

// @lc code=end
