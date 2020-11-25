# http

- [源码地址](./main)

`Go`语言内置`net/http包`提供了`HTTP`客户端和服务端的实现

## net/http介绍

提供了`HTTP`客户端和服务端的实现

## http协议

超文本传输协议（HTTP，`HyperText Transfer Protocol`)是互联网上应用最为广泛的一种网络传输协议，
所有的WWW文件都必须遵守这个标准。设计HTTP最初的目的是为了提供一种发布和接收HTML页面的方法

### http客户端

基本的`HTTP/HTTPS`请求 `Get`、`Head`、`Post`和`PostForm`函数发出HTTP/HTTPS请求

例子

```go
resp, err := http.Get("https://www.baidu.com/")
...
resp, err := http.Post("https://www.baidu.com/upload", "image/jpeg", &buf)
...
resp, err := http.PostForm("https://www.baidu.com//form",
    url.Values{"key": {"Value"}, "id": {"123"}})
```

使用完`response`后必须关闭回复的主体

```go
resp, err := http.Get("https://www.baidu.com/")
if err != nil {
    // handle error
}
defer resp.Body.Close()
body, err := ioutil.ReadAll(resp.Body)
// ...
```

### GET示例

写一个简单发送`http`请求的`client`，如下

```go
package main
import (
    "fmt"
    "io/ioutil"
    "net/http"
)
func main() {
    resp, err := http.Get("https://www.5lmh.com/")
    if err != nil {
        fmt.Println("get failed, err:", err)
        return
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("read from resp.Body failed,err:", err)
        return
    }
    fmt.Print(string(body))
}
```

执行之后就能在终端打印`baidu.com`网站首页的内容了

浏览器，其实就是一个发送和接收`HTTP`协议数据的客户端

浏览器访问网页，其实就是从网站的服务器接收`HTTP`数据，然后浏览器会按照`HTML`、`CSS`等规则将网页渲染展示出来

### 带参数的GET请求示例

`GET`请求的参数需要使用`Go`语言内置的`net/url`这个标准库

例子：

```go
apiUrl := "http://127.0.0.1:9090/get"
// URL param
data := url.Values{}
data.Set("name", "XiaoBai")
data.Set("age", "18")
u, err := url.ParseRequestURI(apiUrl)
if err != nil {
    fmt.Printf("parse url requestUrl failed,err:%v\n", err)
}
u.RawQuery = data.Encode() // URL encode
fmt.Println(u.String())
resp, err := http.Get(u.String())
if err != nil {
    fmt.Println("post failed, err:%v\n", err)
    return
}
defer resp.Body.Close()
b, err := ioutil.ReadAll(resp.Body)
if err != nil {
    fmt.Println("get resp failed,err:%v\n", err)
    return
}
fmt.Println(string(b))
// output
//http://127.0.0.1:9090/get?age=18&name=XiaoBai
//404 page not found
```

对应`Server`的`HandlerFunc`如下

```go
func getHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	data := r.URL.Query()
	fmt.Println(data.Get("name"))
	fmt.Println(data.Get("age"))
	answer := `{"status": "ok"}`
	w.Write([]byte(answer))
}
```

### Post请求示例

```go
	url := "http://127.0.0.1:9090/post"
	// 表单数据
	//contentType := "application/x-www-form-urlencoded"
	//data := "name=小白&age=18
	// json
	contentType := "application/json"
	data := `{"name":"小白","age":18}`
	resp, err := http.Post(url, contentType, strings.NewReader(data))
	if err != nil {
		fmt.Println("post failed, err:%v\n", err)
		return
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("get resp failed,err:%v\n", err)
		return
	}
	fmt.Println(string(b))
}
```

对应的`Server`端`HandlerFunc`如下：

```go
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
```

### 自定义Client

要管理HTTP客户端的头域、重定向策略和其他设置，创建一个`Client`

```go
client := &http.Client{
    CheckRedirect: redirectPolicyFunc,
}
resp, err := client.Get("http://baidu.com")
// ...
req, err := http.NewRequest("GET", "http://baidu.com", nil)
// ...
req.Header.Add("If-None-Match", `W/"wyzzy"`)
resp, err := client.Do(req)
// ...
```

### 自定义Transport

要管理代理、TLS配置、keep-alive、压缩和其他设置，创建一个`Transport`

```go
tr := &http.Transport{
    TLSClientConfig:    &tls.Config{RootCAs: pool},
    DisableCompression: true,
}
client := &http.Client{Transport: tr}
resp, err := client.Get("https://baidu.com")
```

`Client` 和 `Transport` 类型都可以安全的被多个go程同时使用。

出于效率考虑，应一次建立、尽量重用

### 服务端

#### 默认Server

`ListenAndServe`使用指定的监听地址和处理器启动一个HTTP服务端。处理器参数通常是`nil`，这表示采用包变量`DefaultServeMux`作为处理器。

`Handle`和`HandleFunc`函数可以向`DefaultServeMux`添加处理器，如

```go
http.Handle("/foo", fooHandler)
http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
})
log.Fatal(http.ListenAndServe(":8080", nil))
```

### 默认Server示例

使用Go语言中的`net/http包`来编写一个简单的接收HTTP请求的Server端示例，`net/http包`是对`net包`的进一步封装，专门用来处理HTTP协议的数据。具体的代码如下

```go
// http server
func sayHello(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello Xiaobai！")
}
func main() {
    http.HandleFunc("/", sayHello)
    err := http.ListenAndServe(":9090", nil)
    if err != nil {
        fmt.Printf("http server failed, err:%v\n", err)
        return
    }
}
```

#### 自定义Server

```go
s := &http.Server{
    Addr:           ":8080",
    Handler:        myHandler,
    ReadTimeout:    10 * time.Second,
    WriteTimeout:   10 * time.Second,
    MaxHeaderBytes: 1 << 20,
}
log.Fatal(s.ListenAndServe())
```

## 个人总结

本章主要讲述`http`包在`server`、`client`以及`http`请求的应用