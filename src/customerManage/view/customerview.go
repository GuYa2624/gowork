package main

import (
	"customerManage/model"
	"customerManage/service"
	"fmt"
)

type CustomerView struct {
	key string
	loop bool
	// 增加一个字段
	customrtservice *service.CustomerService

}
// 显示所有的客户信息
func (this *CustomerView) list()  {
	// 获取到当前所有的客户信息（在切片中）
	customer := this.customrtservice.List()
	// 显示
	fmt.Println("-----------客户列表-----------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for i := 0; i < len(customer); i++ {
		fmt.Println(customer[i].GetInfo())
	}
	fmt.Println("---------客户列表完成---------\n")
}
// 得到用户的输入，信息构建新的用户，并完成添加
func (this *CustomerView) add() {
	fmt.Println("-----------添加客户-----------")
	fmt.Println("姓名：")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别：")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄：")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话：")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("邮箱：")
	email := ""
	fmt.Scanln(&email)

	// 构建一个新的customer实例
	// 注意id号，没有让用户输入，因为id是唯一的，需要系统分配
	customer := model.NewCustomer1(name, gender, age, phone, email)
	if this.customrtservice.Add(customer) {
		fmt.Println("添加成功")
	} else {
		fmt.Println("添加失败")
	}

}

// 删除id对应的客户信息
func (this *CustomerView) delete() {
	fmt.Println("---------删除客户---------")
	fmt.Println("请输入您要删除的用户id(-1退出)：")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return
	}
	fmt.Println("是否确认删除(y/n):")
	choic := ""
	fmt.Scanln(&choic)
		// 调用servicedelete方法
	if choic == "y" || choic == "Y"{
		 if this.customrtservice.Delete(id) {
		 	fmt.Println("输出成功")
		 } else {
		 	fmt.Println("删除失败")
		 }
	}else if choic == "n" || choic == "N" {
		return
	}
	}



// 显示主菜单
func (this *CustomerView) mainMenu() {
	for {
		fmt.Println("----------客户信息管理软件-----------")
		fmt.Println("----------1.添加客户-----------")
		fmt.Println("----------2.修改用户-----------")
		fmt.Println("----------3.删除用户-----------")
		fmt.Println("----------4.用户列表-----------")
		fmt.Println("----------5.退   出-----------")
		fmt.Print("请输入(1-5):")

		fmt.Scanln(&this.key)
		switch this.key {
		case "1":
			this.add()
		case "2":
			fmt.Println("修改")
		case "3":
			this.delete()
		case "4":
			this.list()
		case "5":
			this.loop = false
		default:
			fmt.Println("输入错误")

		}
		if !this.loop {
			break
		}
	}
	fmt.Println("你退出了客户关系管理软件")

}
func main()  {
	// 在主函数中创建以customerview
	var customer = CustomerView{
		key: "",
		loop: true,
	}
	// 完成初始化
	customer.customrtservice = service.NewCustomerService()
	customer.mainMenu()
}


