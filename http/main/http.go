package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func getHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	data := r.URL.Query()
	fmt.Println(data.Get("name"))
	fmt.Println(data.Get("age"))
	answer := `{"status": "ok"}`
	w.Write([]byte(answer))
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	// 1. 请求类型是application/x-www-form-urlencoded时解析form数据
	r.ParseForm()
	fmt.Println(r.PostForm) // 打印form数据
	fmt.Println(r.PostForm.Get("name"), r.PostForm.Get("age"))
	// 2. 请求类型是application/json时从r.Body读取数据
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("read request.Body failed, err:%v\n", err)
		return
	}
	fmt.Println(string(b))
	answer := `{"status": "ok"}`
	w.Write([]byte(answer))
}

func main() {

	// example for GET
	resp, err := http.Get("https://www.baidu.com/")
	if err != nil {
		fmt.Println("get failed ,err:", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read from resp.Body failed, err:", err)
		return

	}
	fmt.Print(string(body))

	// example for GET with params
	//apiUrl := "http://127.0.0.1:9090/get"
	//// URL param
	//data := url.Values{}
	//data.Set("name", "XiaoBai")
	//data.Set("age", "18")
	//u, err := url.ParseRequestURI(apiUrl)
	//if err != nil {
	//	fmt.Printf("parse url requestUrl failed,err:%v\n", err)
	//}
	//u.RawQuery = data.Encode() // URL encode
	//fmt.Println(u.String())
	//resp, err := http.Get(u.String())
	//if err != nil {
	//	fmt.Println("post failed, err:%v\n", err)
	//	return
	//}
	//defer resp.Body.Close()
	//b, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	fmt.Println("get resp failed,err:%v\n", err)
	//	return
	//}
	//fmt.Println(string(b))

	// example for POST
	//url := "http://127.0.0.1:9090/post"
	// 表单数据
	//contentType := "application/x-www-form-urlencoded"
	//data := "name=小白&age=18
	// json
	//contentType := "application/json"
	//data := `{"name":"小白","age":18}`
	//resp, err := http.Post(url, contentType, strings.NewReader(data))
	//if err != nil {
	//	fmt.Println("post failed, err:%v\n", err)
	//	return
	//}
	//defer resp.Body.Close()
	//b, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	fmt.Println("get resp failed,err:%v\n", err)
	//	return
	//}
	//fmt.Println(string(b))


}
