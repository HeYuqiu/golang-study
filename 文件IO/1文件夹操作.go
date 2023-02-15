package main

import (
	"fmt"
	"io/fs"
	"os"
)

func main() {
	fmt.Println(fs.ModeDir)
	fmt.Println(fs.ModeSymlink)
	dir, _ := os.ReadDir("/Users/bytedance/tesmp/pkg")
	fmt.Println("aaa")
	for _, entry := range dir {
		entry.Type()
		fmt.Println(entry.IsDir())

	}
}
