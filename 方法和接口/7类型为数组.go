package main

import "fmt"

type IPAddr [4]byte

// TODO: 给 IPAddr 添加一个 "String() string" 方法
//func (ip IPAddr) String() string{
//	return string(ip[1])
//}

func main() {
	//hosts := map[string]IPAddr{
	//	"hyq":{127,0,0,1},
	//}

	var arr = IPAddr{1, 2, 3, 4}
	fmt.Println(arr)
	fmt.Println("fff:", string(arr[:]))

}
