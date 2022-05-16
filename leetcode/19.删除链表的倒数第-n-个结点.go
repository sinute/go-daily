package leetcode

/*
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] 删除链表的倒数第 N 个结点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if n <= 0 {
		return head
	}
	r := head
	slow := head
	i := -1 - n
	for head != nil {
		if i >= 0 {
			slow = slow.Next
		}
		head = head.Next
		i++
	}
	if i < -1 {
		return r
	} else if i == -1 {
		return r.Next
	} else if i > -1 {
		slow.Next = slow.Next.Next
	}
	return r
}

// @lc code=end
