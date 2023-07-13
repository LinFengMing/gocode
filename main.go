package main

import (
	"fmt"
)

func main() {
	var score int
	fmt.Println("請輸入成績：")
	fmt.Scanln(&score)
	if score == 100 {
		fmt.Println("獎勵一輛BMW")
	} else if score <= 99 && score > 80 {
		fmt.Println("獎勵一台 iPhone 14 Pro Max")
	} else if score <= 80 && score >= 60 {
		fmt.Println("獎勵一台 iPad")
	} else {
		fmt.Println("沒有獎勵")
	}
}
