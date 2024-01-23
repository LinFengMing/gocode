package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("Error dialing", err)
		return
	}
	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading", err)
			return
		}
		line = strings.Trim(line, "\r\n")
		if line == "exit" {
			fmt.Println("Client exiting")
			break
		}
		n, err := conn.Write([]byte(line + "\n"))
		if err != nil {
			fmt.Println("Error writing", err)
			return
		}
		fmt.Printf("Send %d bytes to server\n", n)
	}
}
