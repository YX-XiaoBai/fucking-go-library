# context

> 在 Go `http包`的`Server`中，每一个请求在都有一个对应的 `goroutine` 去处理。
> 请求处理函数通常会启动额外的 `goroutine` 用来访问后端服务，比如数据库和RPC服务。
> 用来处理一个请求的 `goroutine` 通常需要访问一些与请求特定的数据，
> 比如终端用户的身份认证信息、验证相关的token、请求的截止时间。 当一个请求被取消或超时时，
> 所有用来处理该请求的 `goroutine` 都应该迅速退出，然后系统才能释放这些 `goroutine` 占用的资源

## 为什么需要context

因为在一些场景下需要接受外部命令以实现退出或者结束子`goroutine`

### 全局变量方式

全局变量方式存在的问题：

1. 使用全局变量在跨包调用时不容易统一
2. 如果`worker`中再启动`goroutine`，就不太好控制了。

```go
var wg sync.WaitGroup
// 使用全局变量
var exit bool
func worker() {
	for {
		fmt.Println("worker")
		time.Sleep(time.Second)
		if exit {
			break
		}
	}
	// 如何接收外部命令实现退出
	wg.Done()
}
func main() {
	wg.Add(1)
	go worker()
	// 如何优雅的实现结束子goroutine
	time.Sleep(time.Second * 3)
    // 修改全局变量实现子goroutine的退出
	exit = true
	wg.Wait()
	fmt.Println("over")
}

```

### 通道方式

管道方式存在的问题：

1. 使用全局变量在跨包调用时不容易实现规范和统一，需要维护一个共用的channel

```go
var wg sync.WaitGroup

func worker(exitChan chan struct{}) {
LOOP:
	for {
		fmt.Println("worker")
		time.Sleep(time.Second)
		// 2通道方式
		select {
		case <-exitChan: // 等待接收上级通知
			break LOOP
		default:
		}
	}
	// 如何接收外部命令实现退出
	wg.Done()
}
func main() {
	// 2通道方式
	var exitChan = make(chan struct{})
	wg.Add(1)
	go worker(exitChan)
	// 如何优雅的实现结束子goroutine
	time.Sleep(time.Second * 3)
	exitChan <- struct{}{} // 给子goroutine发送退出信号
	close(exitChan)
	wg.Wait()
	fmt.Println("over")
}
```

### 官方版

```go
package main
import (
    "context"
    "fmt"
    "sync"
    "time"
)
var wg sync.WaitGroup
func worker(ctx context.Context) {
LOOP:
    for {
        fmt.Println("worker")
        time.Sleep(time.Second)
        select {
        case <-ctx.Done(): // 等待上级通知
            break LOOP
        default:
        }
    }
    wg.Done()
}
func main() {
    ctx, cancel := context.WithCancel(context.Background())
    wg.Add(1)
    go worker(ctx)
    time.Sleep(time.Second * 3)
    cancel() // 通知子goroutine结束
    wg.Wait()
    fmt.Println("over")
}
```

当子`goroutine`又开启另外一个`goroutine`时，只需要将`ctx`传入即可：

```go
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
```

### Context

> Go1.7加入了一个新的标准库 `context`，它定义了`Context`类型，
> 专门用来简化 对于处理单个请求的多个 `goroutine` 之间与请求域的数据、取消信号、截止时间等相关操作，
> 这些操作可能涉及多个 `API` 调用。
>
> 对服务器传入的请求应该创建上下文，而对服务器的传出调用应该接受上下文。
> 它们之间的函数调用链必须传递上下文，或者可以使用`WithCancel`、`WithDeadline`、`WithTimeout`或`WithValue`创建的派生上下文。
> 当一个上下文被取消时，它派生的所有上下文也被取消

### Context接口

该接口定义了四个需要实现的方法。具体签名如下

```go
type Context interface {
    Deadline() (deadline time.Time, ok bool)
    Done() <-chan struct{}
    Err() error
    Value(key interface{}) interface{}
}
```

- Deadline方法需要返回当前Context被取消的时间，也就是完成工作的截止时间（deadline）；
- Done方法需要返回一个Channel，这个Channel会在当前工作完成或者上下文被取消之后关闭，多次调用Done方法会返回同一个Channel；
- Err方法会返回当前Context结束的原因，它只会在Done返回的Channel被关闭时才会返回非空的值；
    - 如果当前Context被取消就会返回Canceled错误；
    - 如果当前Context超时就会返回DeadlineExceeded错误；
- Value方法会从Context中返回键对应的值，对于同一个上下文来说，多次调用Value 并传入相同的Key会返回相同的结果，该方法仅用于传递跨API和进程间跟请求域的数据；

### Background()和TODO()

> Go内置两个函数：Background()和TODO()，这两个函数分别返回一个实现了Context接口的background和todo。我们代码中最开始都是以这两个内置的上下文对象作为最顶层的partent context，衍生出更多的子上下文对象。

> Background()主要用于main函数、初始化以及测试代码中，作为Context这个树结构的最顶层的Context，也就是根Context。

> TODO()，它目前还不知道具体的使用场景，如果我们不知道该使用什么Context的时候，可以使用这个。

> background和todo本质上都是emptyCtx结构体类型，是一个不可取消，没有设置截止时间，没有携带任何值的Context。

###  With系列函数

> context包中还定义了四个With系列函数。

### WithCancel

`WithCancel`的函数签名如下

```go
func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
```

> WithCancel返回带有新Done通道的父节点的副本。当调用返回的cancel函数或当关闭父上下文的Done通道时，将关闭返回上下文的Done通道，无论先发生什么情况。

> 取消此上下文将释放与其关联的资源，因此代码应该在此上下文中运行的操作完成后立即调用cancel。

```go
func gen(ctx context.Context) <-chan int {
        dst := make(chan int)
        n := 1
        go func() {
            for {
                select {
                case <-ctx.Done():
                    return // return结束该goroutine，防止泄露
                case dst <- n:
                    n++
                }
            }
        }()
        return dst
    }
func main() {
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel() // 当我们取完需要的整数后调用cancel
    for n := range gen(ctx) {
        fmt.Println(n)
        if n == 5 {
            break
        }
    }
}
```

上面的示例代码中，`gen`函数在单独的`goroutine`中生成整数并将它们发送到返回的通道。 
`gen`的调用者在使用生成的整数之后需要取消上下文，以免`gen`启动的内部`goroutine`发生泄漏。

### WithDeadline

`WithDeadline`的函数签名如下：

```go
func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc)
```

取消此上下文将释放与其关联的资源，因此代码应该在此上下文中运行的操作完成后立即调用cancel

例子

```go

func main() {
	d := time.Now().Add(50 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	// 尽管ctx会过期，但在任何情况下调用它的cancel函数都是很好的实践。
	// 如果不这样做，可能会使上下文及其父类存活的时间超过必要的时间。
	defer cancel()
	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}
```

上面的代码中，定义了一个50毫秒之后过期的deadline，
然后我们调用context.WithDeadline(context.Background(), d)得到一个上下文（ctx）和一个取消函数（cancel），
然后使用一个select让主程序陷入等待：等待1秒后打印overslept退出或者等待ctx过期后退出。 
因为ctx50毫秒后就过期，所以ctx.Done()会先接收到值，上面的代码会打印ctx.Err()取消原因

### WithTimeout

`WithTimeout`的函数签名如下：

`func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)`

取消此上下文将释放与其相关的资源，因此代码应该在此上下文中运行的操作完成后立即调用cancel

通常用于数据库或者网络连接的超时控制。具体示例如下：

```go
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
```

### WithValue

`WithValue` 函数能够将请求作用域的数据与 `Context` 对象建立关系。声明如下：

`func WithValue(parent Context, key, val interface{}) Context`

> 仅对API和进程间传递请求域的数据使用上下文值，而不是使用它来传递可选参数给函数
>
> 所提供的键必须是可比较的，并且不应该是string类型或任何其他内置类型，以避免使用上下文在包之间发生冲突。
WithValue的用户应该为键定义自己的类型。为了避免在分配给interface{}时进行分配，上下文键通常具有具体类型struct{}。
或者，导出的上下文关键变量的静态类型应该是指针或接口

```go
type TraceCode string
var wg01 sync.WaitGroup

func worker01(ctx context.Context) {
	key := TraceCode("TRACE_CODE")
	traceCode, ok := ctx.Value(key).(string) // 在子goroutine中获取trace code
	if !ok {
		fmt.Println("invalid trace code")
	}
LOOP:
	for {
		fmt.Printf("worker01, trace code:%s\n", traceCode)
		time.Sleep(time.Millisecond * 10) // 假设正常连接数据库耗时10毫秒
		select {
		case <-ctx.Done(): // 50毫秒后自动调用
			break LOOP
		default:
		}
	}
	fmt.Println("worker done!")
	wg01.Done()
}

func main() {
	// 设置一个50毫秒的超时
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*50)
	// 在系统的入口中设置trace code传递给后续启动的goroutine实现日志数据聚合
	ctx = context.WithValue(ctx, TraceCode("TRACE_CODE"), "12512312234")
	wg01.Add(1)
	go worker01(ctx)
	time.Sleep(time.Second * 5)
	cancel() // 通知子goroutine结束
	wg01.Wait()
	fmt.Println("over")
}
```

### 使用context注意事项

- 推荐以参数的方式显示传递`Context`
- 以`Context`作为参数的函数方法，应该把`Context`作为第一个参数。
- 给一个函数方法传递`Context`的时候，不要传递`nil`，如果不知道传递什么，就使用`context.TODO()`
- `Context`的`Value`相关方法应该传递请求域的必要数据，不应该用于传递可选参数
- `Context`是线程安全的，可以放心的在多个`goroutine`中传递

### Client超时取消例子

调用服务端API时如何在客户端实现超时控制？

#### server端

- [源码](./server/main.go)

#### client端

- [源码](./client/main.go)