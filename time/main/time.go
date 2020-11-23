package main

import (
	"fmt"
	"time"
)

func main() {

	// 时间类型
	//now := time.Now() //获取当前时间
	//fmt.Printf("current time: %v\n", now)
	//
	//year := now.Year()
	//month := now.Month()   //月
	//day := now.Day()       //日
	//hour := now.Hour()     //小时
	//minute := now.Minute() //分钟
	//second := now.Second() //秒
	//fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)

	// 时间戳
	//now := time.Now()
	//timestamp1 := now.Unix() //时间戳
	//timestamp2 := now.UnixNano() //纳秒时间戳
	//fmt.Printf("timestamp1 is: %v\n", timestamp1)
	//fmt.Printf("timestamp2 is: %v\n", timestamp2)

	// 转换时间戳
	//now := time.Now()
	//timestamp := now.Unix() //时间戳
	//fmt.Println(timestamp)
	//timeObj := time.Unix(timestamp, 0) //将时间戳转为时间格式
	//fmt.Println(timeObj)
	//year := timeObj.Year()     //年
	//month := timeObj.Month()   //月
	//day := timeObj.Day()       //日
	//hour := timeObj.Hour()     //小时
	//minute := timeObj.Minute() //分钟
	//second := timeObj.Second() //秒
	//fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)


	//now := time.Now() // 获取 现在时间
	//now2 := now.Add(time.Hour) //获取 现在时间+1h
	//// Add
	//later := now.Add(time.Hour * 6)
	//fmt.Println(later)
	//// Sub
	//later2 := now2.Sub(now)
	//fmt.Println(later2)
	//// Equal
	//later3 := now.Equal(now2)
	//fmt.Println(later3)
	//// Before
	//later4 := now2.Before(now)
	//fmt.Println(later4)
	//// After
	//later5 := now2.After(now)
	//fmt.Println(later5)

	// 时间格式化
	//now := time.Now()
	//// 24小时制
	//fmt.Println(now.Format("2006-01-02 15:04:05.000 Mon Jan"))
	//// 12小时制
	//fmt.Println(now.Format("2006-01-02 03:04:05.000 PM Mon Jan"))
	//fmt.Println(now.Format("2006/01/02 15:04"))
	//fmt.Println(now.Format("15:04 2006/01/02"))
	//fmt.Println(now.Format("2006/01/02"))

	// 解析字符串格式时间
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
}
