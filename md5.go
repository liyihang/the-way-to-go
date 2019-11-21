package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	md55 := md5.New()
	md55.Write([]byte("zhangsan"))
	res := md55.Sum([]byte(""))
	fmt.Printf("%x\n\n", res)

}
