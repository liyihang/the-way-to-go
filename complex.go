package main

import "fmt"

// complex
//func main() {
//	var x complex64 = 6 + 2i
//
//	fmt.Print(x*x)
//}
//func main() {
//	var x [10] int
//	x[0] = 1
//	x[9] = 10
//	fmt.Printf("%v",x)
//}

func main() {
	x := [10]int8{1, 2, 3, 4, 5, 6}
	fmt.Printf("%v", len(x))
}
