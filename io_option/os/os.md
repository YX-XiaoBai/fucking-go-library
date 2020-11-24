# io

> provides a platform-independent interface to operating system functionality
> 为操作系统功能提供独立于平台的接口

## 输入输出底层原理

终端其实是一个文件

- `os.Stdin`：标准输入的文件实例，类型为`*File`
- `os.Stdout`：标准输出的文件实例，类型为`*File`
- `os.Stderr`：标准错误输出的文件实例，类型为`*File`

以文件方式操作终端

```go
var buf [16]byte
os.Stdin.Read(buf[:])
os.Stdin.WriteString(string(buf[:]))
```

输出

```shell script
$ go run io.go
# input
XiaoBai is me
# output
XiaoBai is me
```

## 打开&关闭文件

- 只读方式打开一个名称为name的文件

`func Open(name string) (file *File, err Error)`

- 打开名称为name的文件，flag是打开的方式，只读、读写等，perm是权限

`func OpenFile(name string, flag int, perm uint32) (file *File, err Error)`


`os.Open()`函数能够打开一个文件，返回一个`*File`和一个`err`

再通过文件实例调用`close()`方法能够关闭文件

举个例子

```go
// Open()
file, err := os.Open("./os.go")
if err != nil {
    fmt.Println("open file failed! Error:", err)
    return
}
fmt.Println("open file success!")
// 关闭文件
file.Close()

// OpenFile()
file1, err := os.OpenFile("os.log", os.O_RDONLY|os.O_CREATE, 0755)
if err != nil {
    log.Fatal("Error:", err)
}
fmt.Println("open file1 success")
if err := file1.Close(); err != nil {
    log.Fatal("Error:", err)
}
fmt.Println("close file1 success")
```

## 创建文件&写文件

- 根据提供的文件名创建新的文件，返回一个文件对象，默认权限是0666

`func Create(name string) (file *File, err Error)`

- 写入byte类型的信息到文件

`func (file *File) Write(b []byte) (n int, err Error)`

- 写入string信息到文件

`func (file *File) WriteString(s string) (ret int, err Error)`

- 在指定位置开始写入byte类型的信息

`func (file *File) WriteAt(b []byte, off int64) (n int, err Error)`


例子

```go
file, err := os.Create("./os.txt")
if err != nil {
    fmt.Println(err)
    return
}
defer file.Close()
for i := 0; i < 5; i ++ {
    file.WriteString("ab\n")
    file.Write([]byte("cd\n"))
}
file.WriteAt([]byte("fe"), 3)
```

## 读文件

- 读取数据到b中

`func (file *File) Read(b []byte) (n int, err Error)`

- 从off开始读取数据到b中

`func (file *File) ReadAt(b []byte, off int64) (n int, err Error)`

举个例子

```go
file, err := os.Open("./os.txt")
if err != nil {
    fmt.Println("open file err:", err)
    return
}
defer file.Close()
var buf [128]byte
var content []byte
for {
    n, err := file.Read(buf[:])
    //n, err := file.ReadAt(buf[:], int64(0))
    //fmt.Println(string(buf[:n]))
    if err == io.EOF {
        break
    }
    if err != nil{
        fmt.Println("read file err", err)
        return
    }
    content = append(content, buf[:n]...)
}

fmt.Println(string(content))
```

输出结果

```shell script
ab
fe
ab
cd
ab
cd
ab
cd
ab
cd
```

## 拷贝文件

例子

```go
srcFile, err := os.Open("./os.txt")
if err != nil {
    fmt.Println("E1:", err)
    return
}
newFile, err2 := os.Create("./io2.txt")
if err2 != nil {
    fmt.Println("E2", err2)
    return
}
buf := make([]byte, 1024)
for {
    n, err := srcFile.Read(buf)
    if err == io.EOF {
        fmt.Println("Read Done")
        break
    }
    if err != nil {
        fmt.Println("E3", err)
        break
    }
    newFile.Write(buf[:n])
}
srcFile.Close()
newFile.Close()
```

执行成功，执行文件同目录下会生成拷贝文件`io2.txt`



