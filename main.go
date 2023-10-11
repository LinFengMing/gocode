package main

import (
	"fmt"
)

func peach(n int) int {
	if n > 10 || n < 1 {
		fmt.Println("輸入天數不對")
		return 0
	}
	if n == 10 {
		return 1
	} else {
		return (peach(n+1) + 1) * 2
	}
}

func main() {
	fmt.Println("第1天桃子的數量是", peach(1))
}
