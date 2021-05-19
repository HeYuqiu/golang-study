package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

type MyFloat float64

type Vertex struct {
	X, Y float64
}

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}
	//fmt.Println(a.Abs()) // 报错
	a = f
	fmt.Println(a.Abs())
	//a = v //报错：实现Abs()的是 *Vertex 而不是 Vertex
	a = &v
	fmt.Println(a.Abs())
}
