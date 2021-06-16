package main

import "hyq.com/study/up"

func main() {
	animal := up.NewAnimal()
	animal.Person.Reconcile()
}
