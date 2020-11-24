# fmt

> implements formatted I/O with functions analogous to C's printf and scanf.
> 使用类似于C的printf和scanf函数实现格式化的I/O。

实现了类似 `C` 的 `prinf` 和 `scanf` 的格式化 `I/O`，主要为 向外输出内容 和 获取输入内容 两部分

## 向外输出

### Print

区别三者
- `Print`函数直接输出内容
- `Printf`函数支持格式化输出字符串
- `Println`函数会输出内容的结尾添加一个换行符(\n)

```go
func Print(a ...interface{})(n int, err error)
func Printf(format string, a ...interface{})(n int, err error)
func Println(a ...interface{})(n int, err error)
```

写个例子来理解

```go
func main(){
    fmt.Print("meet a plmm in the starbucks")
    name := "Nana"
    fmt.Printf("Her name is: %s\n, name")
    fmt.Println("I'm going to have her wechat") 
}
```

运行结果

```text
I meet a plmm in the starbucks.Her name is: Nana
I'm going to have her wechat
```

### Fprint

`Fprint` 系列函数会将内容输出到一个 `io.Writer` 接口类型的变量 `w` 中，通常用这个函数往文件中写入内容

```go
func Fprint(w io.Writer, a ...interface{})(n int, err error)
func Fprintf(w io.Writer, format string, a ...interface{})(n int, err error)
func Fprintln(w io.Writer, a ...interface{})(n int, err error)
```

写个例子来理解

```go
fmt.Fprintln(os.Stdout, "向标准输出写入内容")
fileObj, err := os.OpenFile("./xx.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
if err != nil { 
    fmt.Println("Open File Error, ERROR:", err)
    return
}
name := "YX-XiaoBai"
fmt.Fprintf(fileObj,"put content in the file: %s", name)
```

> 注意：只要满足io.Writer接口的类型都支持写入

### Sprint

`Sprint` 系列函数会把传入的数据生成并返回一个字符串

```go
func Sprint(a ...interface{}) string
func Sprintf(format string, a ...interface{}) string
func Sprintln(a ...interface{}) string
```

写个例子来理解

```go
s1 := fmt.Sprint("YX-XiaoBai")
name := "XiaoBai"
age := 18
s2 := fmt.Sprintf("name:%s,age:%d", name, age)
s3 := fmt.Sprintln("XiaoBai")
fmt.Println(s1, s2, s3)
```

### Errorf

`Errorf` 函数根据 `format` 参数生成格式化字符串并返回一个包含该字符串的错误

```go
func Errorf(format string, a ...interface{}) error
```

举个例子来理解

```go
err := fmt.Errorf("This is a error")
```

## 格式化占位符

`*printf` 系列函数都支持 `format` 格式化参数，在这里我们按照占位符将被替换的变量类型划分，方便查询和记忆

### 通用占位符

| 占位符 | 说明 |
| ---- | ---- |
|%v|值的默认格式表示|
|%+v|类似%v，但输出结构体时会添加字段名|
|%#v|值的Go语法表示|
|%T|打印值的类型|
|%%|百分号|

```go
fmt.Printf("%v\n", 100)
fmt.Printf("%v\n", false)
o := struct {name string}{"XiaoBai"}
fmt.Printf("%v\n", o)
fmt.Printf("%#v\n", o)
fmt.Printf("%+v\n", o)
fmt.Printf("%T\n", o)
fmt.Printf("100%%\n")
```

输出结果

```go
100
false
{XiaoBai}
struct { name string }{name:"XiaoBai"}
{name:XiaoBai}
struct { name string }
100%
```

### 布尔型

|占位符|说明|
|----|----|
|%t|true或false|

### 整型

|占位符|说明|
|----|----|
|%b|表示为二进制|
|%c|该值对应的unicode码值|
|%d|表示为十进制|
|%o|表示为八进制|
|%x|表示为十六进制，使用a-f|
|%X|表示为十六进制，使用A-F|
|%U|表示为Unicode格式：U+1234，等价于”U+%04X”|
|%q|该值对应的单引号括起来的go语法字符字面值，必要时会采用安全的转义表示|

举个例子说明

```go
n := 65
fmt.Printf("%b\n", n)
fmt.Printf("%c\n", n)
fmt.Printf("%d\n", n)
fmt.Printf("%o\n", n)
fmt.Printf("%x\n", n)
fmt.Printf("%X\n", n)
fmt.Printf("%U\n", n)
fmt.Printf("%q\n", n)
```

输出结果

```go
1000001
A
65
101
41
41
U+0041
'A'
```

### 浮点数和复数

|占位符	|说明|
|----|----|
|%b	|无小数部分、二进制指数的科学计数法，如-123456p-78|
|%e	|科学计数法，如-1234.456e+78|
|%E	|科学计数法，如-1234.456E+78|
|%f	|有小数部分但无指数部分，如123.456|
|%F	|等价于%f|
|%g	|根据实际情况采用%e或%f格式（以获得更简洁、准确的输出）|
|%G	|根据实际情况采用%E或%F格式（以获得更简洁、准确的输出）|

举个例子说明

```go
f := 43.21
fmt.Printf("%b\n", f)
fmt.Printf("%e\n", f)
fmt.Printf("%E\n", f)
fmt.Printf("%f\n", f)
fmt.Printf("%g\n", f)
fmt.Printf("%G\n", f)
```

输出结果

```go
6081266871833723p-47
4.321000e+01
4.321000E+01
43.210000
43.210000
43.21
43.21
```

### 字符串和[]byte

|占位符	|说明|
|----|----|
|%s	|直接输出字符串或者[]byte|
|%q	|该值对应的双引号括起来的go语法字符串字面值，必要时会采用安全的转义表示|
|%x	|每个字节用两字符十六进制数表示（使用a-f)|
|%X	|每个字节用两字符十六进制数表示（使用A-F)|

写个例子

```go
s := "XiaoBai"
fmt.Printf("%s\n", s)
fmt.Printf("%q\n", s)
fmt.Printf("%x\n", s)
fmt.Printf("%X\n", s)
```

输出结果

```go
XiaoBai
"XiaoBai"
5869616f426169
5869616F426169
```

### 指针

|占位符	|说明|
|----|----|
|%p	|表示为十六进制，并加上前导的0x|
|%#p|表示为十六进制，不加上前导的0x|

写个例子

```go
a := 18
fmt.Printf("%p\n", &a)
fmt.Printf("%#p\n", &a)
```

输出结果

```go
0xc00001e0a0
c00001e0a0
```

### 宽度标识符

宽度通过一个紧跟在百分号后面的十进制数指定，如果未指定宽度，则表示值时除必需之外不作填充。精度通过（可选的）宽度后跟点号后跟的十进制数指定。如果未指定精度，会使用默认精度；如果点号后没有跟数字，表示精度为0。举例如下

|占位符	|说明|
|----|----|
|%f	|默认宽度，默认精度|
|%9f	|宽度9，默认精度|
|%.2f	|默认宽度，精度2|
|%9.2f	|宽度9，精度2|
|%9.f	|宽度9，精度0|

例子

```go
n := 66.66
fmt.Printf("%f\n", n)
fmt.Printf("%9f\n", n)
fmt.Printf("%.2f\n", n)
fmt.Printf("%9.2f\n", n)
fmt.Printf("%9.f\n", n)
```

输出结果

```go
66.660000
66.660000
66.66
    66.66
       67
```

### 其他flags

|占位符	|说明|
|----|----|
|’+’|	总是输出数值的正负号；对%q（%+q）会生成全部是ASCII字符的输出（通过转义）；|
|’ ‘|   对数值，正数前加空格而负数前加负号；对字符串采用%x或%X时（% x或% X）会给各打印的字节之间加空格|
|’-’|	在输出右边填充空白而不是默认的左边（即从默认的右对齐切换为左对齐）；|
|’#’|	八进制数前加0（%#o），十六进制数前加0x（%#x）或0X（%#X），指针去掉前面的0x（%#p）对%q（%#q），对%U（%#U）会输出空格和单引号括起来的go字面值；|
|‘0’|	使用0而不是空格填充，对于数值类型会把填充的0放在正负号后面；|

例子

```go
fmt.Println("字符串：")
s := "小白"
fmt.Printf("%s\n", s)
fmt.Printf("%5s\n", s)
fmt.Printf("%-5s\n", s)
fmt.Printf("%05s\n", s)
fmt.Println("正数：")
s1 := 10
fmt.Printf("%q\n", s1)
fmt.Printf("%+x\n", s1)
fmt.Printf("% x\n", s1)
fmt.Printf("%#x\n", s1)
fmt.Println("负数：")
s2 := -10
fmt.Printf("%q\n", s2)
fmt.Printf("%+x\n", s2)
fmt.Printf("% X\n", s2)
fmt.Printf("%#x\n", s2)
```

输出结果

```go
字符串：
    小白
       小白
    小白   
    000小白
正数：
    '\n'
    +a
     a
    0xa
负数：
    %!q(int=-10)
    -a
    -A
    -0xa
```

## 获取输入

Go语言fmt包下有`fmt.Scan`、`fmt.Scanf`、`fmt.Scanln`三个函数，可以在程序运行过程中从标准输入获取用户的输入

### fmt.Scan

函数如下

```go
func Scan(a ...interface{})(n int, err error)
```

- `Scan`从标准输入扫描文本，读取由空白符分隔的值保存到传递给本函数的参数中，换行符视为空白符。
- 本函数返回成功扫描的数据个数和遇到的任何错误。如果读取的数据个数比提供的参数少，会返回一个错误报告原因。

例子

```go
func main() {
    var (
        name    string
        age     int
        married bool                    
)        
fmt.Scan(&name, &age, &married)
fmt.Printf("Scan Result is name:%s age:%d married:%t \n", name, age, married)
}
```

输出结果

```go
// 输入以下三个信息
XiaoBai 18 false
// 输出信息
Scan Result is name:XiaoBai age:18 married:false 
```

### fmt.Scanf

函数如下

```go
func Scanf(format string, a ...interface{})(n int, err error)
```

- Scanf从标准输入扫描文本，根据format参数指定的格式去读取由空白符分隔的值保存到传递给本函数的参数中。
- 本函数返回成功扫描的数据个数和遇到的任何错误。

例子

```go
func main() {
    var (
        name    string
        age     int
        married bool
    )
    fmt.Scanf("1:%s 2:%d 3:%t", &name, &age, &married)
    fmt.Printf("Scan Result is name:%s age:%d married:%t \n", name, age, married)
}
```

输出结果

```go
// 必须严格按照顺序和格式写入
1:XiaoBai 2:18 3:false
扫描结果 name:XiaoBai age:18 married:false 
// 错误输入，报错如下
XiaoBai 10 false
扫描结果 name: age:0 married:false 
```

### fmt.Scanln

函数如下

```go
func Scanln(a ...interface{}) (n int, err error)
```

- Scanln类似Scan，它在遇到换行时才停止扫描。最后一个数据后面必须有换行或者到达结束位置。
- 本函数返回成功扫描的数据个数和遇到的任何错误

例子

```go
    func main() {
        var (
            name    string
            age     int
            married bool
        )
        fmt.Scanln(&name, &age, &married)
        fmt.Printf("Scan Result is name:%s age:%d married:%t \n", name, age, married)
    }
```

输出结果

```go
// input
Xiaobai 18 false
// output
Scan Result is name:Xiaobai age:18 married:false 
```

`fmt.Scanln`遇到回车就结束扫描了，这个比较常用

### bufio.NewReader

针对获取输入完整内容，内容中包括空格

例子

```go
func bufioDemo(){
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Please input contents: ")
    text, _ := reader.ReadString('\n')
    text = strings.TrimSpace(text)
    fmt.Printf("%#v\n", text)                            
}
```

输出结果

```go
// input
Please input contents: XiaoBai is a handsome boy, u know?
// output          
"XiaoBai is a handsome boy, u know?"
```

### Fscan

这几个函数功能分别类似于`fmt.Scan`、`fmt.Scanf`、`fmt.Scanln`三个函数，
只不过它们不是从标准输入中读取数据而是从`io.Reader`中读取数据。

函数如下

```go
func Fscan(r io.Reader, a ...interface{}) (n int, err error)
func Fscanln(r io.Reader, a ...interface{}) (n int, err error)
func Fscanf(r io.Reader, format string, a ...interface{}) (n int, err error)
```

### Sscan

这几个函数功能分别类似于fmt.Scan、fmt.Scanf、fmt.Scanln三个函数，
只不过它们不是从标准输入中读取数据而是从`指定字符串`中读取数据。

函数如下

```go
func Sscan(str string, a ...interface{}) (n int, err error)
func Sscanln(str string, a ...interface{}) (n int, err error)
func Sscanf(str string, format string, a ...interface{}) (n int, err error)
```