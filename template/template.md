# template

- [源码地址](./main)

`html/template`包实现了数据驱动的模板，用于生成可对抗代码注入的安全HTML输出。
它提供了和`text/template`包相同的接口，Go语言中输出HTML的场景都应使用`text/template`包

## 模板例子

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Hello</title>
</head>
<body>
    <p>Hello {{.}}</p>
</body>
</html>
```

```go
func sayHello(w http.ResponseWriter, r *http.Request) {
	// 解析指定文件生成模板对象
	tmpl, err := template.ParseFiles("./hello.html")
	if err != nil {
		fmt.Printf("create template failed, err:", err)
		return
	}
	tmpl.Execute(w, "Xiaobai")
}
func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil{
		fmt.Println("http server failed, err", err)
		return
	}
}
```

在任意游览器打开即可看到效果

## 模板语法

`{{.}}`

### 例子

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Hello</title>
</head>
<body>
    <p>Hello {{.Name}}</p>
    <p>Age {{.Age}}</p>
    <p>Gender {{.Gender}}</p>
</body>
</html>
```

```go
type UserInfo struct {
	Name   string
	Gender string
	Age    int
}
func sayHello(w http.ResponseWriter, r *http.Request) {
	// 解析指定文件生成模板对象
	tmpl, err := template.ParseFiles("./hello.html")
	if err != nil {
		fmt.Printf("create template failed, err:", err)
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
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("http server failed, err", err)
		return
	}
}
```

执行成功后，`http://localhost:9090`输出内容如下：

```go
Hello xiaobai
Age 18
Gender boy
```

> 注意：传入的变量是map时，也可以在模板文件中通过.根据key来取值

## pipeline

Go的模板语法中，`pipeline`的概念是传递数据，只要能产生数据的，都是`pipeline`

## 变量

`Action`里可以初始化一个变量来捕获管道的执行结果， 初始化语法如下：

```go
//$variable是变量的名字。声明变量的action不会产生任何输出
$variable := pipeline
```

## 条件判断

Go模板语法中的条件判断有以下几种:

```go
{{if pipeline}} T1 {{end}}
{{if pipeline}} T1 {{else}} T0 {{end}}
{{if pipeline}} T1 {{else if pipeline}} T0 {{end}}
```

### range

Go的模板语法中使用`range`关键字进行遍历，
有两种写法，其中pipeline的值必须是数组、切片、字典或者通道

```go
{{range pipeline}} T1 {{end}}
如果pipeline的值其长度为0，不会有任何输出
{{range pipeline}} T1 {{else}} T0 {{end}}
如果pipeline的值其长度为0，则会执行T0。
```

### with

```go
{{with pipeline}} T1 {{end}}
如果pipeline为empty不产生输出，否则将dot设为pipeline的值并执行T1。不修改外面的dot。
{{with pipeline}} T1 {{else}} T0 {{end}}
如果pipeline为empty，不改变dot并执行T0，否则dot设为pipeline的值并执行T1。
```

## 预定义函数

执行模板时，函数从两个函数字典中查找：首先是模板函数字典，然后是全局函数字典。
一般不在模板内定义函数，而是使用`Funcs`方法添加函数到模板里

预定义的全局函数如下：

```shell script
and
    函数返回它的第一个empty参数或者最后一个参数；
    就是说"and x y"等价于"if x then y else x"；所有参数都会执行；
or
    返回第一个非empty参数或者最后一个参数；
    亦即"or x y"等价于"if x then x else y"；所有参数都会执行；
not
    返回它的单个参数的布尔值的否定
len
    返回它的参数的整数类型长度
index
    执行结果为第一个参数以剩下的参数为索引/键指向的值；
    如"index x 1 2 3"返回x[1][2][3]的值；每个被索引的主体必须是数组、切片或者字典。
print
    即fmt.Sprint
printf
    即fmt.Sprintf
println
    即fmt.Sprintln
html
    返回其参数文本表示的HTML逸码等价表示。
urlquery
    返回其参数文本表示的可嵌入URL查询的逸码等价表示。
js
    返回其参数文本表示的JavaScript逸码等价表示。
call
    执行结果是调用第一个参数的返回值，该参数必须是函数类型，其余参数作为调用该函数的参数；
    如"call .X.Y 1 2"等价于go语言里的dot.X.Y(1, 2)；
    其中Y是函数类型的字段或者字典的值，或者其他类似情况；
    call的第一个参数的执行结果必须是函数类型的值（和预定义函数如print明显不同）；
    该函数类型值必须有1到2个返回值，如果有2个则后一个必须是error接口类型；
    如果有2个返回值的方法返回的error非nil，模板执行会中断并返回给调用模板执行者该错误；
```

## 比较函数

定义：布尔函数会将任何类型的零值视为假，其余视为真。

```go
eq      如果arg1 == arg2则返回真
ne      如果arg1 != arg2则返回真
lt      如果arg1 < arg2则返回真
le      如果arg1 <= arg2则返回真
gt      如果arg1 > arg2则返回真
ge      如果arg1 >= arg2则返回真
```

为简化多参数相等检测，eq（只有eq）可以接受2个或更多个参数，将第一个参数和其余参数依次比较

只适用于基本类型（或重定义的基本类型，如”type Celsius float32”）但是，整数和浮点数不能互相比较

`{{eq arg1 arg2 arg3}}`

## 自定义函数

```go
type UserInfo struct {
	Name   string
	Gender string
	Age    int
}
func sayHello(w http.ResponseWriter, r *http.Request) {
	htmlByte, err := ioutil.ReadFile("./hello.html")
	if err != nil {
		fmt.Println("read html failed, err:", err)
		return
	}
	// 自定义一个kua的模板函数
	kua := func(arg string) (string, error) {
		return arg + "真帅", nil
	}
	// 采用链式操作在Parse之前调用Funcs添加自定义的kua函数
	tmpl, err := template.New("hello").Funcs(template.FuncMap{"kua": kua}).Parse(string(htmlByte))
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
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("http server failed, err", err)
		return
	}
}
```

在html文件中声明`{{kua .Name}}`即可使用

## 嵌套template

这个`template`可以是单独的文件，也可以是通过`define`定义的`template`

### 完整例子

`t.html`文件

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>tmpl test</title>
</head>
<body>
<h1>测试嵌套template语法</h1>
<hr>
{{template "ul.html"}}
<hr>
{{template "ol.html"}}
</body>
</html>
{{ define "ol.html"}}
    <h1>这是ol.html</h1>
    <ol>
        <li>点赞</li>
        <li>关注</li>
        <li>一键三连</li>
    </ol>
{{end}}
```

`ul.html`文件内容如下：

```html
<ul>
    <li>注释</li>
    <li>日志</li>
    <li>测试</li>
</ul>
```

`main.go`文件如下：

```go
// 函数实现
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
    // 注册处理函数
	http.HandleFunc("/", tmplDemo)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("http server failed, err", err)
		return
	}
}
```

执行`main.go`文件，访问`http://localhost`既可以看到页面效果
