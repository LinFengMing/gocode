package main

import (
	"fmt"
)

func main() {
	var num5 = 1.1
	fmt.Printf("num5 的數據類型是 %T \n", num5)
	// 十進制形式
	num6 := 5.12
	num7 := .123 // 0.123
	fmt.Println("num6 =", num6, "num7 =", num7)
	// 科學技數法形式
	num8 := 5.1234e2   // 5.1234 * 10 的 2 次方
	num9 := 5.1234e2   // 5.1234 * 10 的 2 次方
	num10 := 5.1234e-2 // 5.1234 / 10 的 2 次方
	fmt.Println("num8 =", num8, "num9 =", num9, "num10 =", num10)
}
