package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := make([]string, 0, 3)
	arr = append(arr, "hyq")
	arr = append(arr, "jyq")
	arr = append(arr, "ayq")
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})
	fmt.Println(arr)
}
