package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	start := time.Now()
	_ = test()
	log.Printf("main time cost = %v \n", time.Since(start))
}

func test() error {
	start := time.Now()
	channel := make(chan error, 3)
	go func() {
		fmt.Println("method1 ...")
		time.Sleep(3 * time.Second)
		//channel <- errors.New("method1 err")
		channel <- nil
	}()
	go func() {
		fmt.Println("method2 ...")
		time.Sleep(4 * time.Second)
		channel <- nil
	}()
	go func() {
		fmt.Println("method3 ...")
		time.Sleep(5 * time.Second)
		//channel <- errors.New("method3 err")
		channel <- nil
	}()

	for i := 0; i < 3; {
		err := <-channel
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
		i++
	}
	log.Printf("time cost = %v \n", time.Since(start))
	fmt.Println("end ...")
	return nil
}
