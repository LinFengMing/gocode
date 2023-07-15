package main

import (
	"fmt"
)

func main() {
	var second float64
	fmt.Println("請輸入秒數：")
	fmt.Scanln(&second)
	if second <= 8 {
		var gender string
		fmt.Println("請輸入姓別：")
		fmt.Scanln(&gender)
		if gender == "男" {
			fmt.Println("進入男子組的決賽")
		} else {
			fmt.Println("進入女子組的決賽")
		}
	} else {
		fmt.Println("out")
	}
	var month byte
	var age byte
	var price float64 = 60.0
	fmt.Println("請輸入遊玩的月份：")
	fmt.Scanln(&month)
	fmt.Println("請輸入年齡：")
	fmt.Scanln(&age)
	if month >= 4 && month <= 10 {
		if age > 60 {
			fmt.Printf("%v 月 年齡 %v 票價：%v\n", month, age, price/3)
		} else if age >= 18 {
			fmt.Printf("%v 月 年齡 %v 票價：%v\n", month, age, price)
		} else {
			fmt.Printf("%v 月 年齡 %v 票價：%v\n", month, age, price/2)
		}
	} else {
		if age >= 18 && age < 60 {
			fmt.Println("淡季成人票價：40")
		} else {
			fmt.Println("淡季兒童老人票價：20")
		}
	}
}
