package model

import "fmt"

type Customer struct {
	ID      int
	Name    string
	Gender  string
	Age     int
	Phone   int
	Email   string
	Address string
}

func NewCustomer(id int, name string, gender string, age int, phone int,
	email string, address string) Customer {

	return Customer{
		ID:      id,
		Name:    name,
		Gender:  gender,
		Age:     age,
		Phone:   phone,
		Email:   email,
		Address: address,
	}

}

func NewCustomer2(name string, gender string, age int, phone int,
	email string, address string) Customer {
	return Customer{
		Name:    name,
		Gender:  gender,
		Age:     age,
		Phone:   phone,
		Email:   email,
		Address: address,
	}
}

func (customer Customer) GetInfo() string {
	info := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\t%v\t", customer.ID, customer.Name, customer.Gender,
		customer.Age, customer.Phone, customer.Email, customer.Address)
	return info
}
