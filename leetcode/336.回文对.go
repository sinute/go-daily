package leetcode

/*
 * @lc app=leetcode.cn id=336 lang=golang
 *
 * [336] 回文对
 */

// @lc code=start
func isPalindrome336(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func palindromePairs(words []string) [][]int {
	dict := make(map[string]int)
	l := len(words)
	for i := 0; i < l; i++ {
		dict[reverse336(words[i])] = i
	}
	r := make([][]int, 0)
	for i := 0; i < l; i++ {
		if words[i] == "" {
			for j := 0; j < l; j++ {
				if i == j {
					continue
				}
				if isPalindrome336(words[j]) {
					r = append(r, []int{i, j})
				}
			}
			continue
		}
		for j := 0; j < len(words[i]); j++ {
			left := words[i][0:j]
			right := words[i][j:]
			if index, ok := dict[left]; ok && index != i && isPalindrome336(right) {
				// i: abcddd
				// index: cba
				r = append(r, []int{i, index})
			}
			if index, ok := dict[right]; ok && index != i && isPalindrome336(left) {
				// i: abcddd
				// index: dddcba
				r = append(r, []int{index, i})
			}
		}
	}
	return r
}

func reverse336(s string) string {
	r := []byte(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// @lc code=end
