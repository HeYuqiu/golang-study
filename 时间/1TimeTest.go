package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now().Add(time.Duration(int64(360)) * time.Minute))
}
