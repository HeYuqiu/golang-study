package main

import (
	"fmt"
	"time"
)

type person struct {
	name string
	age  int32
}

func mod(p *person, v int32) {
	p.age = v
}

func main() {
	p := person{
		"hyq", 22,
	}
	go mod(&p, 31)
	go mod(&p, 32)
	go mod(&p, 33)
	time.Sleep(10)
	fmt.Println(p)
}
