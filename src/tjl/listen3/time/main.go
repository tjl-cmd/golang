package main

import (
	"fmt"
	"time"
)

func testTime() {
	now := time.Now()
	//fmt.Println(now)
	fmt.Printf("current time:%v\n", now)
	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	send := now.Second()
	//打印当前时间
	fmt.Printf("%02d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, send)
	timeStamp := now.Unix()
	fmt.Printf("timeStamp is :%d\n", timeStamp)
}

//时间戳转Time类型
func testTimestamp(timestamp int64) {
	timeObj := time.Unix(timestamp, 0)
	year := timeObj.Year()
	month := timeObj.Month()
	day := timeObj.Day()
	hour := timeObj.Hour()
	minute := timeObj.Minute()
	send := timeObj.Second()
	fmt.Printf("current timestamp:%d\n", timestamp)
	fmt.Printf("%02d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, send)
}
func processTask() {
	fmt.Println("do task")
}

//定时器
func testTicker() {
	ticker := time.Tick(time.Second)
	fmt.Println(time.Duration(1))
	for i := range ticker {
		fmt.Printf("%v\n", i)
		processTask()
	}
}
func testConst() {
	fmt.Printf("nano second :%d\n", time.Nanosecond)  //纳秒
	fmt.Printf("micro second:%d\n", time.Microsecond) //微秒
	fmt.Printf("mili second:%d\n", time.Millisecond)  //毫秒
	fmt.Printf("second:%d\n", time.Second)            //秒
}

//格式化输出时间  2019/12/30 21:53:08
func testFormat() {
	now := time.Now()
	timeStr := now.Format("2006/01/02 15:04:05")
	fmt.Printf("time:%s\n", timeStr)
}
func testFormat2() {
	now := time.Now()
	timeStr := fmt.Sprintf("%02d/%02d/%02d %02d:%02d:%02d\n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Printf("time:%s\n", timeStr)
}
func testCost() {
	start := time.Now().UnixNano()
	for i := 0; i < 10; i++ {
		time.Sleep(time.Microsecond)
	}
	end := time.Now().UnixNano()
	cost := (end - start) / 1000
	fmt.Printf("code cost:%d us\n", cost)
}
func main() {
	//testTime()
	//	timestamp := time.Now().Unix()
	//	testTimestamp(timestamp)
	//	testTicker()
	//testConst()
	//testFormat2()
	testCost()
}
