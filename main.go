package main

import (
	"fmt"
)

func main() {
	// golang 可以在 if 中，直接定義一個變數
	if age := 20; age > 18 {
		fmt.Println("你的年齡大於18，要對自己的行為負責！")
	}
}
