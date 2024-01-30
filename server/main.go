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
		fmt.Println("mes =", mes)
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
