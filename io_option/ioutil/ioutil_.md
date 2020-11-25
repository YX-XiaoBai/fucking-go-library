# ioutil

- [源码地址](./main)

> implements some I/O utility functions
> 实现一些I/O实用程序函数

## 简介

- 写文件
- 读取文件

# 例子

```go
package main

import (
	"fmt"
	"io/ioutil"
)

func wr() {
	err := ioutil.WriteFile("./ioutil.txt", []byte("www.baidu.com"), 0666)
	if err != nil {
		fmt.Println("E1",err)
		return
	}
}

func re() {
	content, err := ioutil.ReadFile("./ioutil.txt")
	if err != nil {
		fmt.Println("E2",err)
	}
	fmt.Println(string(content))
}

func main() {
	wr()
	re()
}
```

执行文件，便可以实现读写操作