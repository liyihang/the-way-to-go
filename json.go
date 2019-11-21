package main

import (
	"encoding/json"
	"fmt"
)

type Students struct {
	Name string `json:"student_name"`
	Age  int
}

func main() {
	//对数组进行json编码
	x := [5]int{1, 2, 3, 4, 5}
	s, err := json.Marshal(x)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(s))
	//map json 编码
	m := make(map[string]float64)
	m["lisi"] = 130.12
	s2, err2 := json.Marshal(m)
	if err2 != nil {
		panic(err2)
	}
	fmt.Println(string(s2))
	//	对对象进行json编码
	student := Students{"zhangsan", 22}
	s3, err3 := json.Marshal(student)
	if err3 != nil {
		panic(err3)
	}
	fmt.Println(string(s3))
	//	对s3 进行解码
	var s4 interface{}
	json.Unmarshal([]byte(s3), &s4)
	fmt.Printf("%v", s4)
}
