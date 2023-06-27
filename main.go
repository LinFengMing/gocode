package main

import (
	"fmt"
)

func test() int {
	return 90
}

func main() {
	var i int = 10 // 基本賦值
	fmt.Println("i =", i)
	// 有兩個變數，a 和 b，要求將其交換，最終打印結果
	a := 9
	b := 2
	fmt.Printf("交換前的情況是 a = %v , b = %v \n", a, b)
	t := a
	a = b
	b = t
	fmt.Printf("交換後的情況是 a = %v , b = %v \n", a, b)
	a += 17 // a = a + 17
	fmt.Println("a =", a)
	// 賦值運算的執行順序是從右向左
	var c int = a + 3
	fmt.Println("c =", c)
	// 賦值運算符的左邊只能是變數，右邊可以是變數、表達式、常數
	// 表達式：任何有值都可以看做是表達式
	var d int
	d = a
	d = 8 + 2*8
	d = test() + 90
	fmt.Println("d =", d)
	d = 890
	fmt.Println("d =", d)
}
