package main

import "fmt"

//func main() {
//	var info map[string]string
//	//map 必须分配内存 否则报错
//	info = make(map[string]string)
//	info["name"] = "jim"
//	info["sex"] = "boy"
//	for k, v := range info {
//		fmt.Print(k, v)
//	}
//	//fmt.Printf("%v",info)
//}
func main() {
	m := map[string]string{
		"name": "jj",
		"age":  "20",
	}
	name := m["name"]
	fmt.Print(name)

}
