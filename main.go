package main

import (
	"fmt"
	"time"
)

func main() {
	// 取得當前時間
	now := time.Now()
	fmt.Printf("now = %v, type = %T\n", now, now)
	fmt.Printf("年 = %v, \n", now.Year())
	fmt.Printf("月 = %v, \n", int(now.Month()))
	fmt.Printf("日 = %v, \n", now.Day())
	fmt.Printf("時 = %v, \n", now.Hour())
	fmt.Printf("分 = %v, \n", now.Minute())
	fmt.Printf("秒 = %v, \n", now.Second())
	// 格式化日期時間
	dateStr := fmt.Sprintf("%02d-%02d-%02d %02d:%02d:%02d\n", now.Year(),
		int(now.Month()), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Printf("dateStr = %v\n", dateStr)
	fmt.Println(now.Format("2006-01-02 15:04:05"))
}
