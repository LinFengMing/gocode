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
