package up

import (
	"fmt"
	"hyq.com/study/down"
)

type Reconciler interface {
	Reconcile()
}

type Animal struct {
	Person Reconciler
}

type MouthUp struct {
}

// 调不同包的方法：大写大写大写大写！！！
func (an *MouthUp) Breathe() {
	fmt.Println("Up breathe")
}
func (an *MouthUp) eat() {
	fmt.Println("Up eat")
}

func NewAnimal() *Animal {
	fmt.Println("NewAnimal")
	mu := &MouthUp{}
	return &Animal{
		Person: down.NewPerson(mu),
	}
}
