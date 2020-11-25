package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type UserInfo struct {
	Name   string
	Gender string
	Age    int
}

//func sayHello(w http.ResponseWriter, r *http.Request) {
//	// 解析指定文件生成模板对象
//	//tmpl, err := template.ParseFiles("./hello.html")
//	//if err != nil {
//	//	fmt.Printf("create template failed, err:", err)
//	//	return
//	//}
//	//tmpl.Execute(w, "Xiaobai")
//
//
//	htmlByte, err := ioutil.ReadFile("./hello.html")
//	if err != nil {
//		fmt.Println("read html failed, err:", err)
//		return
//	}
//	// 自定义一个kua的模板函数
//	kua := func(arg string) (string, error) {
//		return arg + "真帅", nil
//	}
//	// 采用链式操作在Parse之前调用Funcs添加自定义的kua函数
//	tmpl, err := template.New("hello").Funcs(template.FuncMap{"kua": kua}).Parse(string(htmlByte))
//	if err != nil {
//		fmt.Println("create template failed, err:", err)
//		return
//	}
//
//	user := UserInfo{
//		Name: "xiaobai",
//		Gender: "boy",
//		Age: 18,
//	}
//	tmpl.Execute(w, user)
//}

func tmplDemo(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./t.html", "./ul.html")
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}
	user := UserInfo{
		Name: "xiaobai",
		Gender: "boy",
		Age: 18,
	}
	tmpl.Execute(w, user)
}

func main() {
	//http.HandleFunc("/", sayHello)
	http.HandleFunc("/", tmplDemo)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("http server failed, err", err)
		return
	}
}
