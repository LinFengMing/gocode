package model

import (
	"fmt"
)

type account struct {
	name     string
	password string
	balance  float64
}

func NewAccount(name string, password string, balance float64) *account {
	if len(name) < 6 || len(name) > 10 {
		fmt.Println("name length should be between 6 and 10")
		return nil
	}
	if len(password) != 6 {
		fmt.Println("password length should be 6")
		return nil
	}
	if balance < 20 {
		fmt.Println("balance should be greater than 20")
		return nil
	}

	return &account{
		name:     name,
		password: password,
		balance:  balance,
	}
}

func SetName(a *account, name string) {
	if len(name) >= 6 && len(name) <= 10 {
		a.name = name
	} else {
		fmt.Println("name length should be between 6 and 10")
	}
}

func GetName(a *account) string {
	return a.name
}

func SetPassword(a *account, password string) {
	if len(password) == 6 {
		a.password = password
	} else {
		fmt.Println("password length should be 6")
	}
}

func GetPassword(a *account) string {
	return a.password
}

func SetBalance(a *account, balance float64) {
	if balance >= 20 {
		a.balance = balance
	} else {
		fmt.Println("balance should be greater than 20")
	}
}

func GetBalance(a *account) float64 {
	return a.balance
}
