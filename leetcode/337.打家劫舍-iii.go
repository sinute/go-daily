package leetcode

/*
 * @lc app=leetcode.cn id=337 lang=golang
 *
 * [337] 打家劫舍 III
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rob(root *TreeNode) int {
	v := dfs337(root)
	return max(v[0], v[1])
}

func dfs337(root *TreeNode) []int {
	if root == nil {
		return []int{0, 0}
	}
	left := dfs337(root.Left)
	right := dfs337(root.Right)
	// 当前节点选中的最大值 = 当前节点的值 + 左侧节点未选中的最大值 + 右侧节点未选中的最大值
	selected := root.Val + left[1] + right[1]
	// 当前节点未选中的最大值 = 左侧节点选中或不选中的最大值 + 右侧节点选中或不选中的最大值
	unSelected := max(left[0], left[1]) + max(right[0], right[1])
	return []int{selected, unSelected}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
