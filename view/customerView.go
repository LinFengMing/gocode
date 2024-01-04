package main

import (
	"fmt"
	"gocode/model"
	"gocode/service"
)

type CustomerView struct {
	key             string
	loop            bool
	customerService *service.CustomerService
}

func (this *CustomerView) list() {
	customers := this.customerService.List()
	fmt.Println("------------客戶列表------------")
	fmt.Println("编號\t姓名\t性别\t年齡\t電話\t\t信箱")
	for i := 0; i < len(customers); i++ {
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Printf("\n----------客戶列表完成----------\n\n")
}
func (this *CustomerView) add() {
	fmt.Println("------------新增客户------------")
	fmt.Println("姓名：")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别：")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年齡：")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("電話：")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("信箱：")
	email := ""
	fmt.Scanln(&email)
	customer := model.NewCustomer2(name, gender, age, phone, email)
	if this.customerService.Add(customer) {
		fmt.Println("------------新增完成------------")
	} else {
		fmt.Println("------------新增失敗------------")
	}
}
func (this *CustomerView) delete() {
	fmt.Println("------------刪除客户------------")
	fmt.Println("請輸入待刪除客户編號(-1退出)：")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return
	}
	fmt.Println("確認是否刪除(Y/N)：")
	choice := ""
	fmt.Scanln(&choice)
	if choice == "y" || choice == "Y" {
		if this.customerService.Delete(id) {
			fmt.Println("------------刪除完成------------")
		} else {
			fmt.Println("------------刪除失敗------------")
		}
	}
}
func (this *CustomerView) exit() {
	fmt.Println("確認是否退出(Y/N)：")
	for {
		fmt.Scanln(&this.key)
		if this.key == "y" || this.key == "Y" || this.key == "n" || this.key == "N" {
			break
		}
		fmt.Println("輸入錯誤，請重新輸入")
	}
	if this.key == "y" || this.key == "Y" {
		this.loop = false
	}
}
func (this *CustomerView) mainMenu() {
	for {
		fmt.Println("----------客戶訊息管理軟體----------")
		fmt.Println("            1 新增客户")
		fmt.Println("            2 修改客户")
		fmt.Println("            3 刪除客户")
		fmt.Println("            4 客户列表")
		fmt.Println("            5 退    出")
		fmt.Print("請選擇(1-5): ")
		fmt.Scanln(&this.key)
		switch this.key {
		case "1":
			this.add()
		case "2":
			fmt.Println("修改客户")
		case "3":
			this.delete()
		case "4":
			this.list()
		case "5":
			this.exit()
		default:
			fmt.Println("輸入錯誤，請重新輸入")
		}
		if !this.loop {
			break
		}
	}
}
func main() {
	customerView := CustomerView{
		key:  "",
		loop: true,
	}
	customerView.customerService = service.NewCustomerService()
	customerView.mainMenu()
	fmt.Println("退出軟體")
}
