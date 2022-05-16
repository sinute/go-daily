package leetcode

/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l := l1
	carry := 0
	for true {
		l1.Val += l2.Val
		if l1.Val >= 10 {
			carry = 1
			l1.Val -= 10
		}
		if l1.Next == nil && l2.Next != nil {
			l1.Next = &ListNode{
				Val:  carry,
				Next: nil,
			}
		} else if l1.Next != nil && l2.Next == nil {
			l2.Next = &ListNode{
				Val:  carry,
				Next: nil,
			}
		} else if l1.Next == nil && l2.Next == nil {
			if carry == 1 {
				l1.Next = &ListNode{
					Val:  carry,
					Next: nil,
				}
			}
			break
		} else {
			l1.Next.Val += carry
		}
		carry = 0
		l1, l2 = l1.Next, l2.Next
	}
	return l
}

// @lc code=end
