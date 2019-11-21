package main

import (
	"fmt"
	"time"
)

func test_channel() {
	ch := make(chan int, 1)
	ch <- 1
	fmt.Println("file run")
}
func main() {
	go test_channel()
	time.Sleep(2 * time.Second)
	fmt.Println("run test")
}
