package main

import (
	"fmt"
	"sync"
	"time"
)

var mMap map[int]int

var lock = sync.Mutex{}

func main() {
	mMap = make(map[int]int)

	for i := 0; i < 500; i++ {
		go func() {
			lock.Lock()
			defer lock.Unlock()
			mMap[0] = i
		}()
	}
	time.Sleep(10)
	fmt.Println(mMap[0])
}
