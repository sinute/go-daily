package main

import "fmt"

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
    Val  int
    Next *ListNode
}

//19. 删除链表的倒数第N个节点
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

func printListNode(head *ListNode) {
    for head != nil {
        fmt.Print(head.Val, " ")
        head = head.Next
    }
    fmt.Println("nil")
}

func main() {
    printListNode(removeNthFromEnd(&ListNode{
        Val: 1,
        Next: &ListNode{
            Val:  2,
            Next: nil,
        },
    }, 1))
    printListNode(&ListNode{
        Val:  1,
        Next: nil,
    })

    fmt.Println("===")

    printListNode(removeNthFromEnd(&ListNode{
        Val: 1,
        Next: &ListNode{
            Val:  2,
            Next: nil,
        },
    }, 2))
    printListNode(&ListNode{
        Val:  2,
        Next: nil,
    })

    fmt.Println("===")

    printListNode(removeNthFromEnd(&ListNode{
        Val: 1,
        Next: &ListNode{
            Val: 2,
            Next: &ListNode{
                Val: 3,
                Next: &ListNode{
                    Val: 4,
                    Next: &ListNode{
                        Val:  5,
                        Next: nil,
                    },
                },
            },
        },
    }, 2))
    printListNode(&ListNode{
        Val: 1,
        Next: &ListNode{
            Val: 2,
            Next: &ListNode{
                Val: 3,
                Next: &ListNode{
                    Val:  5,
                    Next: nil,
                },
            },
        },
    })

    fmt.Println("===")

    printListNode(removeNthFromEnd(&ListNode{
        Val:  1,
        Next: nil,
    }, 1))
    printListNode(nil)

    fmt.Println("===")
}
