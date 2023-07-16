package main

type person struct {
	name string
	age  int
	mm   map[string]string
}

func main() {

	pp := map[string]*person{}
	pp["111"] = &person{
		name: "ffff",
		age:  0,
	}
	delete(pp, "222")
	delete(pp, "111")
	delete(pp, "111")

	pp["111"].mm["aaa"] = "ffff"
	p, ok := pp["1222"]
	if p == nil {
		println(p)
	}
	println(p)

	kv := map[string]string{"k1": "v1", "k2": "v2"}
	kv["hyq"] = ""
	kv["hyq"] = ""
	kv["hyq1"] = ""
	kv["k1"] = "222"
	s := kv["aaaa"]
	if s == "" {
		println("aaa")
	}
	println(s)
	value, ok := kv["k1"]
	if ok {
		println(value)
	} else {
		println("ss")
	}
	println("end")
}
