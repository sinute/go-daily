package main

import (
    "fmt"
)

//43. 字符串相乘
func multiply(num1 string, num2 string) string {
    carry := 0
    r := make([]byte, 0)
    l1 := len(num1)
    l2 := len(num2)
    l := l1 + l2
    values := make([][]byte, l)
    for j := l2 - 1; j >= 0; j-- {
        for i := l1 - 1; i >= 0; i-- {
            v := (num1[i]-'0')*(num2[j]-'0') + byte(carry)
            carry = int(v / 10)
            v %= 10
            values[l1-1-i+l2-1-j] = append(values[l1-1-i+l2-1-j], v)
        }
        if carry > 0 {
            values[l1+l2-1-j] = append(values[l1+l2-1-j], byte(carry))
            carry = 0
        }
    }

    carry = 0
    for i := 0; i < l; i++ {
        v := carry
        for j := len(values[i]) - 1; j >= 0; j-- {
            v += int(values[i][j])
        }
        carry = v / 10
        v %= 10
        r = append([]byte{byte(v) + '0'}, r...)
    }

    for len(r) > 1 && r[0] == '0' {
        r = r[1:]
    }
    return string(r)
}

func main() {
    fmt.Println(multiply("6964594125027226699998128707627435007621143975736747759751", "333791918659904899647584436187733004125181273682766434"), "2324725235680339589575434145174345450376468286967007130548581359508676923461769510883584014763890133705678997934")
    fmt.Println(multiply("0", "1233"), "0")
    fmt.Println(multiply("9", "9"), "81")
    fmt.Println(multiply("99", "99"), "9801")
    fmt.Println(multiply("999", "999"), "998001")
    fmt.Println(multiply("2", "3"), "6")
    fmt.Println(multiply("123", "456"), "56088")
}
