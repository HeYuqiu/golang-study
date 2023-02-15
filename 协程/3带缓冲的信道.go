package main

import "fmt"

// 就是设置信道队列的最大长度
// 有缓冲和无缓冲的区别：https://www.shuzhiduo.com/A/KE5Q49MkJL/
func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	//ch <- 3 // 报错
	go func() { ch <- 3 }()
	fmt.Println(ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
