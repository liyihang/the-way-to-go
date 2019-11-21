package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	timeout := make(chan int, 1)
	go func() {
		time.Sleep(time.Second)
		timeout <- 1
	}()
	//time.Sleep(time.Second)
	select {
	case <-ch:
		fmt.Println("ch")
	case <-timeout:
		fmt.Println("timeout")
	}
	fmt.Println("end of code ")
}
