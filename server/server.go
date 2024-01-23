package main

import (
	"fmt"
	"net"
)

func proccess(conn net.Conn) {
	defer conn.Close()
	for {
		buf := make([]byte, 1024)
		fmt.Printf("伺服器在等待用戶端%s發送訊息\n", conn.RemoteAddr().String())
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("conn.Read err =", err)
			return
		}
		fmt.Print(string(buf[:n]))
	}
}
func main() {
	fmt.Println("伺服器開始監聽")
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("listen error =", err)
		return
	}
	defer listen.Close()
	for {
		fmt.Println("等待用戶端連接")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept error =", err)
		} else {
			fmt.Printf("Accept success conn = %v, ip = %v\n", conn, conn.RemoteAddr().String())
		}
		go proccess(conn)
	}
}
