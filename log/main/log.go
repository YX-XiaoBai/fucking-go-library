package main

import (
	"log"
	"os"
)

func main() {

	// logger使用
	//log.Println("This is a normal log")
	//v := "normal"
	//log.Printf("this is a %s log\n", v)
	//log.Fatalln("this is a log from fatal")
	//log.Panicln("this is a log from panic")

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

	// flag选项
	//log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	//log.Println("This is a normal log.")

	// 配置log前缀
	//log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	//log.Println("This is a normal log")
	//log.SetPrefix("[pprof]")
	//log.Println("This is a normal log")

	// 配置log输出位置
	//logFile, err := os.OpenFile("./xx.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	//if err != nil {
	//	fmt.Println("open log file failed, err:", err)
	//	return
	//}
	//log.SetOutput(logFile)
	//log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	//log.Println("This is a normal log")
	//log.SetPrefix("[XiaoBai]")
	//log.Println("This is a normal log")

	// 创建logger
	logger := log.New(os.Stdout, "<New>", log.Lshortfile|log.Ldate|log.Ltime)
	logger.Println("this is a custom logger")


}