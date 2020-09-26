package main

import (
	message "chatroom/common"
	"encoding/binary"
	"encoding/json"
	"fmt"
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
	pkgLen =  binary.BigEndian.Uint32(buf[:4])
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