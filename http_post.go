package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	res, err := http.Post("http://www.baidu.com", "application/x-www-form-urlencoded", strings.NewReader("id=1"))
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}
