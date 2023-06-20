package main

import (
	"fmt"
)

func main() {
	var c1 byte = 'a'
	var c2 byte = '0'
	// 是輸出對應字符的 UTF-8 碼值
	fmt.Println("c1 =", c1)
	fmt.Println("c2 =", c2)
	// 需要格式化輸出
	fmt.Printf("c1 = %c c2 = %c\n", c1, c2)
	var c3 int = '北'
	fmt.Printf("c3 = %c c3 對應的碼值 = %d", c3, c3)
}
