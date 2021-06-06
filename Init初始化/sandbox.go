package main

import "fmt"

// 最先执行
var _ int64 = s()

// 第二执行
func init() {
	fmt.Println("init in sandbox.go")
}

func s() int64 {
	fmt.Println("calling s() in sandbox.go")
	return 1
}

// 最后执行
func main() {
	fmt.Println("main")
}
