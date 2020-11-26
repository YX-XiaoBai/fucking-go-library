package main
import (
	"fmt"
	"io/ioutil"
	"math/rand"
)
type Person1 struct {
	Name string
	Age  int
	Sex  string
}
// 二进制写出
func writerJson(filename string) (err error) {
	var persons []*Person1
	// 假数据
	for i := 0; i < 10; i++ {
		p := &Person1{
			Name: fmt.Sprintf("name%d", i),
			Age:  rand.Intn(100),
			Sex:  "male",
		}
		persons = append(persons, p)
	}
	// 二进制json序列化
	data, err := msgpack.Marshal(persons)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = ioutil.WriteFile(filename, data, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}
// 二进制读取
func readJson(filename string) (err error) {
	var persons []*Person
	// 读文件
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 反序列化
	err = msgpack.Unmarshal(data, &persons)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range persons {
		fmt.Printf("%#v\n", v)
	}
	return
}
func main() {
	//err := writerJson("D:/person.dat")
	//if err != nil {
	// fmt.Println(err)
	// return
	//}
	err := readJson("D:/person.dat")
	if err != nil {
		fmt.Println(err)
		return
	}
}