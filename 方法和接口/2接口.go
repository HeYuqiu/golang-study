package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
	test(s string)
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (f MyFloat) test(s string) {
	fmt.Println(s)
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	//fmt.Println(a.Abs()) // 报错
	a = f
	fmt.Println(a.Abs())
	a.test("fff")
}
