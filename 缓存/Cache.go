package main

import (
	"fmt"
	"sync"
	"time"
)

var globalMap sync.Map

func Set(key string, data interface{}, timeout int) {
	globalMap.Store(key, data)
	time.AfterFunc(time.Second*time.Duration(timeout), func() {
		globalMap.Delete(key)
	})
}

func Get(key string) (interface{}, bool) {
	return globalMap.Load(key)
}

func main() {
	Set("hyq", 1, 5)
	Set("hyq", "13213", 5)
	time.Sleep(time.Second * 2)
	get, b := Get("hyq")
	go func() {
		get, b := Get("hyq")
		if b {
			fmt.Printf("runtime %s ", get)
		}
	}()

	if b {
		fmt.Println(get)
	}

	time.Sleep(time.Second * 2)

}
