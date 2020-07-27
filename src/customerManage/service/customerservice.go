package service

import "customerManage/model"

//该Customer结构体的增删改查
type CustomerService struct {
	customer []model.Customer
	// 声明一个字段，表示当前切片含有多少个客户
	// 该字段后面，还可以作为新客户的id + 1
	customerNum int
}