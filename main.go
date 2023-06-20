package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// bool 類型只能是 true or false
	var b = false
	fmt.Println("b =", b)
	// bool 類型占用儲存空間是一個字節
	fmt.Println("b 的佔用空間 =", unsafe.Sizeof(b))
}
