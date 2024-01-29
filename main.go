package main

import (
	"fmt"
	"gocode/client"
)

var userId int
var userPwd string

func main() {
	var key int
	var loop bool = true
	for loop {
		fmt.Println("----------歡迎登入多人聊天系統----------")
		fmt.Println("\t\t\t 1 登入聊天室")
		fmt.Println("\t\t\t 2 註冊用戶")
		fmt.Println("\t\t\t 3 退出系統")
		fmt.Println("\t\t\t 請選擇(1-3):")
		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("登入聊天室")
			loop = false
		case 2:
			fmt.Println("註冊用戶")
			loop = false
		case 3:
			fmt.Println("退出系統")
			loop = false
		default:
			fmt.Println("輸入錯誤，請重新輸入")
		}
	}
	if key == 1 {
		// 進行用戶登入
		fmt.Println("請輸入用戶的ID:")
		fmt.Scanf("%d\n", &userId)
		fmt.Println("請輸入用戶的密碼:")
		fmt.Scanf("%s\n", &userPwd)
		err := client.Login(userId, userPwd)
		if err != nil {
			fmt.Println("登入失敗")
		} else {
			fmt.Println("登入成功")
		}
	} else if key == 2 {
		fmt.Println("註冊用戶的流程")
	} else {
		fmt.Println("退出系統")
	}
}
