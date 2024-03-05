package main

import "fmt"

type People struct {
	name string
}
type Eat interface {
	eatFood()
}

func (p People) eatFood() {
	fmt.Printf("%s eat food", p.name)
}

var DefaultPerson Eat = &People{
	name: "hyqtest",
}

func main() {
	people := DefaultPerson.(*People)
	fmt.Println(people.name)
	people.eatFood()
	//DefaultPerson.eatFood()
}
