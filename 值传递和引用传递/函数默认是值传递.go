package main

import "fmt"

// golang只有值传递没有引用传递
// 就算传的指针它也是值，只是这个值是指针，不像java
func main() {
	var s []int
	for i := 0; i < 3; i++ {
		s = append(s, i)
	}
	modifySlice(s)
	fmt.Println(s)
}

// 特殊情况：传的值是[]int的指针，如果声明为*[]int就表示传的是指针的指针；
func modifySlice(s []int) {
	s = append(s, 2048) //append后产生的是新引用，而不是原始引用，slice是按值传递的，
	s[0] = 1024         // 修改成功
	fmt.Println(s)
}

func modifySlice1(s []int) {
	s = append(s, 2048)
	s = append(s, 4096)
	s[0] = 1024 // 修改失败，因为发生了扩容，修改只发生在新内存中
	fmt.Println(s)
}
