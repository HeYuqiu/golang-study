package main

import (
	"fmt"
	"hyq.com/study/pkgtest"
)

func main() {
	a := 1
	test(&a)
	fmt.Println(a)

	mm := make(map[string]string)
	mm["a"] = "a"
	mm["b"] = "b"
	delete(mm, "a")
	delete(mm, "c")
	pkgtest.PkgMethod("hyq")
}

func test(int2 *int) {
	int3 := 10
	int2 = &int3
}
