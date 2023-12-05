package main

import (
	"fmt"
)

type Account struct {
	AccountNo string
	Pwd       string
	Balance   float64
}

func (account *Account) Deposite(money float64, pwd string) {
	if pwd != account.Pwd {
		fmt.Println("密碼錯誤")
		return
	}
	if money <= 0 {
		fmt.Println("金額錯誤")
		return
	}
	account.Balance += money
	fmt.Println("存款成功")
}

func (account *Account) WithDraw(money float64, pwd string) {
	if pwd != account.Pwd {
		fmt.Println("密碼錯誤")
		return
	}
	if money <= 0 || money > account.Balance {
		fmt.Println("金額錯誤")
		return
	}
	account.Balance -= money
	fmt.Println("提款成功")
}

func (account *Account) Query(pwd string) {
	if pwd != account.Pwd {
		fmt.Println("密碼錯誤")
		return
	}
	fmt.Printf("帳號:%v 餘額:%v\n", account.AccountNo, account.Balance)
}

func main() {
	account := Account{
		AccountNo: "gs123456",
		Pwd:       "666666",
		Balance:   100.0,
	}
	account.Query("666666")
	account.Deposite(200.0, "666666")
	account.Query("666666")
	account.WithDraw(150.0, "666666")
	account.Query("666666")
}
