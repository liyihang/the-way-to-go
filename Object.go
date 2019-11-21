package main

import "fmt"

type Integer struct {
	value int
}

func (a Integer) compare(b Integer) bool {
	return a.value > b.value
}

func main() {
	a := Integer{1}
	b := Integer{2}
	fmt.Printf("%v", a.compare(b))

}
