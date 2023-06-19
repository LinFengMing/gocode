package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var n1 = 100
	// 查看變數的數據類型
	fmt.Printf("n1 的類型 %T \n", n1)
	var n2 int64 = 10
	// 查看變數的占用字節大小和數據類型
	fmt.Printf("n2 的類型 %T n2 占用字節大小是 %d \n", n2, unsafe.Sizeof(n2))
	// 整數變數在使用時，遵守保小不保大的原則，在保證程序正確運行下，盡量使用佔用空間小的數據類型
	var age byte = 90
	fmt.Println("age =", age)
}
