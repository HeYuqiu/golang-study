package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {
	var i I
	//i.M() // 这里肯定要报错，没有实现类
	var t *T
	i = t
	describe(i)
	i.M() // 如果是java就报空指针了，但是go不会报，会继续执行方法

	i = &T{"hyq"}
	describe(i)
	i.M()

	// 空接口：类似java的Object，因为任何类型都至少实现了空接口
	var nilInterface interface{}
	describe1(nilInterface)

}

func describe(i I) {
	fmt.Printf("%v, %T \n", i, i)
}
func describe1(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
