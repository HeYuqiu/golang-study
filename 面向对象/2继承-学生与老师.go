package main

import "fmt"

type person struct {
	name string
	age  int
}

//如果一个struct嵌套了另一个匿名结构体，就可以直接访问匿名结构体的字段或方法，从而实现继承
type student struct {
	person //匿名字段，struct
	mobile string
}

//如果一个struct嵌套了另一个【有名】的结构体，叫做组合
type teacher struct {
	p      person //有名字段，struct
	mobile string
}

func (p *person) run() {
	fmt.Println(p.name, " run")
}

func (p *person) reading() {
	fmt.Println(p.name, " reading")
}

func (s *student) reading() {
	fmt.Println(s.name, " reading")
}

func main() {
	p := person{"zhangsan", 22}
	s := student{person{"lisi", 20}, "000"}
	t := teacher{person{"wangwu", 25}, "000"}

	fmt.Println(s.name)   //访问【匿名】结构体的字段
	fmt.Println(t.p.name) //访问【有名】结构体的字段。不是继承，需要指定结构体
	p.run()
	s.run() // 类型作为调用者时，有继承传递
}
