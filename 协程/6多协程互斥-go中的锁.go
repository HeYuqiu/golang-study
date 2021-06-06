package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// 不带锁
	//num := 0
	//for i := 0; i < 1000; i++ {
	//	go func() { num = num + 1 }()
	//}
	//time.Sleep(time.Second)
	//fmt.Println(num)

	num := 0
	var mutex = sync.Mutex{}
	for i := 0; i < 1000; i++ {
		go func() {
			mutex.Lock()
			defer mutex.Unlock()
			num = num + 1
		}()
	}
	time.Sleep(time.Second)
	fmt.Println(num)
}
