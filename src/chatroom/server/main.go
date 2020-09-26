package main

import (
	message "chatroom/common"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"net"
)

func readPkg(conn net.Conn) (mes message.Message, err error)  {
	buf := make([]byte, 8096)
	fmt.Println("读取客户端发送的数据")
	// conn.Read之后再conn没有被关闭的情况下，才会阻塞
	// 如果客户端关闭了conn，就不会阻塞
	n, err := conn.Read(buf[:4])
	if n != 4 || err != nil {
		// fmt.Println("conn.Read error=", err)
		return
	}
	// 根据读到的长度转换成一个uint32类型的
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[:4])
	// 根据pkglen读取消息内容
	n,err = conn.Read(buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Read(buf[:pkgLen]) error", err)
		return
	}
	// 将pkglen反序列化成 -> message.Message
	err = json.Unmarshal(buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarshal(buf[:pkgLen], &mes) error" ,err)
		return
	}
	return

}

func WritePkg(conn net.Conn, data []byte) (err error) {
	// 先发送给一个长度给对方
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
	// 发送data数据本身
	n, err = conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Write(buf[0:4]) error", err)
		return
	}
	return
}
// 编写一个函数sverProcessLogin函数，专门处理登录请求
func severProcessLogin(conn net.Conn, mes *message.Message) (err error)  {

	// 1.先从mes中取出mes.Data, 并直接反序列化成LoginMes
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal error", err)
		return
	}
	// 1.先声明一个resMes
	var resMes message.Message
	resMes.Type = message.LoginResMesType

	// 2.再声明一个LoginResMes
	var loginResMes message.LoginResMes

	// 如果用户id=100,密码是123456就正确，否则就是不合法
	if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
		// 合法
		loginResMes.Code = 200 // 200 状态码表示用户存在
	} else {
		// 不合法
		loginResMes.Code = 500 // 500 状态码表示该用户不存在
		loginResMes.Error = "用户不存在，请先注册"
	}
	// 3.将loginResMes序列化
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal fail", err)
		return
	}
	// 4.将data赋值给resMes
	resMes.Data = string(data)

	// 5.对resMes序列化，并发送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal fail", err)
		return
	}
	// 6.发送data 封装到writePkg函数中
	err = WritePkg(conn, data)
	return err
}
// 编写一个ServerProcessMes函数
// 功能：根据客户端发送消息种类不同，决定调用哪一个函数
func serverProcessMes(conn net.Conn, mes *message.Message) (err error)  {
	switch mes.Type {
		case message.LoginMesType:
			// 处理登录逻辑
			err = severProcessLogin(conn, mes)
		case message.RegisterMesType:
			// 处理注册
	default:
		fmt.Println("消息无法处理")
	}
	return

}

// 处理和客户端的通信
func process(conn net.Conn)  {

	// 这里需要延时关闭conn
	defer conn.Close()
	// 读客户端发送的信息
	for {
		// 这里我们将读取数据包直接封装成一个函数readPkg()
		// 循环的读取客户端发的消息
		mes, err := readPkg(conn)
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端关闭了连接")
				return
			}else {
				fmt.Println("readPkg(conn)", err)
				return
			}
		}
		fmt.Println("mes=", mes)
		err = serverProcessMes(conn, &mes)
		if err != nil {
			return
		}
	}
}

func main()  {
	// 提示信息
	fmt.Println("服务器在端口8889端口监听")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	defer listen.Close()
	if err != nil {
		fmt.Println("net.Listen error=", err)
		return
	}
	// 一旦监听成功，就等待客户端来连接服务器
	for {
		fmt.Println("等待客户端谅解服务器。。。。")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept() error=", err)
			return
		}

		// 一旦连接成功，就启动一个协程和客户端保持通讯。。
		go process(conn)
	}
}
