package main

import (
	message "chatroom/common"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

// 先写一个函数， 完成登录
func login(userId int, userPwd string) (err error) {
	// 下一步就要开始定协议了
	//fmt.Printf("userId=%d, userPwd=%s\n", userId, userPwd)
	//
	//return nil

	// 1.连接打服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial error=", err)
	}
	// 延时关闭
	defer conn.Close()

	// 2.准备通过conn发送消息给服务器
	var mes message.Message
	mes.Type = message.LoginMesType

	// 3.创建一个loginMes结构体
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd

	// 4.将loginMes序列化
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Marsha error=", err)
		return
	}
	// 5.把data赋给了mes.Data字段
	mes.Data = string(data)

	// 6.将mes进行序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marsha error=", err)
		return
	}
	// 7.data就是我们要发送的数据了
	// 7.1先把data的长度发送给服务器
	// 先获取到data的长度，然后再转成表示长度的byte切片
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	// 8.发送长度
	n, err := conn.Write(buf[0:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(buf[0:4]) error", err)
		return
	}
	// fmt.Printf("客户端发送消息的长度成功长度是%v,内容是%s\n", len(data), data)
	// 发送消息本身
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write(data) error", err)
		return
	}

	// 这里还需要处理服务器端返回的消息.
	mes, err = readPkg(conn)
	if err != nil {
		fmt.Println("readPkg()")
		return
	}

	// 将mes的data部分反序列化成LoginResMes
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResMes.Code == 200 {
		fmt.Println("登录成功")
	} else if loginResMes.Code == 50 {
		fmt.Println(loginResMes.Error)
	}
	return
}
