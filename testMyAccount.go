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
	flag := false
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
			if flag {
				fmt.Println(details)
			} else {
				fmt.Println("目前沒有收支明細")
			}
		case "2":
			fmt.Println("本次收入金額:")
			fmt.Scanln(&money)
			balance += money
			fmt.Println("本次收入說明:")
			fmt.Scanln(&note)
			details += fmt.Sprintf("\n收入\t%v\t\t%v\t\t%v", balance, money, note)
			flag = true
		case "3":
			fmt.Println("本次支出金額:")
			fmt.Scanln(&money)
			if money > balance {
				fmt.Println("餘額不足")
				break
			}
			balance -= money
			fmt.Println("本次支出說明:")
			fmt.Scanln(&note)
			details += fmt.Sprintf("\n支出\t%v\t\t%v\t\t%v", balance, money, note)
			flag = true
		case "4":
			fmt.Println("你確定要退出嗎? y/n")
			choice := ""
			for {
				fmt.Scanln(&choice)
				if choice == "y" || choice == "n" {
					break
				}
				fmt.Println("請輸入y/n")
			}
			if choice == "y" {
				loop = false
			}
		default:
			fmt.Println("請輸入正確選項")
		}
		if !loop {
			break
		}
	}
	fmt.Println("你退出家庭收支記帳軟體的使用")
}
