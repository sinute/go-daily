package leetcode

/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 */

// @lc code=start
var alpDig = [][]string{
	{},
	{},
	{"a", "b", "c"},
	{"d", "e", "f"},
	{"g", "h", "i"},
	{"j", "k", "l"},
	{"m", "n", "o"},
	{"p", "q", "r", "s"},
	{"t", "u", "v"},
	{"w", "x", "y", "z"},
}

//17. 电话号码的字母组合
func letterCombinations(digits string) []string {
	l := len(digits)
	if l == 0 {
		return []string{}
	}
	r := alpDig[digits[0]-48]
	for i := 1; i < l; i++ {
		r = Combinations(r, digits[i])
	}
	return r
}

func Combinations(strs []string, chr byte) []string {
	l := len(strs)
	r := make([]string, 0)
	for i := 0; i < l; i++ {
		for j := 0; j < len(alpDig[chr-48]); j++ {
			r = append(r, strs[i]+alpDig[chr-48][j])
		}
	}
	return r
}

// @lc code=end
