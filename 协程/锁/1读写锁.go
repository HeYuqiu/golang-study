package main

import (
	"fmt"
	"sync"
	"time"
)

var muxTaintGuardMap sync.RWMutex

type NodeReconciler struct {
	taintGuardNodeMap map[string]string
}

func main() {
	r := &NodeReconciler{
		taintGuardNodeMap: make(map[string]string),
	}

	go func() {
		for i := 0; i < 100000; i++ {
			muxTaintGuardMap.RLock()
			_ = r.taintGuardNodeMap["1"]
			muxTaintGuardMap.RUnlock()
		}
	}()

	go func() {
		for i := 0; i < 100000; i++ {
			muxTaintGuardMap.Lock()
			r.taintGuardNodeMap["1"] = "1"
			muxTaintGuardMap.Unlock()
		}
	}()
	time.Sleep(5 * time.Second)
	fmt.Println("done")
}
