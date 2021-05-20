package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	d1 := []byte("hello\ngo\n")
	ioutil.WriteFile("dat1", d1, 0644)
	f, err := os.Create("dat1")
	defer f.Close() // 打开文件后，习惯立即使用 defer 调用文件的 Close操作。类似java的finally
	check(err)
	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	fmt.Printf("wrote %d bytes\n", n2)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
