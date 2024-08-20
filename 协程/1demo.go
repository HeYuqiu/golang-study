package main

import (
	"fmt"
	"sync"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	// 5. 全局通用限流（异步）
	go func() {
		fmt.Println("1 start")
		defer wg.Done()
		defer func() {
			if p := recover(); p != nil {
				fmt.Println("recover")
			}
		}()
		time.Sleep(1 * time.Second)
		fmt.Println("1 end")
	}()
	// 6. 大资源限流（异步）
	go func() {
		fmt.Println("2 start")
		defer wg.Done()
		defer func() {
			if p := recover(); p != nil {
				fmt.Println("recover")
			}
		}()
		time.Sleep(2 * time.Second)
		fmt.Println("2 end")
	}()

	// main 函数等待异步线程？
	fmt.Println("wait start")
	wg.Wait()
	fmt.Println("wait end")
}
