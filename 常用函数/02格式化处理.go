package main

import (
	"encoding/json"
	"fmt"
)

type point struct {
	x, y int
	m    *man
}
type man struct {
	name string
}

func main() {
	p := point{
		x: 1,
		y: 2,
		m: &man{
			name: "hyq",
		},
	}
	marshal, _ := json.Marshal(p)
	fmt.Println(string(marshal))
	fmt.Printf("%v \n", p)
	fmt.Printf("%+v \n", p) // 包含字段名
	fmt.Printf("%#v \n", p) // 值的语法表示
	fmt.Printf("%T \n", p)  // 值的类型
}
