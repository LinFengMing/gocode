package client

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"gocode/common/message"
	"net"
)

func Login(userId int, userPwd string) (err error) {
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err =", err)
		return
	}
	// 延時關閉
	defer conn.Close()
	var mes message.Message
	mes.Type = message.LoginMesType
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd
	// 將 loginMes 序列化
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Marshal err =", err)
		return
	}
	mes.Data = string(data)
	// 將 mes 序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err =", err)
		return
	}
	// 先傳送 data 的長度給服務器
	// 先取得到 data 的長度 -> 轉成一個表示長度的切片
	var pkgLen uint32 = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[:4], pkgLen)
	// 發送長度
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(buf) err =", err)
		return
	}
	// fmt.Printf("客戶端發送消息的長度 = %d, 內容 = %s", len(data), string(data))
	// 發送消息本身
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write(data) err =", err)
		return
	}
	mes, err = readPkg(conn)
	if err != nil {
		fmt.Println("readPkg err =", err)
		return
	}
	// 將 mes 的 Data 部分反序列化成 LoginResMes
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResMes.Code == 200 {
		fmt.Println("登入成功")
	} else if loginResMes.Code == 500 {
		fmt.Println(loginResMes.Error)
	}
	return
}
