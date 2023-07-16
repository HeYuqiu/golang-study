package main

import (
	"fmt"
	"time"
)

func main() {
	A("hyq")

	time.Sleep(5 * time.Second)
}

func A(cmd string) {
	fmt.Printf("A1 %s \n", cmd)
	aa := "aa"
	go func() {
		aa = B(cmd)
	}()
	fmt.Println(aa)
	fmt.Printf("A2 %s \n", cmd)
	time.Sleep(1 * time.Second)
	cmd = "change"
	fmt.Printf("A3 %s \n", cmd)
	time.Sleep(3 * time.Second)
	fmt.Println(aa)

}

func B(cmdb string) string {
	fmt.Printf("B1 %s \n", cmdb)
	time.Sleep(3 * time.Second)
	fmt.Printf("B2 %s \n", cmdb)
	return "bbb"
}
