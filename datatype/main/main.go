package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	//Name string
	//Hobby string

	Age       int    `json:"age,string"`
	Name      string `json:"name"`
	Niubility bool   `json:"niubility"`
}

func main() {
	// json
	//p := Person{"xiaobai", "girl"}
	//b, err := json.Marshal(p)
	//if err != nil {
	//	fmt.Println("json err", err)
	//}
	//fmt.Println(string(b))
	//// format output
	//b, err =json.MarshalIndent(p, "","	")
	//if err != nil {
	//	fmt.Println("json err", err)
	//}
	//fmt.Println(string(b))

	//通过map生成json
	//student := make(map[string]interface{})
	//student["name"] = "Xiaobai"
	//student["age"] = 18
	//student["sex"] = "man"
	//b, err := json.Marshal(student)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(string(b))

	// 解析到结构体
	//b := []byte(`{"age":"18","name":"Xiaobai","niubility":false}`)
	//var p Person
	//err := json.Unmarshal(b, &p)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(p)

	// 解析到interface
	// int和float64都当float64
	b := []byte(`{"age":1.6,"name":"Xiaobai","niubility":false}`)
	var i interface{}
	err := json.Unmarshal(b, &i)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(i)
	m := i.(map[string]interface{})
	for k, v := range m {
		switch vv := v.(type) {
		case float64:
			fmt.Println(k, "is float64", vv)
		case string:
			fmt.Println(k, "is string", vv)
		default:
			fmt.Println("other")
		}
	}
}
