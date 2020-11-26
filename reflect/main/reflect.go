package main

import (
	"fmt"
	"reflect"
)

func reflect_type(a interface{}) {
	t := reflect.TypeOf(a)
	fmt.Println("Type is:", t)
	k := t.Kind()
	fmt.Println(k)
	switch k {
	case reflect.Float64:
		fmt.Println("a is float64\n")
	case reflect.String:
		fmt.Println("string\n")
	}
}

func reflect_value(a interface{}) {
	v := reflect.ValueOf(a)
	fmt.Println("Value is:", v)
	k := v.Kind()
	fmt.Println(k)
	switch k {
	case reflect.Float64:
		fmt.Println("a is", v.Float())
	}
}

func reflect_set_value(a interface{}){
	v := reflect.ValueOf(a)
	k := v.Kind()
	switch k {
	case reflect.Float64:
		v.SetFloat(6.8)
		fmt.Println("a is", v.Float())
	case reflect.Ptr:
		v.Elem().SetFloat(8.8)
		fmt.Println("a is", v.Elem().Float())
		fmt.Println(v.Pointer())
	}
}

func main() {
	var x float64 = 3.4
	reflect_set_value(&x)
	fmt.Println("main:", x)
	//reflect_type(x)
	//reflect_value(x)
}