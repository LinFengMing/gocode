package main

import (
	"fmt"
)

func main() {
	key := ""
	loop := true
	// 顯示主選單
	for {
		fmt.Println("--------------------家庭收支記帳軟體--------------------")
		fmt.Println("                     1.收支明細")
		fmt.Println("                     2.登記收入")
		fmt.Println("                     3.登記支出")
		fmt.Println("                     4.退出軟體")
		fmt.Print("請選擇(1-4): ")
		fmt.Scanln(&key)
		switch key {
		case "1":
			fmt.Println("--------------------當前收支明細記錄--------------------")
		case "2":
		case "3":
			fmt.Println("--------------------登記支出--------------------")
		case "4":
			loop = false
		default:
			fmt.Println("請輸入正確選項")
		}
		if !loop {
			break
		}
	}
	fmt.Println("你退出家庭收支記帳軟體的使用")
}
