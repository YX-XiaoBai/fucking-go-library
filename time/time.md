# time

主要介绍`Go`内置的`time`包基本用法

## time包

`time`包提供了`时间的显示`和`测量用的函数`。日历的计算采用的是公历

## 时间类型

`time.Time`类型表示时间。我们可以通过`time.Now()`函数获取当前的时间对象，
然后获取时间对象的年月日时分秒等信息

例子

```go
func timeDemo(){
    now := time.Now() //获取当前时间
    fmt.Printf("current time: %v\n", now)
    
    year := now.Year()
    month := now.Month()   //月
    day := now.Day()       //日
    hour := now.Hour()     //小时
    minute := now.Minute() //分钟
    second := now.Second() //秒
    fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}
```

输出结果

```go
current time: 2020-11-20 17:25:43.410152 +0800 CST m=+0.000119957
2020-11-20 17:25:43

```

## 时间戳

时间戳是自1970年1月1日（08:00:00GMT）至当前时间的总毫秒数
它也被称为`Unix`时间戳（`UnixTimestamp`）

获取时间戳例子

```go
now := time.Now()
timestamp1 := now.Unix() //时间戳
timestamp2 := now.UnixNano() //纳秒时间戳
fmt.Printf("timestamp1 is: %v\n", timestamp1)
fmt.Printf("timestamp2 is: %v\n", timestamp2)
```

输出结果

```go
timestamp1 is: 1605864737
timestamp2 is: 1605864737255257000
```

`time.Unix()`函数将时间戳转为时间格式

例子

```go
now := time.Now()
timestamp := now.Unix() //时间戳
fmt.Println(timestamp)
timeObj := time.Unix(timestamp, 0) //将时间戳转为时间格式
fmt.Println(timeObj)
year := timeObj.Year()     //年
month := timeObj.Month()   //月
day := timeObj.Day()       //日
hour := timeObj.Hour()     //小时
minute := timeObj.Minute() //分钟
second := timeObj.Second() //秒
fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
```

输出结果

```go
1605864990
2020-11-20 17:36:30 +0800 CST
2020-11-20 17:36:30
```

## 时间间隔

`time.Duration`是`time`包定义的一个类型，它代表两个时间点之间经过的时间，以纳秒为单位。

`time.Duration`表示一段时间间隔，可表示的最长时间段大约290年

`time`包中定义的时间间隔类型的常量如下

```go
const (
    Nanosecond  Duration = 1
    Microsecond          = 1000 * Nanosecond
    Millisecond          = 1000 * Microsecond
    Second               = 1000 * Millisecond
    Minute               = 60 * Second
    Hour                 = 60 * Minute
)
```

## 时间操作

### Add

要求时间+时间间隔，time对象提供`Add`方法

```go
func (t Time) Add(d Duration) Time
```

### Sub

求两个时间之间的差值

```go
func (t Time) Sub(u Time) Duration
```

### Equal

判断两个时间是否相同，会考虑时区的影响，因此不同时区标准的时间也可正确比较
本方法和用t==u不同，还会比较地点和时区信息

```go
func (t Time) Equal(u Time) bool
```

### Before

```go
func (t Time) Before(u Time) bool
```

### After

```go
func (t Time) After(u Time) bool
```

举个例子说明上面方法

```go 
// 定义两个时间
now := time.Now() // 获取 现在时间
now2 := now.Add(time.Hour) //获取 现在时间+1h
// Add
later := now.Add(time.Hour * 6)
fmt.Println(later)
// Sub
later2 := now2.Sub(now)
fmt.Println(later2)
// Equal
later3 := now.Equal(now2)
fmt.Println(later3)
// Before
later4 := now2.Before(now)
fmt.Println(later4)
// After
later5 := now2.After(now)
fmt.Println(later5)
```

## 时间格式化

`Format`对时间进行格式化

注意：时间模板不是常见的Y-m-d H:M:S

而是：Go的诞生时间2006年1月2号15点04分（记忆口诀为2006 1 2 3 4）

举个例子

```go
now := time.Now()
// 24小时制
fmt.Println(now.Format("2006-01-02 15:04:05.000 Mon Jan"))
// 12小时制
fmt.Println(now.Format("2006-01-02 03:04:05.000 PM Mon Jan"))
fmt.Println(now.Format("2006/01/02 15:04"))
fmt.Println(now.Format("15:04 2006/01/02"))
fmt.Println(now.Format("2006/01/02"))
```

输出结果

```go
2020-11-20 18:03:22.838 Fri Nov
2020-11-20 06:03:22.838 PM Fri Nov
2020/11/20 18:03
18:03 2020/11/20
2020/11/20
```

## 解析字符串格式时间

不多说，直接上例子

```go
now := time.Now()
fmt.Println(now)
// 加载时区
loc, err := time.LoadLocation("Asia/Shanghai")
if err != nil {
    fmt.Println(err)
    return
}
// 按照指定时区和指定格式解析字符串时间
timeObj, err := time.ParseInLocation("2006/01/02 15:04:05", "2020/11/21 14:15:20", loc)
if err != nil {
    fmt.Println(err)
    return
}
fmt.Println(timeObj)
fmt.Println(timeObj.Sub(now))
```