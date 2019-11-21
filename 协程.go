package main

import (
	"fmt"
	"strconv"
	"time"
)

var ch chan int

func main() {
	ch = make(chan int)
	go func() {
		for i := 1; i < 100; i++ {
			if i == 10 {
				//runtime.Gosched()  //主动让出cpu
				<-ch
			}
			fmt.Println("routine 1:" + strconv.Itoa(i))
		}
	}()
	go func() {
		for i := 100; i < 200; i++ {
			fmt.Println("routine 2:" + strconv.Itoa(i))
		}
		ch <- 1
	}()
	time.Sleep(time.Second)
}
