package main

import (
	"fmt"
)

//337. 打家劫舍 III
/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//func rob(root *TreeNode) int {
//	if root == nil {
//		return 0
//	}
//	a := root.Val
//	b := 0
//	if root.Left != nil {
//		a += rob(root.Left.Left) + rob(root.Left.Right)
//		b += rob(root.Left)
//	}
//	if root.Right != nil {
//		a += rob(root.Right.Left) + rob(root.Right.Right)
//		b += rob(root.Right)
//	}
//	if a > b {
//		return a
//	}
//	return b
//}

func rob(root *TreeNode) int {
	v := dfs(root)
	return max(v[0], v[1])
}

func dfs(root *TreeNode) []int {
	if root == nil {
		return []int{0, 0}
	}
	left := dfs(root.Left)
	right := dfs(root.Right)
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

func main() {
	fmt.Println(rob(&TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val:  1,
			Left: nil,
			Right: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}), 7)

	fmt.Println(rob(&TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:  2,
			Left: nil,
			Right: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:  3,
			Left: nil,
			Right: &TreeNode{
				Val:   1,
				Left:  nil,
				Right: nil,
			},
		},
	}), 7)

	fmt.Println(rob(&TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val:   1,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:  5,
			Left: nil,
			Right: &TreeNode{
				Val:   1,
				Left:  nil,
				Right: nil,
			},
		},
	}), 9)
}
