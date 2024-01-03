package service

import "gocode/model"

type CustomerService struct {
	customers   []model.Customer
	customerNum int
}

func NewCustomerService() *CustomerService {
	customerService := &CustomerService{}
	customerService.customerNum = 1
	customer := model.NewCustomer(1, "jiro", "男", 42, "0960535590", "jiro853141@gmail.com")
	customerService.customers = append(customerService.customers, customer)
	return customerService
}
func (this *CustomerService) List() []model.Customer {
	return this.customers
}
func (this *CustomerService) Add(customer model.Customer) bool {
	this.customerNum++
	customer.Id = this.customerNum
	this.customers = append(this.customers, customer)
	return true
}
func (this *CustomerService) FindById(id int) int {
	index := -1
	for i := 0; i < len(this.customers); i++ {
		if this.customers[i].Id == id {
			index = i
		}
	}
	return index
}
func (this *CustomerService) Delete(id int) bool {
	index := this.FindById(id)
	if index == -1 {
		return false
	}
	this.customers = append(this.customers[:index], this.customers[index+1:]...)
	return true
}
