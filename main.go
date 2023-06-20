package main

import (
	"fmt"
)

func main() {
	// 字串一旦賦值了，字串就不能修改了：在 GO 中字串是不能變的
	var address string = "北京長城 110 hello world"
	fmt.Println("address =", address)
}
