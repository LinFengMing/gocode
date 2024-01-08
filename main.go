package main

import (
	"flag"
	"fmt"
)

func main() {
	var user string
	var pwd string
	var host string
	var port int
	flag.StringVar(&user, "u", "", "用户名，默認為空")
	flag.StringVar(&pwd, "pwd", "", "密碼，默認為空")
	flag.StringVar(&host, "h", "localhost", "主機名，默認為localhost")
	flag.IntVar(&port, "port", 3306, "端口號，默認為3306")
	// 重要，轉換，必須使用該方法
	flag.Parse()
	fmt.Printf("user = %v, pwd = %v, host = %v, port = %v", user, pwd, host, port)
}
