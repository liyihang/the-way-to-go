package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Student struct {
	Person
	class string
}

func main() {
	student := Student{Person{"zhangsan", 35}, "math"}
	fmt.Printf("%v", student)
}
