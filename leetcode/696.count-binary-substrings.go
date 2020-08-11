package main

import "fmt"

//696. 计数二进制子串
func countBinarySubstrings(s string) int {
    l := len(s)
    if l <= 1 {
        return 0
    }
    r := 0
    curr := s[0]
    for i := 1; i < l; {
        start, end := i-2, i+1
        if s[i] != curr {
            curr = s[i]
            r++
            for ; start >= 0 && end < l && s[end] == curr && s[start] != curr; start, end = start-1, end+1 {
                r++
            }
        }
        if end < l {
            i = end
        } else {
            break
        }
    }
    return r
}

//func countBinarySubstrings(s string) int {
//    l := len(s)
//    if l <= 1 {
//        return 0
//    }
//    r := 0
//    continuous := make([]int, l)
//    x := 0
//    curr := s[0]
//    for i := 0; i < l; i++ {
//        if curr != s[i] {
//            x++
//            curr = s[i]
//        }
//        continuous[x]++
//    }
//    for i := 0; i < len(continuous)-1; i++ {
//        if continuous[i] == 0 {
//            break
//        }
//        r += min(continuous[i], continuous[i+1])
//    }
//    return r
//}
//
//func min(a, b int) int {
//    if a < b {
//        return a
//    }
//    return b
//}

func main() {
    fmt.Println(countBinarySubstrings("01"), 1)
    fmt.Println(countBinarySubstrings("1010001"), 4)
    fmt.Println(countBinarySubstrings("00110011"), 6)
    fmt.Println(countBinarySubstrings("10101"), 4)
}
