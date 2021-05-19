package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// 类似java中类的toString方法
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	fmt.Println(a)
}
