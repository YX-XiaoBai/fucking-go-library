# bufio

> implements buffered I/O. It wraps an io.Reader or io.Writer object, creating another object (Reader or Writer) that also implements the interface
> 实现缓冲I/O。它包装io.Reader或者io.Writer对象，创建另一个同样实现接口的对象（读取器或编写器）

## 简介

- 实现了带缓冲区的读写，是对文件读写的封装
- `bufio`缓冲写数据

|模式|含义|
|os.O_WRONLY|只写|
|os.O_CREATE|创建文件|
|os.O_RDONLY|只读|
|os.O_RDWR|读写|
|os.O_TRUNC|清空|
|os.O_APPEND|追加|

## 例子

```go
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func wr() {
	file, err := os.OpenFile("./xx.txt", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for i := 0; i < 1; i++ {
		writer.WriteString("hello\n")
	}
	writer.Flush()
}

func re() {
	file, err := os.Open("./xx.txt")
	if err != nil {
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			return
		}
		fmt.Println(string(line))
	}
}

func main(){
	wr()
	re()
}
```