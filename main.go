package main

import (
	"fmt"
)

func test(char byte) byte {
	return char + 1
}

func main() {
	var key byte
	fmt.Println("請輸入一個字母 a,b,c,d,e,f,g")
	fmt.Scanf("%c", &key)
	switch test(key) {
	case 'a':
		fmt.Println("星期一")
	case 'b':
		fmt.Println("星期二")
	case 'c':
		fmt.Println("星期三")
	case 'd':
		fmt.Println("星期四")
	case 'e':
		fmt.Println("星期五")
	case 'f':
		fmt.Println("星期六")
	case 'g':
		fmt.Println("星期日")
	default:
		fmt.Println("輸入有誤")
	}
	var n1 int32 = 20
	var n2 int32 = 20
	switch n1 {
	case n2, 10, 5:
		fmt.Println("ok")
	default:
		fmt.Println("沒有一致")
	}
}
