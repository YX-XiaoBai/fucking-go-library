package main

import (
	"fmt"
	"io/ioutil"
)

func wr() {
	err := ioutil.WriteFile("./ioutil.txt", []byte("www.baidu.com"), 0666)
	if err != nil {
		fmt.Println("E1",err)
		return
	}
}

func re() {
	content, err := ioutil.ReadFile("./ioutil.txt")
	if err != nil {
		fmt.Println("E2",err)
	}
	fmt.Println(string(content))
}

func main() {
	wr()
	re()
}