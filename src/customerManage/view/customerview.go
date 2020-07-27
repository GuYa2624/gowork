package main

import "fmt"

type CustomerView struct {
	key string
	loop bool

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
			fmt.Println("添加")
		case "2":
			fmt.Println("修改")
		case "3":
			fmt.Println("删除")
		case "4":
			fmt.Println("列表")
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
	customer.mainMenu()
}
