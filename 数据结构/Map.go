package main

func main() {
	kv := map[string]string{"k1": "v1", "k2": "v2"}
	kv["hyq"] = ""
	kv["hyq"] = ""
	kv["hyq1"] = ""
	kv["k1"] = "222"
	value, ok := kv["k1"]
	if ok {
		println(value)
	} else {
		println("ss")
	}
	println("end")
}
