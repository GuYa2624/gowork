package main

import (
	"fmt"
)

func main()  {

	// 声明一个变量，保存和接收用户输入
	var key string = ""

	// 声明一个变量，控制是否退出for
	loop := true
	// 定义账户余额
	balance := 10000.0
	// 每次收支的金额
	money := 0.0
	// 每次收支的说明
	note := ""
	// 收支的详情，使用字符串来记录
	// 当前收支是，只需要对details进行拼接就好
	details := "收支\t账户金额\t收支金额\t说 明"
	// 显示主菜单
	for {
		fmt.Println("\n-----------家庭收支记账软件-----------")
		fmt.Println("             1.收入明细")
		fmt.Println("             2.登记收入")
		fmt.Println("             3.登记支出")
		fmt.Println("             4.退出软件")
		fmt.Println("请选择(1-4):")

		fmt.Scanln(&key) // 接收用户输入的指令

		// 使用switch来实现用户选的的功能
		switch key {
		case "1":
			fmt.Println("-----------当前收支明细记录-----------")
			fmt.Println(details)
		case "2":
			fmt.Println("本次收入金额：")
			fmt.Scanln(&money)
			balance += money // 修改账户余额
			fmt.Println("本次收入说明：")
			fmt.Scanln(&note)
			// 将收入情况拼接到details变量
			details += fmt.Sprintf("\n收入\t%v\t%v\t%v", balance, money, note)
		case "3":
			fmt.Println("本次支出金额：")
			fmt.Scanln(&money)
			// 做一个金额的判断
			if  money > balance{
				fmt.Println("余额不足")
				break
			} else {
				balance -= money
			}
			fmt.Println("本次支出说明：")
			fmt.Scanln(&note)
			// 将支出情况拼接到details变量
			details += fmt.Sprintf("\n支出\t%v\t%v\t%v", balance, money, note)
		case "4":
			fmt.Println("确定退出吗？(y/n):")
			for {
				fmt.Scanln(&key)
				if key == "y" || key == "n" {
					break
				}
				fmt.Println("请输入y或n")
			}
			if key == "y" {
				loop = false
			}
		default:
			fmt.Println("请输入正确的选项..")

		}
		if !loop {
			break
		}

	}
	fmt.Println("您退出了家庭记账软件的使用")

}
