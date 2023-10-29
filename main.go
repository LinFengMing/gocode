package main

import (
	"fmt"
)

func main() {
	var arr01 [3]int
	arr01[0] = 1
	arr01[1] = 3
	arr01[2] = 5
	fmt.Println(arr01)
	// 預設值，數字為 0，字串為 ""，boolen 為 falst
	var arr02 [3]float32
	var arr03 [3]string
	var arr04 [3]bool
	fmt.Printf("arr02 = %v, arr03 = %v, arr04 = %v\n", arr02, arr03, arr04)
}
