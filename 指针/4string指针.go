package main

import "fmt"

func strr(str *string) {
	fmt.Println(str)
}

func main() {
	str := "fff"
	//strr(str) // 报错
	//strr("fff") // 报错
	//strr(&"fff") // 报错
	strr(&str)
}
