package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string `json:"name1" db:"name2"`
}

func main() {
	var s Student
	v := reflect.ValueOf(&s)
	t := v.Type()
	f := t.Elem().Field(0)
	fmt.Println(f.Tag.Get("json"))
	fmt.Println(f.Tag.Get("db"))
}