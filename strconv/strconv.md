# Strconv

- [源码地址](./main)

> implements conversions to and from string representations of basic data types.
> 实现了基本数据类型与其字符串表示的转换

## 常用函数

### Atoi()

`Atoi()`函数用于将`string`类型的整数转换为`int`类型，函数签名如下

```go
func Atoi(s string)(i int, err error)
```

### Itoa()

`Itoa()`函数用于将`int`类型数据转换为对应的`string`，函数签名如下

```go
i2 := 200
s2:= strconv.Itoa(i2)
fmt.Printf("type:%T value:%#v\n", s2, s
```

> 为什么string类型被称作a呢？这是C语言遗留下的典故。
> C语言中没有string类型而是用字符数组(array)表示字符串，所以Itoa对很多C系的程序员很好理解

### Parse系列函数

`Parse类函数` 用于转换 `字符串` 为 `给定类型的值` ：ParseBool()、ParseFloat()、ParseInt()、ParseUint()

#### ParsePool()

```go
func PraseBool(str string) (value bool, err, error)
```

返回字符串表示的`bool`值。它接受`1、0、t、f、T、F、true、false、True、False、TRUE、FALSE`；否则返回错误

#### ParseInt()

```go
func ParseInt(s string, base int, bitSize int) (i int64, err error)
```

返回字符串表示的整数值，接受`正负号`

`base`指定进制（2到36），如果base为0，则会从字符串前置判断，”0x”是16进制，”0”是8进制，否则是10进制；

`bitSize`指定结果必须能无溢出赋值的整数类型，0、8、16、32、64 分别代表 int、int8、int16、int32、int64；

返回的err是*NumErr类型的，如果语法有误，err.Error = ErrSyntax；如果结果超出类型范围err.Error = ErrRange

#### ParseUnit()

```go
func ParseUint(s string, base int, bitSize int) (n uint64, err error)
```

`ParseUint`类似`ParseInt`但不接受正负号，用于无符号整型。

#### ParseFloat()

```go
func ParseFloat(s string, bitSize int) (f float64, err error)
```

解析一个表示浮点数的字符串并返回其值。

如果s合乎语法规则，函数会返回最为接近s表示值的一个浮点数（使用IEEE754规范舍入）。

`bitSize`指定了期望的接收类型，32是float32（返回值可以不改变精确值的赋值给float32），64是float64；

返回值err是*NumErr类型的，语法有误的，err.Error=ErrSyntax；结果超出表示范围的，返回值f为±Inf，err.Error= ErrRange。

### Parse例子

```go
b, err := strconv.ParseBool("true")
if err != nil {
    fmt.Println(err)
}
fmt.Printf("type:%T value:%#v\n", b, b)
f, err1 := strconv.ParseFloat("3.1415", 64)
if err != nil {
    fmt.Println(err1)
}
fmt.Printf("type:%T value:%#v\n", f, f)
i, err2 := strconv.ParseInt("-2", 10, 64)
if err2 != nil {
    fmt.Println(err2)
}
fmt.Printf("type:%T value:%#v\n", i, i)
u, err3 := strconv.ParseUint("2", 10, 64)
if err3 != nil {
    fmt.Println(err3)
}
fmt.Printf("type:%T value:%#v\n", u, u)
```

输出结果

```go
type:bool value:true
type:float64 value:3.1415
type:int64 value:-2
type:uint64 value:0x2
```

### Format系列函数

`Format系列函数` 实现了将 `给定类型数据` 格式化为 `string类型数据` 的功能

#### FormatBool()

```go
func FormatBool(b bool) string
```

根据b的值返回”true”或”false”

#### FormatInt()

```go
func FormatInt(i int64, base int) string
```

返回i的base进制的字符串表示。base 必须在2到36之间，结果中会使用小写字母’a’到’z’表示大于10的数字

#### FormatUint()

```go
func FormatUint(i uint64, base int) string
```

#### FormatFloat()

```go
func FormatFloat(f float64, fmt byte, prec, bitSize int) string
```

函数将浮点数表示为字符串并返回。

`bitSize`表示f的来源类型（32：float32、64：float64），会据此进行舍入。

`fmt`表示格式：’f’（-ddd.dddd）、’b’（-ddddp±ddd，指数为二进制）、’e’（-d.dddde±dd，十进制指数）、’E’（-d.ddddE±dd，十进制指数）、’g’（指数很大时用’e’格式，否则’f’格式）、’G’（指数很大时用’E’格式，否则’f’格式）。

`prec`控制精度（排除指数部分）：对’f’、’e’、’E’，它表示小数点后的数字个数；对’g’、’G’，它控制总的数字个数。如果prec 为-1，则代表使用最少数量的、但又必需的数字来表示f。

### Format例子

```go
s1 := strconv.FormatBool(true)
s2 := strconv.FormatFloat(3.1415, 'E', -1, 64)
s3 := strconv.FormatInt(-2, 16)
s4 := strconv.FormatUint(2, 16)
fmt.Printf("s1:%#v s2:%#v s3:%#v s4:%#v", s1, s2, s3, s4)
```

### 其他函数

#### isPrint()

```go
func IsPrint(r rune) bool
```

例子

```go
ss1 := strconv.IsPrint(123)
fmt.Println(ss1)
//s := []rune("小白")
//ss1 := strconv.IsPrint(s[0])
//fmt.Println(ss1)
```

返回一个字符是否是可打印的，和unicode.IsPrint一样，r必须是：字母（广义）、数字、标点、符号、ASCII空格

#### CanBackquote()

```go
func CanBackquote(s string) bool
```

判断返回字符串s是否可以不被修改的
表示为一个单行的、没有空格和tab之外控制字符的反引号字符串

例子

```go
s2 := "xiao bai"
ss2 := strconv.CanBackquote(s2)
fmt.Println(ss2)
```

## 个人总结

主要讲述了类型之间如何进行转换以及回转换