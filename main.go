package main

import (
	"fmt"
)

func main() {
	// 嚴格區分大小寫
	// 不能包含空格
	// _ 是空標識符，用於佔用
	// 25個保留關鍵字
	var num int = 10
	var Num int = 20
	fmt.Printf("num = %v, Num = %v \n", num, Num)
}
