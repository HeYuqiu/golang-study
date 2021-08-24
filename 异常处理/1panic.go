package main

import "fmt"

// panic 表示进程内的错误。panic 的原因来自于代码的逻辑 bug，比如强制类型转换失败，比如数组越界。这个代表了程序员的责任不到位，导致了程序的panic。 error 代表进程外的错误。比如输入符合预期。比如访问外部的服务失败。这些都不是程序员可以设计控制的。这些情况的错误处理是业务逻辑的一部分。
// Go中可以抛出一个panic的异常，然后在defer中通过recover捕获这个异常，然后正常处理。
func main() {

	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		fmt.Println("c")
		if err := recover(); err != nil {
			fmt.Println(err) // 这里的err其实就是panic传入的内容，55
		}
		fmt.Println("d")
	}()

	f()
}

func f() {
	fmt.Println("a")
	panic(55)
	fmt.Println("b")
	fmt.Println("f")
}
