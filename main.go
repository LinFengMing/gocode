package main

import (
	"fmt"
)

func main() {
	var i int32 = 100
	// i 轉成 float
	var n1 float32 = float32(i)
	var n2 int8 = int8(n1)
	var n3 int64 = int64(i)
	fmt.Printf("i = %v, n1 = %v, n2 = %v, n3 = %v", i, n1, n2, n3)
}
