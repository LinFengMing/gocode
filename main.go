package main

import (
	"fmt"
)

func main() {
	var i int = 5
	// 二進制輸出
	fmt.Printf("i = %b \n", i)
	// 八進制：0-7，滿8進1，以數字0開頭表示
	var j int = 011 // 8 + 1 = 9
	fmt.Println("j =", j)
	// 十六進制：0-9和A-F，滿16進1，以0x或0X開頭表示
	var k int = 0x11 // 16 + 1 = 17
	fmt.Println("k =", k)
}
