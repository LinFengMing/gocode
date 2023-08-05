package main

import (
	"fmt"
)

func main() {
	var sum int = 0
	for i := 1; i <= 100; i++ {
		sum += i
		if sum > 20 {
			fmt.Println("當 sum > 20 時，當前數是 ", i)
			break
		}
	}
	var name string
	var pwd string
	var loginChance int = 3
	for i := 1; i <= 3; i++ {
		fmt.Println("請輸入用戶名")
		fmt.Scanln(&name)
		fmt.Println("請輸入密碼")
		fmt.Scanln(&pwd)
		if name == "jiro" && pwd == "888" {
			fmt.Println("登陸成功")
			break
		} else {
			loginChance--
			if loginChance > 0 {
				fmt.Printf("還有%v次登陸機會\n", loginChance)
			}
		}
	}
	if loginChance == 0 {
		fmt.Println("沒有登陸成功")
	}
}
