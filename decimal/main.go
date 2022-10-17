package main

import (
	"fmt"

	jsoniter "github.com/json-iterator/go"
	"github.com/shopspring/decimal"
)

type Test struct {
	a decimal.Decimal `json:"a"`
}

func main() {
	x := decimal.NewFromFloat(123456789.9)
	fmt.Println(x.StringFixed(8))
	fmt.Printf("%.8f\n", 123456789.9)
	return

	decimal.DivisionPrecision = 3
	a := decimal.NewFromFloat(1.0)
	fmt.Println(a.Equal(decimal.Zero))
	t := Test{
		a: a,
	}
	v, err := jsoniter.Marshal(t)
	fmt.Println(string(v), err)
	return

	z := a

	a = a.Sub(a)
	b := decimal.NewFromFloat(2)
	c := a.Mul(b)
	fmt.Printf("%s\n", a)
	fmt.Printf("%s\n", z)
	fmt.Println(b)
	fmt.Println(c)
	// fmt.Println(decimal.NewFromFloat(1.123456789).RoundCeil(2).String()) // output: "-1.5"
}
