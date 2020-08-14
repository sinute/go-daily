package main

import (
    "fmt"
)

type ListNode struct {
    Val  int
    Next *ListNode
}

//2. 两数相加
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

func main() {
    l1 := &ListNode{
        2,
        &ListNode{
            4,
            &ListNode{
                3,
                nil,
            },
        },
    }
    l2 := &ListNode{
        5,
        &ListNode{
            6,
            &ListNode{
                4,
                nil,
            },
        },
    }

    l1 = &ListNode{
        0,
        nil,
    }
    l2 = &ListNode{
        7,
        &ListNode{
            3,
            nil,
        },
    }

    l1 = &ListNode{
        9,
        &ListNode{
            9,
            nil,
        },
    }
    l2 = &ListNode{
        9,
        nil,
    }

    l := addTwoNumbers(l1, l2)
    for l != nil {
        fmt.Println(l.Val)
        l = l.Next
    }
    //fmt.Println(twoSum([]int{2,5,15,7,11}, 9))
    //fmt.Println(twoSum([]int{-1,0}, -1))
}
