package main

import (
	"fmt"
)

type Point struct {
	x int
	y int
}

func main() {
	var a interface{}
	var point Point = Point{1, 2}
	a = point
	var b Point
	// b = a // error
	b = a.(Point) // 型別斷言
	fmt.Println(b)
	var x interface{}
	var b2 float32 = 1.1
	x = b2
	if y, flag := x.(float32); flag {
		fmt.Println("轉換成功")
		fmt.Printf("y 的型別是 %T 值是 = %v\n", y, y)
	} else {
		fmt.Println("轉換失敗")
	}
	fmt.Println("繼續執行")
}
