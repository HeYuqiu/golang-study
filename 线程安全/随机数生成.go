package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var (
	// randPool used to get a rand.Rand and generate a random number thread-safely,
	// which improve the performance of using rand.Rand with a locker
	randPool = &sync.Pool{
		New: func() interface{} {
			return rand.New(rand.NewSource(rand.Int63()))
		},
	}
)

func main() {
	for i := 0; i < 100; i++ {
		fmt.Println(randFloat())
	}
}

func randFloat() float64 {
	rnd := randPool.Get().(*rand.Rand)
	ret := rnd.Float64()
	randPool.Put(rnd)
	return ret
}
