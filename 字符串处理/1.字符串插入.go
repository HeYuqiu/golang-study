package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "https://fdasfw.wf.com://fff"
	b := "hyq"
	replace := strings.Replace(a, "://", fmt.Sprintf("://%s.", b), 1)
	fmt.Println(replace)

	providerID := "volcengine://i-ybqzudcktgkdvbnfzcul"
	if providerID != "" && strings.Contains(providerID, "volcengine://") {
		right := strings.TrimLeft(providerID, "volcengine://")
		fmt.Println(right)
	}
}
