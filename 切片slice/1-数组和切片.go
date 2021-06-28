package main

import "fmt"

var slice1 = make([]int, 100)
var slice2 = []int{}

var array = [100]int{}

type person struct {
	name string
}

func main() {
	var arr = make([]person, 0)
	arr = append(arr, person{
		"hyq",
	})
	arr = append(arr, person{
		"zs",
	})
	//arr[0] = person{
	//	"hyq",
	//}
	for _, item := range arr {
		fmt.Println(item)
	}

	println(len(slice1))
	println(cap(slice1))
	slice1 = append(slice1, 1)
	println(len(slice1))
	println(cap(slice1))

}
