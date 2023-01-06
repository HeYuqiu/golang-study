package main

import (
	"fmt"
	"regexp"
)

func main() {
	context := "UE9||{{.size10}}|{{.size20}}|a{{.namespace}}aPQ"
	pwd1, _ := regexp.Compile("{{\\.size.*?}}")
	context1 := pwd1.ReplaceAllString(context, "111111")
	findString := pwd1.FindString(context)
	allString := pwd1.FindAllString(context, 20)
	fmt.Println(context1, findString, allString)
}
