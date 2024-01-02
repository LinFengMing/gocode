package main

import (
	"fmt"
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
			fmt.Println("新增客户")
		case "2":
			fmt.Println("修改客户")
		case "3":
			fmt.Println("刪除客户")
		case "4":
			this.list()
		case "5":
			this.loop = false
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
