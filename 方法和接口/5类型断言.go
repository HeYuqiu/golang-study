package main

import "fmt"

func main() {
	var i interface{} = "hello"
	s := i.(string) // 类型断言提供了访问接口底层类型具体值的能力
	fmt.Println(s)
	s, ok := i.(string)
	fmt.Println(s, ok)
	f, ok := i.(float64)
	fmt.Println(f, ok)
	i = 10.1
	ff, ok := i.(float64)
	if ok {
		fmt.Println(ff)
	}
	i = "hyq"
	f = i.(float64) // 报错(panic)
	fmt.Println(f)
}
