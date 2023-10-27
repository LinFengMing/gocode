package main

import (
	"fmt"
)

func main() {
	// var score [5]float64
	// for i := 0; i < len(score); i++ {
	// 	fmt.Printf("請輸入第 %d 個元素的值", i+1)
	// 	fmt.Scanln(&score[i])
	// }
	// for i := 0; i < len(score); i++ {
	// 	fmt.Printf("score[%d] = %v\n", i, score[i])
	// }
	// 初始化陣列的方式
	var numArr01 [3]int = [3]int{1, 2, 3}
	fmt.Println("numArr01 =", numArr01)
	var numArr02 = [3]int{4, 5, 6}
	fmt.Println("numArr02 =", numArr02)
	var numArr03 = [...]int{7, 8, 9}
	fmt.Println("numArr03 =", numArr03)
	var numArr04 = [...]int{1: 800, 0: 900, 2: 1000}
	fmt.Println("numArr04 =", numArr04)
	numArr05 := [...]int{1: 1100, 0: 1300, 2: 1200}
	fmt.Println("numArr05 =", numArr05)
}
