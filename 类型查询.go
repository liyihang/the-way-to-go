package main

import "fmt"

func main() {
	var v1 interface{}
	v1 = "3.14"
	//if v,ok := v1.(float64);ok{
	//	fmt.Print(v,ok)
	//}

	switch v1.(type) {
	case int:
	case float64:
		fmt.Print("float64")
	case string:
		fmt.Printf("string")
	}
}
