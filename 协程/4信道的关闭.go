package main

import (
	"fmt"
	"time"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	fmt.Println(<-c)
	for i := range c {
		fmt.Println(i)
	}
	time.Sleep(10)
	// 信道已经关闭了，再获取直接返回0，不会阻塞造成死锁
	fmt.Println(<-c)
}
