package main

import "fmt"

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
    Val  int
    Next *ListNode
}

//21. 合并两个有序链表
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

func printList(l *ListNode) {
    for l != nil {
        fmt.Print(string(l.Val+'0') + " -> ")
        l = l.Next
    }
    fmt.Println("nil")
}

func main() {
    l1 :=
        &ListNode{
            Val: 1,
            Next: &ListNode{
                Val: 2,
                Next: &ListNode{
                    Val:  4,
                    Next: nil,
                },
            },
        }
    l2 := &ListNode{
        Val: 1,
        Next: &ListNode{
            Val: 3,
            Next: &ListNode{
                Val:  4,
                Next: nil,
            },
        },
    }
    l3 := &ListNode{
        Val: 1,
        Next: &ListNode{
            Val: 1,
            Next: &ListNode{
                Val: 2,
                Next: &ListNode{
                    Val: 3,
                    Next: &ListNode{
                        Val: 4,
                        Next: &ListNode{
                            Val:  4,
                            Next: nil,
                        },
                    },
                },
            },
        },
    }
    printList(mergeTwoLists(l1, l2))
    printList(l3)
}
