package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"gocode/common/message"
	"io"
	"net"
)

func readPkg(conn net.Conn) (mes message.Message, err error) {
	buf := make([]byte, 8096)
	fmt.Printf("伺服器在等待客戶端 %s 傳送訊息\n", conn.RemoteAddr().String())
	_, err = conn.Read(buf[:4])
	if err != nil {
		return
	}
	// 根據 buf[:4] 轉成一個 uint32 數據類型
	var pkgLen uint32 = binary.BigEndian.Uint32(buf[:4])
	// 根據 pkgLen 讀取消息內容
	n, err := conn.Read(buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		return
	}
	// 把 pkgLen 反序列化成 -> message.Message
	err = json.Unmarshal(buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarshal err =", err)
		return
	}
	return
}
func writePkg(conn net.Conn, data []byte) (err error) {
	// 先發送一個長度給對方
	var pkgLen uint32 = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[:4], pkgLen)
	// 發送長度
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(buf) err =", err)
		return
	}
	// 發送消息本身
	n, err = conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Write(data) err =", err)
		return
	}
	return
}
func serverProcessLogin(conn net.Conn, mes *message.Message) (err error) {
	// 取出 mes.Data，並直接反序列化成 LoginMes
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal err =", err)
		return
	}
	var resMes message.Message
	resMes.Type = message.LoginResMesType
	var loginResMes message.LoginResMes
	// id = 100, password = 123456
	if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
		loginResMes.Code = 200
	} else {
		loginResMes.Code = 500
		loginResMes.Error = "該用戶不存在，請註冊後再使用"
	}
	// 序列化 loginResMes
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal err =", err)
		return
	}
	resMes.Data = string(data)
	// 將 resMes 序列化
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal err =", err)
		return
	}
	err = writePkg(conn, data)
	return
}
func serverProcessMes(conn net.Conn, mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMesType:
		// 處理登入
		err = serverProcessLogin(conn, mes)
	case message.RegisterMesType:
		// 處理註冊
	default:
		fmt.Println("消息類型不存在，無法處理")
	}
	return
}
func process(conn net.Conn) {
	defer conn.Close()
	for {
		mes, err := readPkg(conn)
		if err != nil {
			if err == io.EOF {
				fmt.Println("客戶端退出了，伺服器也退出")
				return
			} else {
				fmt.Println("readPkg err =", err)
				return
			}
		}
		err = serverProcessMes(conn, &mes)
		if err != nil {
			return
		}
	}
}
func main() {
	fmt.Println("伺服器在8889埠監聽")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	defer listen.Close()
	if err != nil {
		fmt.Println("net.Listen err =", err)
		return
	}
	for {
		fmt.Println("等待客戶端連接服務器")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err =", err)
		}
		fmt.Printf("客戶端 %v 連接服務器成功\n", conn.RemoteAddr().String())
		go process(conn)
	}
}
