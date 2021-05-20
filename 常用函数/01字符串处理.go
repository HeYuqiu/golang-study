package main

import (
	"fmt"
	"strings"
)

var p = fmt.Println

func main() {
	p("Contains:", strings.Contains("test", "es"))
	p("Count:", strings.Count("test", "t"))
	p("Split:", strings.Split("a-b-c-d-e", "-"))
}
