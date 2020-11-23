# flag

## 简介

`Go`语言内置的`flag`包实现了命令行参数的解析，`flag`包使得开发命令行工具更为简单

## os.Args

使用`os.Args`来获取命令行参数

```go
//os.Args
if len(os.Args) > 0 {
    for index, arg := range os.Args {
        fmt.Printf("arg[%d]=%v\n", index, arg)
    }
}
```

直接执行`go build -o "args_demo"(打包)编译，可能出现报错

```shell script
go: cannot find main module, but found .git/config in /Users/lidean/github_project/fucking-go-library
        to create a module there, run:
        cd ../.. && go mod init
```

根据提示，我们要先创建一个`go.mod`文件

```shell script
go mod init example.com/m
```

然后进入到执行文件的目录下，使用命令`go build -o "args_demo`编译，执行成功会生成一个`args_demo.exe`

接着执行`./args_demo a b c d`，执行结果如下

```shell script
$ ./args_demo a b c d
args[0]=./args_demo
args[1]=a
args[2]=b
args[3]=c
args[4]=d
```

`os.Args`是一个存储命令行参数的字符串切片，它的第一个元素是执行文件的名称

## flag包基本使用

### flag参数类型

|flag参数|有效值|
|----|----|
|字符串flag	|合法字符串|
|整数flag	|1234、0664、0x1234等类型，也可以是负数。|
|浮点数flag	|合法浮点数|
|bool类型flag	|1, 0, t, f, T, F, true, false, TRUE, FALSE, True, False|
|时间段flag	|任何合法的时间段字符串。如”300ms”、”-1.5h”、”2h45m”。 合法的单位有”ns”、”us” /“µs”、”ms”、”s”、”m”、”h”|

### 定义命令行flag参数

#### flag.Type()

基本语法: `flag.Type(flag名, 默认值, 帮助信息)*Type`

举个例子

```go
name := flag.String("name", "张三", "姓名")
age := flag.Int("age", 18, "年龄")
married := flag.Bool("married", false, "婚否")
delay := flag.Duration("d", 0, "时间间隔")
println(name, age, married, delay)
```

输出结果

```text
// 此时name、age、married、delay均为对应类型的指针
0xc000010240 0xc0000140a0 0xc0000140a8 0s
```

#### flag.TypeVar()

基本语法:  `flag.TypeVar(Type指针, flag名, 默认值, 帮助信息)`

举个例子

```go
var name string
var age int
var married bool
var delay time.Duration
flag.StringVar(&name, "name", "张三", "姓名")
flag.IntVar(&age, "age", 18, "年龄")
flag.BoolVar(&married, "married", false, "婚否")
flag.DurationVar(&delay, "d", 0, "时间间隔")
fmt.Println(name, age, married, delay)
```

输出结果

```text
张三 18 false 0s
```

#### flag.Parse()

`flag.Type()`和`flag.TypeVar()`定义命令行`flag参数`，调用flag.Parse()来对命令行参数进行解析

支持的命令行参数格式有以下几种：

- flag xxx （使用空格，一个-符号）
- --flag xxx （使用空格，两个-符号）
- flag=xxx （使用等号，一个-符号）
- --flag=xxx （使用等号，两个-符号）

注意：布尔类型的参数必须使用等号的方式指定

### flag其他函数

- flag.Args() //返回命令行参数后的其他参数，以[]string类型
- flag.NArg() //返回命令行参数后的其他参数个数
- flag.NFlag() //返回使用的命令行参数个数

### 完整例子

下方代码放入`main`函数中，使用`go build -o "flag_demo"`打包好

```go
var name string
var age int
var married bool
var delay time.Duration
flag.StringVar(&name, "name", "张三", "姓名")
flag.IntVar(&age, "age", 18, "年龄")
flag.BoolVar(&married, "married", false, "婚否")
flag.DurationVar(&delay, "d", 0, "时间间隔")

// 解析命令行参数
flag.Parse()
fmt.Println(name, age, married, delay)
//返回命令行参数后的其他参数
fmt.Println(flag.Args())
//返回命令行参数后的其他参数个数
fmt.Println(flag.NArg())
//返回使用的命令行参数个数
fmt.Println(flag.NFlag())
```

### flag使用

命令行参数提示:

```shell script
$ ./flag_demo --help 
Usage of ./flag_demo:
  -age int
        年龄 (default 18)
  -d duration
        时间间隔
  -married
        婚否
  -name string
        姓名 (default "张三")
```

正常使用

```shell script
$ ./flag_demo -name pprof --age 28 -married=false -d=1h30m
pprof 28 false 1h30m0s
[]
0
4
```

使用非flag命令行参数

```shell script
$ ./flag_demo a b c
张三 18 false 0s
[a b c]
3
0
```