package main

import "fmt"

func main() {
	fmt.Println("===================== arrayFunc")
	arrays()
	fmt.Println("===================== slice")
	slice()

}

// 数组
func arrays() {
	var test1 [6]int
	fmt.Println("内容：", test1[1:4])
}

// 切片
func slice() {
	// 比数组更强大的序列接口,比数组更常用
	primes := [6]int{2, 3, 5, 1, 7}
	var s []int = primes[1:4]
	fmt.Println(s)
}
