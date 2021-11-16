package main

import "fmt"

func main() {
	user := map[string]string{"k": "v", "h": "y"}
	for s, s2 := range user {
		fmt.Println(s, s2)
	}
}
