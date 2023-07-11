package main

import (
	"fmt"
)

func main() {
	var age int
	fmt.Println("請輸入年齡：")
	fmt.Scanln(&age)
	if age > 18 {
		fmt.Println("你的年齡大於18，要對自己的行為負責！")
	}
}
