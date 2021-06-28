package main

import "fmt"

func main() {
	v1 := [6]int{1, 2, 3, 4, 5, 6}
	v2 := v1[0:5]
	v3 := v1[1:4]
	fmt.Println(v2)
	fmt.Println(v3)
	v3[0] = 0
	fmt.Println(v2)
	fmt.Println(v3)
}
