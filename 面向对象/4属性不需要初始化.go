package main

import "fmt"

type Animal struct {
	Person *Person
}

type Person struct {
}

func (p *Person) Do() {
	fmt.Println("person do")
}

func main() {
	an := &Animal{
		//Person: &Person{},
	}
	an.Person.Do()
}
