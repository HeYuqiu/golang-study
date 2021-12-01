package main

import "fmt"

func main() {
	result := make(map[int]*string)
	strArr := []string{"1", "2", "3", "4"}
	for i, s := range strArr {
		result[i] = &s
	}
	fmt.Println(result)
}
