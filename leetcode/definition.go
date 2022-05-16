package leetcode

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Node struct {
	Val       int
	Neighbors []*Node
}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}
