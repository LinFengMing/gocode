package main

import (
	"fmt"
	"strings"
)

func main() {
	// 字串去除頭尾空格
	str := strings.TrimSpace(" hello haha ")
	fmt.Printf("str = %q\n", str)
	// 字串頭尾去除字串
	str2 := strings.Trim("! hello! ", " !")
	fmt.Printf("str2 = %q\n", str2)
	// 字串頭去除字串
	str3 := strings.TrimLeft("!hello!", "!")
	fmt.Printf("str3 = %q\n", str3)
	str4 := strings.TrimRight("!hello!", "!")
	fmt.Printf("str4 = %q\n", str4)
	// 判斷字串是否用某字串開始
	b := strings.HasPrefix("ftp://192.168.10.1", "ftp")
	fmt.Printf("b = %v\n", b)
	// 判斷字串是否用某字串結束
	b2 := strings.HasSuffix("NLT_abc.jpg", "abc")
	fmt.Printf("b2 = %v\n", b2)
}
