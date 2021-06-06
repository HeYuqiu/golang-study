package main

import (
	"fmt"
	"time"
)

// 同时监听多个信道，哪个消息来了就处理哪个，如果同时来了，就随机处理

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()
	for i := 0; i < 2; i++ {
		fmt.Println(i)
		select { // 如果不用default，会一直阻塞
		case msg1 := <-c1:
			fmt.Println("received1", msg1)
		case msg2 := <-c2:
			fmt.Println("received2", msg2)
		}
	}
}
