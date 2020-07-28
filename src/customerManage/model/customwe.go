package model

import "fmt"

type Customer struct {
	Id int
	Name string
	Gender string
	Age int
	Phone string
	Email string
}

// 使用一个工厂模式，返回一个Customer的实例

func NewCustomer(id int, name string, gender string,
				age int, phone string, email string) Customer  {
	return Customer{
		Id: id,
		Name: name,
		Gender: gender,
		Age: age,
		Phone: phone,
		Email: email,
	}

}
// 第二种不带id的实例
func NewCustomer1(name string, gender string,
	age int, phone string, email string) Customer  {
	return Customer{
		Name: name,
		Gender: gender,
		Age: age,
		Phone: phone,
		Email: email,
	}

}

// 返回格式化后的用户的信息(字符串)
func (this Customer) GetInfo() string {
	info := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\t", this.Id, this.Name, this.Gender, this.Age, this.Phone, this.Email)
	return info
}