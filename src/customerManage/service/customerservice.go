package service

import "customerManage/model"

//该Customer结构体的增删改查
type CustomerService struct {
	customer []model.Customer
	// 声明一个字段，表示当前切片含有多少个客户
	// 该字段后面，还可以作为新客户的id + 1
	customerNum int
}

// 编写一个方法，返回*CustomrtService
func NewCustomerService() *CustomerService {
	// 为了能够看到有客户在切片中，我们初始化一个客户
	customerService := &CustomerService{}
	customerService.customerNum = 1
	customer := model.NewCustomer(1, "张三", "男", 20, "12345", "123@gmail.com")
	customerService.customer = append(customerService.customer, customer)
	return customerService
}

// 加一个list方法，返回客户切片
func (this *CustomerService) List() []model.Customer{
	return this.customer

}

// 添加客户到customer切片中
func (this *CustomerService) Add(customer model.Customer) bool {
	// 我们确定一个分配id的规则
	this.customerNum++
	customer.Id = this.customerNum
	this.customer = append(this.customer, customer)
	return true
}

// 根据id从切片中删除
func (this *CustomerService) Delete (id int) bool {
	index := this.FindById(id)
	if index == -1 {
		println("无此用户")
		return false
	} else {
		// 如何从切片中删除元数
		this.customer = append(this.customer[:index], this.customer[index+1:]...)
		return true
	}
}

// 根据id查找客户在切片中对应的下标，如果没有该客户，返回-1
func (this *CustomerService) FindById(id int) int {
	index := -1
	// 遍历切片
	for i := 0; i < len(this.customer); i++ {
		if this.customer[i].Id == id {
			index = i
		}
	}
	return index
}