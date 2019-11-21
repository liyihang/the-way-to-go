package main

import "fmt"

//func main() {
//	a := int(10)
//	if a > 10 {
//		fmt.Print(10)
//	} else if a > 50 {
//		fmt.Print(50)
//	} else {
//		fmt.Print(1000)
//	}
//}

func main() {
	var class int
	class = 80
	switch class {
	case 80:
		fmt.Print("优秀")
	case 90:
		fmt.Print("厉害")
	default:
		fmt.Print("牛啤")
	}
}
