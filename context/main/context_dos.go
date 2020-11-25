package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg3 sync.WaitGroup

func worker3(ctx context.Context) {
	go worker4(ctx)
LOOP:
	for {
		fmt.Println("worker")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done(): // 等待接收上级通知
			break LOOP
		default:
		}
	}
	wg3.Done()
}

func worker4(ctx context.Context) {
LOOP:
	for {
		fmt.Println("worker2")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done(): // 等待接收上级通知
			break LOOP
		default:
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg3.Add(1)
	go worker3(ctx)
	// 如何优雅的实现结束子goroutine
	time.Sleep(time.Second * 3)
	cancel() // 通知子goroutine结束
	wg3.Wait()
	fmt.Println("over")
}