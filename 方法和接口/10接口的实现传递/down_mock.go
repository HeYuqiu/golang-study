package main

import "fmt"

type MockPerson struct {
}

func (m *MockPerson) breathe() {
	fmt.Println("mock breathe")
}
