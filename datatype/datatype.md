# datatype

## 介绍

是系统中数据交互不可缺少的内容
这里主要介绍`JSON`、`XML`、`MSGPack`

## json


- `json`是完全独立于语言的文本格式，是k-v的形式 `name:zs`
- 应用场景：前后端交互，系统间数据交互
- `json`使用go语言内置的 `encoding/json` 标准库
- 编码json使用`json.Marshal()`函数可以对一组数据进行JSON格式的编码

函数签名如下

```go
func Marshal(v interface{}) ([]byte, error)
```

### 例子

```go
type Person struct {
	Name string
	Hobby string
}

func main() {
	p := Person{"xiaobai", "girl"}
	b, err := json.Marshal(p)
	if err != nil {
		fmt.Println("json err", err)
	}
	fmt.Println(string(b))
	// format output
	b, err =json.MarshalIndent(p, "","	")
	if err != nil {
		fmt.Println("json err", err)
	}
	fmt.Println(string(b))
}
```

### struct tag

```go
type Person struct {
    Name string `json:"-"`
    Hobby string `json:"hobby"`
}
```

### 通过Map生成Json

```go
student := make(map[string]interface{})
student["name"] = "Xiaobai"
student["age"] = 18
student["sex"] = "man"
b, err := json.Marshal(student)
if err != nil {
    fmt.Println(err)
}
fmt.Println(string(b))
```

### 解码json

签名如下

```go
func Unmarshal(data []byte, v interface{}) error
```

#### 解析到结构体

```go
type Person struct {
	Age       int    `json:"age,string"`
	Name      string `json:"name"`
	Niubility bool   `json:"niubility"`
}
b := []byte(`{"age":"18","name":"Xiaobai","niubility":false}`)
var p Person
err := json.Unmarshal(b, &p)
if err != nil {
    fmt.Println(err)
}
fmt.Println(p)
```

#### 解析到interface

```go 
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
```

## xml

- 是可扩展标记语言，包含声明、根标签、子元素和属性
- 应用场景：配置文件以及`webService`

- [源码文件_go](main/xml.go)
- [源码文件_xml](main/test.xml)

## MSGPack

`MSGPack`是二进制的`json`，性能更快，更省空间
需要安装第三方包：`go get -u github.com/vmihailenco/msgpack`

- [源码文件_go](main/msgPack.go)