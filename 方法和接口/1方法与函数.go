package main

import (
	"fmt"
	"math"
)

// Go的结构体，类似java的实体类
type Vertex struct {
	X, Y float64
}

// 方法，Go没有类，但我们可以为结构体定义方法
func (v Vertex) Abs() float64 {
	v.X = 10
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// 函数
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{
		X: 3,
		Y: 4,
	}
	fmt.Println(Abs(v))
	fmt.Println(v.Abs())
}
