package main

import (
	"fmt"
)

func sum(n1 int, n2 int) int {
	// 當執行到 defer，暫時不執行，會將 defer 後面的語法放到 defer 的 stack
	// 當函式執行完後，再從 defer 的 stack，按照後進先出的方式出 stack 執行
	defer fmt.Println("ok1 n1 =", n1)
	defer fmt.Println("ok2 n2 =", n2)
	n1++
	n2++
	result := n1 + n2
	fmt.Println("ok3 result =", result)
	return result
}

func main() {
	result := sum(10, 20)
	fmt.Println("result =", result)
}
