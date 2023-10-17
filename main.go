package main

import (
	"fmt"
)

func printPyramid(totalLevel int) {
	for i := 1; i <= totalLevel; i++ {
		for k := 1; k <= totalLevel-i; k++ {
			fmt.Print(" ")
		}
		for j := 1; j <= 2*i-1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func printMult(num int) {
	for i := 1; i <= num; i++ {
		for j := 1; j <= 1; j++ {
			fmt.Printf("%v * %v = %v", j, i, j*i)
		}
		fmt.Println()
	}
}

func main() {
	var n int
	fmt.Println("請輸入打印金字塔的層數")
	fmt.Scanln(&n)
	printPyramid(n)
	var num int
	fmt.Println("請輸入九九表的對應數")
	fmt.Scanln(&num)
	printMult(num)
}
