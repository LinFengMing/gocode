package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("Error dialing", err)
		return
	}
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading", err)
		return
	}
	n, err := conn.Write([]byte(line))
	if err != nil {
		fmt.Println("Error writing", err)
		return
	}
	fmt.Printf("Send %d bytes to server\n", n)
}
