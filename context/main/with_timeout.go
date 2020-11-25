package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg0 sync.WaitGroup

func worker0(ctx context.Context) {
LOOP:
	for {
		fmt.Println("db connecting ...")
		time.Sleep(time.Millisecond * 10) // 假设正常连接数据库耗时10毫秒
		select {
		case <-ctx.Done(): // 50毫秒后自动调用
			break LOOP
		default:
		}
	}
	fmt.Println("worker done!")
	wg0.Done()
}
func main() {
	// 设置一个50毫秒的超时
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*50)
	wg0.Add(1)
	go worker0(ctx)
	time.Sleep(time.Second * 5)
	cancel() // 通知子goroutine结束
	wg0.Wait()
	fmt.Println("over")
}
