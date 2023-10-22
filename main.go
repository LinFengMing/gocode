package main

import (
	"fmt"
	"strings"
)

func main() {
	// 尋找字串最後一次出現的位置，沒有回傳 -1
	index := strings.LastIndex("go golang", "go")
	fmt.Printf("index = %v\n", index)
	// 字串取代
	str := strings.Replace("go go hello", "go", "hi", -1)
	fmt.Printf("str = %v\n", str)
	// 字串分割
	strArr := strings.Split("hello,world,ok", ",")
	fmt.Printf("strArr = %v\n", strArr)
	// 字串轉大寫
	str2 := "goLang Hello"
	str2 = strings.ToUpper(str2)
	fmt.Printf("str2 = %v\n", str2)
	// 字串轉小寫
	str3 := "goLang Hello"
	str3 = strings.ToLower(str2)
	fmt.Printf("str3 = %v\n", str3)
}
