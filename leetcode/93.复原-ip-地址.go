package leetcode

import (
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode.cn id=93 lang=golang
 *
 * [93] 复原 IP 地址
 */

// @lc code=start
func restoreIpAddresses(s string) []string {
	r := []string{}
	parts := split(s, []string{})
	for i := range parts {
		r = append(r, strings.Join(parts[i], "."))
	}
	return r
}

func split(s string, ss []string) [][]string {
	//ss = []string{"1","100"}
	l := len(s)
	if len(ss) >= 4 {
		// 拆分成4部分的情况
		if l == 0 {
			// 如果原始字符串已被拆分完，则当前部分就是拆分结果
			return [][]string{ss}
		} else {
			// 否则不是一个合理的拆分，返回空
			return [][]string{}
		}
	}
	// 如果还没拆分成4部分时原始字符串就已经为空，则不是一个合理的拆分，返回空
	if l == 0 {
		return [][]string{}
	}
	r := [][]string{}
	// 1位
	ss1 := append([]string{s[len(s)-1:]}, ss...)
	r = append(r, split(s[0:len(s)-1], ss1)...)
	// 2位
	if l >= 2 {
		v := s[len(s)-2:]
		if v[0] != '0' {
			ss2 := append([]string{s[len(s)-2:]}, ss...)
			r = append(r, split(s[0:len(s)-2], ss2)...)
		}
	}
	// 3位
	if l >= 3 {
		v := s[len(s)-3:]
		intV, _ := strconv.Atoi(v)
		if v[0] != '0' && intV <= 255 {
			ss3 := append([]string{s[len(s)-3:]}, ss...)
			r = append(r, split(s[0:len(s)-3], ss3)...)
		}
	}
	return r
}

// @lc code=end
