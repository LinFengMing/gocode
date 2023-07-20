package main

import (
	"fmt"
)

func main() {
	var char byte
	fmt.Println("請輸入一個字母 a,b,c,d,e")
	fmt.Scanln(&char)
	switch char {
	case 'a':
		fmt.Println("A")
	case 'b':
		fmt.Println("B")
	case 'c':
		fmt.Println("C")
	case 'd':
		fmt.Println("D")
	case 'e':
		fmt.Println("E")
	default:
		fmt.Println("other")
	}
	var score float64
	fmt.Println("請輸入成績")
	fmt.Scanln(&score)
	switch int(score / 60) {
	case 1:
		fmt.Println("及格")
	case 0:
		fmt.Println("不及格")
	default:
		fmt.Println("輸入錯誤")
	}
	var month byte
	fmt.Println("請輸入月份")
	fmt.Scanln(&month)
	switch month {
	case 3, 4, 5:
		fmt.Println("spring")
	case 6, 7, 8:
		fmt.Println("summer")
	case 9, 10, 11:
		fmt.Println("autumn")
	case 1, 2, 12:
		fmt.Println("winter")
	default:
		fmt.Println("輸入錯誤")
	}
}
