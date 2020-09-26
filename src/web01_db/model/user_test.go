package model

import (
	"fmt"
	"testing"
)

func TestAddUser(t *testing.T)  {
	fmt.Println("测试添加用户")
	// user := &User{}
	// 调用添加用户的方法
	// user.AddUser()
	t.Run("测试获取用户信息", testGetUserById)
}

// 测试获取一个user
func testGetUserById(t *testing.T)  {
	fmt.Println("测试查询一条信息")
	user := User{
		ID: 1,
	}
	// 调用获取user方法
	u, _ := user.GetUserById()
	fmt.Println("得到user的信息是" ,u)

}
