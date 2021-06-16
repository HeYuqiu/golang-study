package down

import "fmt"

type Mouse interface {
	Breathe()
}

func NewPerson(mouse Mouse) *Person {
	fmt.Println("NewPerson")
	return &Person{
		Mouse: mouse,
	}
}

type Person struct {
	Mouse Mouse
}

func (p *Person) Reconcile() {
	fmt.Println("down reconcile")
	p.Mouse.Breathe()
}
