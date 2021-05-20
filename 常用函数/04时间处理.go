package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println
	now := time.Now()

	p(now)
	p(now.Day())
	p(now.Weekday())
	p(now.Format("2006-01-02T15:04:05"))

}
