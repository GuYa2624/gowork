package model

import (
	"fmt"
	"web01_db/utils"
)

// User结构体
type User struct {
	ID int
	Username string
	Password string
	Email 	 string
}

// 添加user的方法
func (user *User) AddUser() error {
	// 写sql语句
	sqlStr := "insert into users(username, password, email) values(?,?,?)"
	// 预编译
	inStmt, err := utils.Db.Prepare(sqlStr) // 传入一个查询语句
	if err != nil {
		fmt.Println("预编译失败", err)
		return err
	}
	// 执行
	_, err2 := inStmt.Exec("admin", "123456", "123@root.com")
	if err2 != nil {
		fmt.Println("执行错误", err2)
		return err2
	}
	return nil
}

// GetUserById 根据用户id查询
func (user *User) GetUserById() (*User, error){
	// 写查询的sql语句
	sqlStr := "select id,username,password,email from users where id = ?"
	// 执行
	row := utils.Db.QueryRow(sqlStr, user.ID)
	// 声明
	var id int
	var username string
	var passwod string
	var email string
	err := row.Scan(&id, &username, &passwod, &email)
	if err != nil {
		return nil, err
	}

	u := &User{
		ID:id,
		Username: username,
		Password: passwod,
		Email: email,
	}
	return u, nil

}
