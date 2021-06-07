package main

import "fmt"

func judgeType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("类型：%T，值：%v\n", v, v)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}
func main() {
	judgeType(21)
	judgeType("hello")
	judgeType(true)
}
