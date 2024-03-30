package main

import (
	"fmt"
	"os"

	"github.com/28267/customer-manage/model"
	"github.com/28267/customer-manage/service"
)

type CustomerView struct {
	key string
	//是否循环显示菜单
	// loop bool
	customerService *service.CustomerService
}

func (cu *CustomerView) MainMenu() {

	for {
		fmt.Println("------------------客户信息管理系统--------------------")
		fmt.Println()
		fmt.Println("                  1.添加客户          ")
		fmt.Println("                  2.修改客户          ")
		fmt.Println("                  3.删除客户          ")
		fmt.Println("                  4.查询客户          ")
		fmt.Println("                  5.客户列表          ")
		fmt.Println("                  6.退出系统          ")
		fmt.Println("请选择1~5")
		fmt.Scanln(&cu.key)
		switch cu.key {
		case "1":
			cu.viewAdd()
		case "2":
			cu.viewUpdate()
		case "3":
			cu.viewDelete()
		case "4":
			cu.viewInquire()
		case "5":
			cu.viewList()
		case "6":
			fmt.Println("您已成功退出客户信息管理系统..")
			os.Exit(0)
		default:
			fmt.Println("您的输入有误")
		}
	}

}

func (customerView *CustomerView) viewList() {
	fmt.Println("ID\t姓名\t性别\t年龄\t手机号\t\t电子邮件\t\t地址")
	customers := customerView.customerService.List()
	for _, v := range customers {
		fmt.Println(v.GetInfo())
	}
}

func (customerView *CustomerView) viewAdd() {
	fmt.Println("-------------------添加客户---------------------")
	fmt.Println("请输入姓名")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("请输入性别")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("请输入年龄")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("请输入手机号")
	phone := 0
	fmt.Scanln(&phone)
	fmt.Println("请输入电子邮件")
	email := ""
	fmt.Scanln(&email)
	fmt.Println("请输入地址")
	address := ""
	fmt.Scanln(&address)
	customer := model.NewCustomer2(name, gender, age, phone, email, address)
	if customerView.customerService.Add(customer) {
		fmt.Println("----------添加完成----------")
	}
}

func (customerView *CustomerView) viewDelete() {
	fmt.Println("-------------------删除客户---------------------")
	fmt.Println("请输入要删除的客户的ID")
	id := 0
	fmt.Scanln(&id)
	if customerView.customerService.Delete(id) {
		fmt.Printf("ID为%v的客户删除完成\n", id)
	}
}

func (customerView *CustomerView) viewUpdate() {
	fmt.Println("-------------------修改客户---------------------")
	fmt.Println("请输入要修改的客户的ID")
	id := 0
	fmt.Scanln(&id)
	index := customerView.customerService.FindByID(id)
	if index == -1 {
		fmt.Println("客户ID不存在")
		return
	}
	fmt.Println("请输入姓名")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("请输入性别")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("请输入年龄")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("请输入手机号")
	phone := 0
	fmt.Scanln(&phone)
	fmt.Println("请输入电子邮件")
	email := ""
	fmt.Scanln(&email)
	fmt.Println("请输入居住地址")
	address := ""
	fmt.Scanln(&address)
	customer := model.NewCustomer2(name, gender, age, phone, email, address)
	if customerView.customerService.Update(id, customer) {
		fmt.Printf("ID为%v的客户修改完成\n", id)
	}
}

func (customerView *CustomerView) viewInquire() {
	fmt.Println("-------------------查询客户---------------------")
	fmt.Println("请输入要查询的客户ID")
	id := 0
	fmt.Scanln(&id)
	index := customerView.customerService.FindByID(id) + 1
	if index == 0 {
		fmt.Printf("ID为%v的客户不存在\n", id)
		return
	}
	customers := customerView.customerService.List()
	for _, v := range customers {
		if v.ID == index {
			fmt.Printf("ID为%v的客户已找到,信息如下\n", id)
			fmt.Println("ID\t姓名\t性别\t年龄\t手机号\t\t电子邮件\t\t地址")
			fmt.Println(v.GetInfo())
		}
	}
}

func main() {

	customerView := &CustomerView{
		key: "",
	}
	customerView.customerService = service.NewCustomerService()
	customerView.MainMenu()

}
