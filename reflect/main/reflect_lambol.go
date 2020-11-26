package main
import (
	"fmt"
	"reflect"
)
// 定义结构体
type User1 struct {
	Id   int
	Name string
	Age  int
}
// 匿名字段
type Boy struct {
	User1
	Addr string
}
func main() {
	m := Boy{User1{1, "xiaobai", 20}, "sz"}
	t := reflect.TypeOf(m)
	fmt.Println(t)
	// Anonymous：匿名
	fmt.Printf("%#v\n", t.Field(0))
	// 值信息
	fmt.Printf("%#v\n", reflect.ValueOf(m).Field(0))
}