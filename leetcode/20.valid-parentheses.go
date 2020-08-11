package main

import "fmt"

//20. 有效的括号
func isValid(s string) bool {
    l := len(s)
    if l == 0 {
        return true
    }
    m := map[byte]byte{
        ')': '(',
        '}': '{',
        ']': '[',
    }
    stack := []byte{}
    for i := 0; i < l; i++ {
        if v, ok := m[s[i]]; !ok {
            stack = append(stack, s[i])
        } else {
            if len(stack) == 0 {
                return false
            }
            end := stack[len(stack)-1]
            if end != v {
                return false
            }
            stack = stack[:len(stack)-1]
        }
    }
    return len(stack) == 0
}

func main() {
    fmt.Println(isValid(")"), false)
    fmt.Println(isValid("("), false)
    fmt.Println(isValid("()"), true)
    fmt.Println(isValid("()[]{}"), true)
    fmt.Println(isValid("(]"), false)
    fmt.Println(isValid("([)]"), false)
    fmt.Println(isValid("{[]}"), true)
}
