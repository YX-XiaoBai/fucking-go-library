# 反射

反射是指在程序运行期对程序本身进行访问和修改的能力

## 变量的内在机制

- 变量包含类型信息和值信息
    - 类型信息：是静态的元信息，是预先定义好的
    - 值信息：是程序运行过程中动态改变的
    
## 反射的作用

- `reflect`包封装了反射相关的方法
- 获取类型信息：`reflect.TypeOf`，是静态的
- 获取值信息：`reflect.ValueOf`，是动态的

## 空接口与反射

- 反射可以在运行时动态获取程序的各种详细信息
- 反射获取`interface`类型信息

### 反射获取类型信息

```go
func reflect_type(a interface{}) {
	t := reflect.TypeOf(a)
	fmt.Println("类型是:", t)
	k := t.Kind()
	fmt.Println(k)
	switch k {
	case reflect.Float64:
		fmt.Println("a is float64\n")
	case reflect.String:
		fmt.Println("string\n")
	}
}
```

### 反射获取值信息

```go
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
```

### 反射修改值信息

```go
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
```


## 结构体与反射

### 查看类型、字段和方法

```go
// 定义结构体
type User struct {
	Id   int
	Name string
	Age  int
}
// 绑方法
func (u User) Hello() {
	fmt.Println("Hello")
}
// 传入interface{}
func Poni(o interface{}) {
	t := reflect.TypeOf(o)
	fmt.Println("Type is：", t)
	fmt.Println("String Type is：", t.Name())
	// 获取值
	v := reflect.ValueOf(o)
	fmt.Println(v)
	// 可以获取所有属性
	// 获取结构体字段个数：t.NumField()
	for i := 0; i < t.NumField(); i++ {
		// 取每个字段
		f := t.Field(i)
		fmt.Printf("%s : %v", f.Name, f.Type)
		// 获取字段的值信息
		// Interface()：获取字段对应的值
		val := v.Field(i).Interface()
		fmt.Println("val :", val)
	}
	fmt.Println("=================Method====================")
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Println(m.Name)
		fmt.Println(m.Type)
	}
}
func main() {
	u := User{1, "Xiaobai", 20}
	Poni(u)
}
```

### 查看匿名字段

```go
type User1 struct {
	Id   int
	Name string
	Age  int
}
type Boy struct {
	User1
	Addr string
}
func main() {
	m := Boy{User1{1, "xiaobai", 20}, "shenzhen"}
	t := reflect.TypeOf(m)
	fmt.Println(t)
	fmt.Printf("%#v\n", t.Field(0))
	fmt.Printf("%#v\n", reflect.ValueOf(m).Field(0))
}
```

### 修改结构体值

```go
type User3 struct {
	Id int
	Name string
	Age int
}
func SetValue(o interface{}) {
	v := reflect.ValueOf(o)
	fmt.Println(v)
	// 获取指针指向的元素
	v = v.Elem()
	fmt.Println(v)
	// 取字段
	f := v.FieldByName("Name")
	fmt.Println(f)
	if f.Kind() == reflect.String {
		f.SetString("xiaobai")
	}
}
func main() {
	u := User3{1, "xiaohei", 20}
	SetValue(&u)
	fmt.Println(u)
}
```

### 调用方法

```go
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
```

### 获取字段tag

```go
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
```

## 反射练习

- 任务：解析如下配置文件

    - 序列化：将结构体序列化为配置文件数据并保存到硬盘
    - 反序列化：将配置文件内容反序列化到程序的结构体
- 配置文件有server和mysql相关配置

```.env
#this is comment
#this a comment
#[]表示一个section
[server]
ip = 10.238.2.2
port = 8080
[mysql]
username = root
passwd = admin
database = test
host = 192.168.10.10
port = 8000
timeout = 1.2
```

- [reflect_practice](../reflect_practice)