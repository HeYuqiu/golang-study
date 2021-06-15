package main

import "fmt"

type DB interface {
	Get(key string) (int, error)
}

func GetFromDB(db DB, key string) int {
	if value, err := db.Get(key); err == nil {
		fmt.Println("success" + key)
		return value
	}
	fmt.Println("fail" + key)
	return -1
}
