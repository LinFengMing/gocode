package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	for {
		buf := make([]byte, 8096)
		fmt.Printf("伺服器在等待客戶端 %s 傳送訊息\n", conn.RemoteAddr().String())
		n, err := conn.Read(buf[:4])
		if n != 4 || err != nil {
			fmt.Println("conn.Read err =", err)
			return
		}
		fmt.Println("讀到的 buf =", buf[:4])
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
