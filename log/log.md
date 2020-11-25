# log

- [源码地址](./main)

> implements a simple logging package
> 实现一个简单的日志记录包

`Go`语言内置的`log`包实现了简单的日志服务

## 使用Logger

`log`包定义了`Logger`类型，该类型提供了一些格式化输出的方法

直接调用下面三个方法会比自己创建一个`logger`对象更`easy`

- Print系列(Print|Printf|Println）
- Fatal系列（Fatal|Fatalf|Fatalln）
- Panic系列（Panic|Panicf|Panicln）

举个简单的例子

```go
log.Println("This is a normal log。")
v := "normal"
log.Printf("this is a %s log。\n", v)
log.Fatalln("this is a log from fatal")
log.Panicln("this is a log from panic")
```

输出信息如下(日期、时间，默认输出到系统的标准错误)

```shell script
2020/11/23 15:30:11 This is a normal log
2020/11/23 15:30:11 this is a normal log
2020/11/23 15:30:11 this is a log from fatal
exit status 1
```

`Fatal`系列函数会在写入日志信息后调用`os.Exit(1)`。`Panic`系列函数会在写入日志信息后`panic`

## 配置logger

默认`logger`只提供了时间和日志信息，我们往往不够用，所以`log`标准库还提供了设置的方法

- `Flags`函数会`返回`标准`logger`的输出配置
- `SetFlags`函数用来`设置`标准`logger`的输出配置

```go
func Flags() int
func SetFlags(flag int)
```

## flag选项

`log`标准库提供的`flag`选项，是一系列已定义好的常量

```go
const (
    // 控制输出日志信息的细节，不能控制输出的顺序和格式。
    // 输出的日志在每一项后会有一个冒号分隔：例如2009/01/23 01:23:23.123123 /a/b/c/d.go:23: message
    Ldate         = 1 << iota     // 日期：2009/01/23
    Ltime                         // 时间：01:23:23
    Lmicroseconds                 // 微秒级别的时间：01:23:23.123123（用于增强Ltime位）
    Llongfile                     // 文件全路径名+行号： /a/b/c/d.go:23
    Lshortfile                    // 文件名+行号：d.go:23（会覆盖掉Llongfile）
    LUTC                          // 使用UTC时间
    LstdFlags     = Ldate | Ltime // 标准logger的初始值
)
```

直接在log打印前使用`log.SetFlags`设置下输出选项即可

```go
og.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
log.Println("This is a normal log.")
```

输出结果

```shell script
$ go run log.go
2020/11/23 15:43:13.804053 /Users/lidean/github_project/fucking-go-library/log/main/log.go:28: This is a normal log.
```

## 配置log前缀

log库还提供了关于日志信息前缀的两个方法

```go
func Prefix() string
func SetPrefix(prefix string)
```

- `Prefix`函数用来查看标准`logger`的输出前缀
- `SetPrefix`函数用来设置输出前缀

举个例子

```go
log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
log.Println("This is a normal log")
log.SetPrefix("[pprof]")
log.Println("This is a normal log")
```

输出结果

```shell script
2020/11/23 16:00:34.814164 /Users/lidean/github_project/fucking-go-library/log/main/log.go:32: This is a normal log
[pprof]2020/11/23 16:00:34.814294 /Users/lidean/github_project/fucking-go-library/log/main/log.go:34: This is a normal log
```

## 配置log输出位置

语法 

```go
func SetOutput(w io.Writer)
```

`SetOutput`函数用来设置标准`logger`的输出目的地，默认是标准错误输出

举个例子，将日志输出到同目录下的`xx.log`中

```go
logFile, err := os.OpenFile("./xx.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
if err != nil {
    fmt.Println("open log file failed, err:", err)
    return
}
log.SetOutput(logFile)
log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
log.Println("This is a normal log")
log.SetPrefix("[XiaoBai]")
log.Println("This is a normal log")
```

如果你要使用标准的`logger`，我们通常会把上面的配置操作写到`init`函数中，如下

```go
func init() {
    logFile, err := os.OpenFile("./xx.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
    if err != nil {
        fmt.Println("open log file failed, err:", err)
        return
    }
    log.SetOutput(logFile)
    log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
}
```

执行成功后，会生成`xx.log`文件，内容如下

```log
2020/11/23 16:08:44.167391 /Users/lidean/github_project/fucking-go-library/log/main/log.go:48: This is a normal log
[XiaoBai]2020/11/23 16:08:44.167507 /Users/lidean/github_project/fucking-go-library/log/main/log.go:50: This is a normal log
```

## 创建logger

`log`库中还提供了一个创建新`logger`对象的构造函数`New`，支持我们创建自己的`logger`示例。New函数的签名如下

```go
func New(out io.Writer, prefix string, flag int) *Logger
```

创建logger

```go
logger := log.New(os.Stdout, "<New>", log.Lshortfile|log.Ldate|log.Ltime)
logger.Println("this is a custom logger")
```

输出结果

```shell script
<New>2020/11/23 16:21:17 log.go:53: this is a custom logger
```

## 个人总结

总的来说，log 库的功能还是十分有限的，遇到需要记录不用级别日志时就不够用了，第三方的日志库有 logrus / zap 等等。
整个文档来说，首先是讲述了三个格式化输出的方法，然后讲述了 log 配合可以定义常量后自定义 log 输出、配置 log 前缀以及输出在一个单独的log文件中，
最后介绍使用 log.New() 方法去创建一个自定义 log