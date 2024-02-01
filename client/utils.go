package client

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"gocode/common/message"
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
