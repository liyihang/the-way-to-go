package main

import "fmt"

//func main() {
//	x:=[10]int{1,2,3,4,5,6,7,8,9,10}
//	y:=x[1:3]
//	fmt.Print(y)
//}
func main() {
	x := make([]int, 3, 5) // 创建cap为5 length为3的slice
	x = append(x, 4, 5, 6, 77)
	//fmt.Print(x)
	fmt.Print(cap(x)) //cap 10  动态增加了
}
