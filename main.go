package main

import (
	"fmt"
)

func main() {
	var i int = 10
	fmt.Println("i 的地址 =", &i)
	// ptr 是一個指針變數
	// ptr 的類型是 *int
	// ptr 的值是 &i
	var ptr *int = &i
	fmt.Printf("ptr = %v \n", ptr)
	fmt.Printf("ptr 的地址 = %v \n", &ptr)
	fmt.Printf("ptr 指向的值 = %v \n", *ptr)
}
