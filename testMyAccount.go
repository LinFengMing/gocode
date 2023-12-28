package main

import (
	"fmt"
)

func main() {
	key := ""
	loop := true
	balance := 10000.0
	money := 0.0
	note := ""
	details := "收支\t帳戶金額\t收支金額\t說明"
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
			fmt.Println(details)
		case "2":
			fmt.Println("本次收入金額:")
			fmt.Scanln(&money)
			balance += money
			fmt.Println("本次收入說明:")
			fmt.Scanln(&note)
			details += fmt.Sprintf("\n收入\t%v\t\t%v\t\t%v", balance, money, note)
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
