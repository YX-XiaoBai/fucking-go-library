package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

// 抽取单个server对象
type Server struct {
	ServerName string `xml:"serverName"`
	ServerIP   string `xml:"serverIP"`
}
type Servers struct {
	Name    xml.Name `xml:"servers"`
	Version int      `xml:"version"`
	Servers []Server `xml:"server"`
}

func main() {
	data, err := ioutil.ReadFile("./test.xml")
	if err != nil {
		fmt.Println(err)
		return
	}
	var servers Servers
	err = xml.Unmarshal(data, &servers)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("xml: %#v\n", servers)
}