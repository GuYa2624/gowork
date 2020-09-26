package main

import (
	"fmt"
	"os"
)

// 定义两个变量，一个是id，一个是密码
var userId int
var userPwd string

func main() {
	// 登陆界面
	// 接收用户选择
	var key int
	// 判断是否还继续显示菜单
	// var loop = true

	for {
		fmt.Println("----------欢迎登录多人聊太系统-----------")
		fmt.Println("\t\t\t 1 登录聊天室")
		fmt.Println("\t\t\t 2 注册用户")
		fmt.Println("\t\t\t 3 退出系统")
		fmt.Println("\t\t\t 请选择(1-3):")

		fmt.Scanf("%d\n", &key)

		switch key {
		case 1:
			fmt.Println("登录聊天室")
			goto login
		case 2:
			fmt.Println("注册用户")
			break
		case 3:
			fmt.Println("退出系统")
			os.Exit(0)
		default:
			fmt.Println("输入有误")
		}

	}
login :
	// 根据用户的输入显示新的提示信息
	if key == 1 {
		// 说明用户要登录了
		fmt.Println("请输入用户的id号:")
		fmt.Scanf("%d\n", &userId)
		fmt.Println("请输入用户的密码:")
		fmt.Scanf("%s\n", &userPwd)

		// 先把登录函数写到另外一个文件，先写到login.go
		login(userId, userPwd)
		//if err != nil {
		//	fmt.Println("登录失败")
		//} else {
		//	fmt.Println("登录成功")
		//}

	} else if key == 2 {
		fmt.Println("进行用户注册的界面")

	}
}
