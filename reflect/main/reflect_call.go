package main
import (
	"fmt"
	"reflect"
)
// 定义结构体
type User5 struct {
	Id   int
	Name string
	Age  int
}
func (u User5) Hello(name string) {
	fmt.Println("Hello：", name)
}
func main() {
	u := User5{1, "xiaobai", 20}
	v := reflect.ValueOf(u)
	fmt.Println(v)
	// 获取方法
	m := v.MethodByName("Hello")
	// 构建一些参数
	args := []reflect.Value{reflect.ValueOf("6666")}
	// 没参数的情况下：var args2 []reflect.Value
	// 调用方法，需要传入方法的参数
	m.Call(args)
}