package main

import (
	"fmt"
	"gocode/model"
)

func main() {
	a := model.NewAccount("abc123", "123456", 100)
	if a == nil {
		fmt.Println("create account failed")
	} else {
		fmt.Println("create account success")
	}
	model.SetName(a, "abc123456")
	model.SetPassword(a, "123456")
	model.SetBalance(a, 50)
	fmt.Printf("name = %v password = %v balance = %v", model.GetName(a), model.GetPassword(a), model.GetBalance(a))
}
