package main

import "fmt"

type Animal1 interface {
	fly()
	run()
}
type Animal2 interface {
	fly()
}
type Bird1 struct {
}

func (bird Bird1) fly() {
	fmt.Print("bird can flying")
}

func (bird Bird1) run() {
	fmt.Print("bird can running")
}
func main() {
	var animal1 Animal1
	var animal2 Animal2
	bird := new(Bird1)
	animal1 = bird
	animal2 = animal1
	animal2.fly()
	animal1.fly()
	animal1.run()

}
