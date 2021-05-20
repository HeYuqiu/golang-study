package main

import "fmt"

type point struct {
	x, y int
}

func main() {

	p := point{1, 2}
	fmt.Printf("%v \n", p)
	fmt.Printf("%+v \n", p) // 包含字段名
	fmt.Printf("%#v \n", p) // 值的语法表示
	fmt.Printf("%T \n", p)  // 值的类型
}
