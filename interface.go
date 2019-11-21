package main

import "fmt"

type Animal interface {
	fly()
	run()
}
type Bird struct {
}

func (bird Bird) fly() {
	fmt.Print("bird can flying")
}

func (bird Bird) run() {
	fmt.Print("bird can running")
}
func main() {
	var animal Animal
	bird := new(Bird)
	animal = bird
	animal.fly()
	animal.run()
}
