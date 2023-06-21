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
	fmt.Printf("i = %v, n1 = %v, n2 = %v, n3 = %v \n", i, n1, n2, n3)
	// 被轉換的是變數儲存的值，變數本身的數據類型別沒有改變
	fmt.Printf("i type is %T \n", i)
	// 只是轉換的結果是按溢出處理，和我們希望的結果不同
	var num1 int64 = 999999
	var num2 int8 = int8(num1)
	fmt.Println("num2 =", num2)
	var n4 int32 = 12
	var n5 int64
	var n6 int8
	n5 = int64(n4) + 20
	n6 = int8(n4) + 20
	fmt.Println("n5 =", n5, "n6 =", n6)
}
