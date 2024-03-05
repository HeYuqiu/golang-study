package main

import "fmt"

func main() {
	// 内置的 strconv 包提供了数字的解析功能
	// float64转int
	var a float64
	a = 3.6
	b := int32(a)

	fmt.Println(b)
}
