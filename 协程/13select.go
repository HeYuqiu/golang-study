package main

import (
	"fmt"
	"time"
)

/**
处理超时
使用 channel 时需要小心，比如对于下面的简单用法：

i := <-ch
碰到永远没有往 ch 中写入数据的情况，那么这个读取动作将永远也无法从 ch 中读取到数据，导致的结果就是整个 goroutine 永远阻塞并且没有挽回的机会。如果 channel 只是被同一个开发者使用，那样出问题的可能性还低一些。但如果一旦对外公开，就必须考虑到最差情况并对程序进行维护。

Golang 没有提供直接的超时处理机制，但可以利用 select 机制变通地解决。因为 select 的特点是只要其中一个 case 已经完成，程序就会继续往下执行，而不会考虑其它的 case。基于此特性我们来实现一个 channel 的超时机制：
*/

func main() {
	ch := make(chan int)
	// 首先实现并执行一个匿名的超时等待函数
	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(1e9) // 等待 1 秒
		timeout <- true
	}()
	// 然后把 timeout 这个 channel 利用起来
	select {
	case <-ch:
		// 从 ch 中读取到数据
	case <-timeout:
		// 一直没有从 ch 中读取到数据，但从 timeout 中读取到了数据
		fmt.Println("Timeout occurred.")
	}
}
