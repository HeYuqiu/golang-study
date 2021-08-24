package main

import "fmt"

type person struct {
	name *string
}

func strr(str *string) {
	fmt.Println(str)
}

func main() {
	//name := "name"
	p := &person{
		//name: &name,
	}
	//fmt.Println("hyq" + p.name)
	fmt.Println("hyq" + *p.name)

	str := "fff"
	//strr(str) // 报错
	//strr("fff") // 报错
	//strr(&"fff") // 报错
	strr(&str)
}
