package main

var slice1 = make([]int, 100)
var slice2 = []int{}

var array = [100]int{}

func main() {
	println(len(slice1))
	println(cap(slice1))
	slice1 = append(slice1, 1)
	println(len(slice1))
	println(cap(slice1))
}
