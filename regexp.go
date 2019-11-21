package main

import (
	"fmt"
	"regexp"
)

func main() {
	isFit, _ := regexp.MatchString("[a-zA-Z]{3}", "wqeheql")
	fmt.Printf("%v", isFit)
}
