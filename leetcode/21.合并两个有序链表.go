package leetcode

/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}
	if l1.Val > l2.Val {
		return mergeTwoLists(l2, l1)
	}

	head := l1
	previous := l1
	l1 = l1.Next
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			previous.Next = l2
			l2 = l2.Next
			previous.Next.Next = l1
		} else {
			l1 = l1.Next
		}
		previous = previous.Next
	}
	if l2 != nil {
		previous.Next = l2
	}
	return head
}

// @lc code=end
