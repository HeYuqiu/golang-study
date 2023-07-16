package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	errCh := make(chan string, 5)
	var wg sync.WaitGroup // 等待所有协程完成
	wg.Add(5)

	for i := 0; i < 5; i++ {
		go func(n int) {
			defer wg.Done()
			time.Sleep(1 * time.Second)
			errCh <- fmt.Sprintf("%d", n)
		}(i)
	}

	wg.Wait()
	var sss string

	for i := 0; i < 10; i++ {
		select {
		case sss = <-errCh:
			fmt.Println("str:", sss)
			fmt.Println(len(errCh))
		default:
			fmt.Println("default")
		}
	}

	fmt.Println("ops:")
}
