package leetcode

/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] 最长公共前缀
 */

// @lc code=start
func longestCommonPrefix(strs []string) string {
	l := len(strs)
	if l == 0 {
		return ""
	} else if l == 1 {
		return strs[0]
	}
	longest := 0
	b := byte(0)
	for true {
		for i := 0; i < l; i++ {
			if longest >= len(strs[i]) {
				return strs[0][0:longest]
			}
			if i == 0 {
				b = strs[i][longest]
			} else if b != strs[i][longest] {
				return strs[0][0:longest]
			}
		}
		longest++
	}
	return strs[0][0:longest]
}

// @lc code=end
