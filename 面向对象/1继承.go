package main

type A struct {
}
type B struct {
	A //B is-a A
}

func save(A) {
	//do something
}
func main() {
	b := B{}
	save(b) // 类型作为方法参数时，没有继承传递，
}
