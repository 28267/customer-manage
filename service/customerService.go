package service

import (
	"fmt"

	"github.com/28267/customer-manage/model"
)

type CustomerService struct {
	Customers []model.Customer
	//客户id
	CustomerNum int
}

func NewCustomerService() *CustomerService {
	customerService := &CustomerService{}
	customerService.CustomerNum = 1
	Customer := model.NewCustomer(1, "tom", "男", 18, 17536259220, "accardokori@gmail.com", "london")
	customerService.Customers = append(customerService.Customers, Customer)
	return customerService
}

func (customerService CustomerService) List() []model.Customer {
	return customerService.Customers
}

func (customerService *CustomerService) Add(customer model.Customer) bool {
	customerService.CustomerNum++
	customer.ID = customerService.CustomerNum
	customerService.Customers = append(customerService.Customers, customer)
	return true
}

func (customerService *CustomerService) Delete(id int) bool {
	index := customerService.FindByID(id)
	customerService.Customers = append(customerService.Customers[:index],
		customerService.Customers[index+1:]...)
	return true
}

func (customerService *CustomerService) JudgeIdExist(id int) {

}

func (customerService *CustomerService) Update(id int, customer model.Customer) bool {
	index := customerService.FindByID(id)
	if index == -1 {
		fmt.Println("您需要修改的客户ID不存在")
		return false
	}
	customerService.Customers = append(customerService.Customers[:index],
		customerService.Customers[index+1:]...)
	customer.ID = id
	customerService.Customers = append(customerService.Customers, customer)
	return true
}

func (customerService *CustomerService) FindByID(id int) int {
	index := -1
	for i := 0; i < len(customerService.Customers); i++ {
		if customerService.Customers[i].ID == id {
			index = i
		}
	}
	return index
}
