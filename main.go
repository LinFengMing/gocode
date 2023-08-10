package main

import (
	"fmt"
)

func main() {
here:
	for i := 0; i < 2; i++ {
		for j := 1; j < 4; j++ {
			if j == 2 {
				continue here
			}
			fmt.Println("i =", i, "j =", j)
		}
	}
	for k := 1; k <= 100; k++ {
		if k%2 == 0 {
			continue
		}
		fmt.Println("奇數是", k)
	}
	var positiveCount int
	var negativeCount int
	var num int
	for {
		fmt.Println("請輸入一個整數")
		fmt.Scanln(&num)
		if num == 0 {
			break
		}
		if num > 0 {
			positiveCount++
			continue
		}
		negativeCount++
	}
	fmt.Printf("正數個數是%v, 負數個數是%v", positiveCount, negativeCount)
}
