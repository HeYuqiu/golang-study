package main

import "fmt"

type Animal struct {
	*Person
}

type Person struct {
}

func (p *Person) Do() {
	fmt.Println("person do")
}

//
//func (p *Animal) Do()  {
//	fmt.Println("Animal do")
//}

func main() {
	an := &Animal{
		Person: &Person{},
	}
	an.Do()
}
