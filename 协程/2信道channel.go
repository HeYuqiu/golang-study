package main

import "fmt"

// 信道主要用户协程之间数据传递？

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // 从c取值，信道相当于一个队列，因为有两个协程，所以有两个c的值，获取时会阻塞
	fmt.Println(x, y, x+y)
}
