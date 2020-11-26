package main

import (
	"fmt"
	"reflect"
)

type User3 struct {
	Id int
	Name string
	Age int
}

func SetValue(o interface{}) {
	v := reflect.ValueOf(o)
	fmt.Println(v)
	// 获取指针指向的元素
	v = v.Elem()
	fmt.Println(v)
	// 取字段
	f := v.FieldByName("Name")
	fmt.Println(f)
	if f.Kind() == reflect.String {
		f.SetString("xiaobai")
	}
}

func main() {
	u := User3{1, "xiaohei", 20}
	SetValue(&u)
	fmt.Println(u)
}