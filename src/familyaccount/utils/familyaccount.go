package utils

import (
	"fmt"
)

type FamilyAccount struct {
	// 声明必要的字段
	// 声明一个变量，保存和接收用户输入
	key string

	// 声明一个变量，控制是否退出for
	loop bool
	// 定义账户余额
	balance float64
	// 每次收支的金额
	money float64
	// 每次收支的说明
	note string
	// 收支的详情，使用字符串来记录
	// 当前收支是，只需要对details进行拼接就好
	details string
	// 显示主菜单
}

// 将各个功能写成一个方法
func (this *FamilyAccount) showDatails()  {
	fmt.Println("-----------当前收支明细记录-----------")
	fmt.Println(this.details)
}

func (this *FamilyAccount) income() {
	fmt.Println("本次收入金额：")
	fmt.Scanln(&this.money)
	this.balance += this.money // 修改账户余额
	fmt.Println("本次收入说明：")
	fmt.Scanln(&this.note)
	// 将收入情况拼接到details变量
	this.details += fmt.Sprintf("\n收入\t%v\t%v\t%v", this.balance, this.money, this.note)
}

func (this *FamilyAccount) pay() {
	fmt.Println("本次支出金额：")
	fmt.Scanln(&this.money)
	// 做一个金额的判断
	if  this.money > this.balance{
		fmt.Println("余额不足")
	} else {
		this.balance -= this.money
	}
	fmt.Println("本次支出说明：")
	fmt.Scanln(&this.note)
	// 将支出情况拼接到details变量
	this.details += fmt.Sprintf("\n支出\t%v\t%v\t%v", this.balance, this.money, this.note)

}

func (this *FamilyAccount) exit()  {
	fmt.Println("确定退出吗？(y/n):")
	for {
		fmt.Scanln(&this.key)
		if this.key == "y" || this.key == "n" {
			break
		}
		fmt.Println("请输入y或n")
	}
	if this.key == "y" {
		this.loop = false
	}

}
// 给该结构体绑定相应的方法
// 显示主菜单
// 编写一个工厂模式的方法
func NewFamilyAccount() *FamilyAccount  {
	return &FamilyAccount{
		key : "",
		loop : true,
		balance: 1000.00,
		money: 0.0,
		note: "",
		details: "收支\t账户金额\t收支金额\t说 明",
	}

}

func (this *FamilyAccount) MainMenu()  {
	for {
		fmt.Println("\n-----------家庭收支记账软件-----------")
		fmt.Println("             1.收入明细")
		fmt.Println("             2.登记收入")
		fmt.Println("             3.登记支出")
		fmt.Println("             4.退出软件")
		fmt.Println("请选择(1-4):")
		fmt.Scanln(&this.key) // 接收用户输入的指令
		// 使用switch来实现用户选的的功能
		switch this.key {
		case "1":
			this.showDatails()
		case "2":
			this.income()
		case "3":
			this.pay()
		case "4":
			this.exit()
		default:
			fmt.Println("请输入正确的选项..")

		}
		if !(this.loop) {
			break
		}

	}
	fmt.Println("您退出了家庭记账软件的使用")



}
