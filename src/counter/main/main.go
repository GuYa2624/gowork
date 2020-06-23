package main

import (
	operator2 "counter/operator"
	"fmt"
)

func main() {
	var n1 float64
	var n2 float64
	var operator byte
	for {

		fmt.Println("*************简易计算器*************")
		fmt.Println("1.输入1为加法\n" +
			"2.输入2为减法\n" +
			"3.输入3为乘法\n" +
			"4.输入4为除法\n" +
			"5.输入5退出\n")
		var op int
		fmt.Scan(&op)
		switch op {
		case 1:
			operator = '+'
			fmt.Println("你选择的是加法")
			fmt.Println("请输入数据")
			fmt.Scan(&n1, &n2)
			res := operator2.Operator(n1, n2, operator)
			fmt.Println("计算结果res=", res)
		case 2:
			operator = '-'
			fmt.Println("你选择的是减法")
			fmt.Println("请输入数据")
			fmt.Scan(&n1, &n2)
			res := operator2.Operator(n1, n2, operator)
			fmt.Println("计算结果res=", res)
		case 3:
			operator = '*'
			fmt.Println("你选择的是乘法")
			fmt.Println("请输入数据")
			fmt.Scan(&n1, &n2)
			res := operator2.Operator(n1, n2, operator)
			fmt.Println("计算结果res=", res)
		case 4:
			operator = '/'
			fmt.Println("你选择的是除法")
			fmt.Println("请输入数据")
			fmt.Scan(&n1, &n2)
			res := operator2.Operator(n1, n2, operator)
			fmt.Println("计算结果res=", res)
		default:
			fmt.Println("你选择的是除法")
		}

	}

}
