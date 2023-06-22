package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str string = "ok"
	var b bool
	b, _ = strconv.ParseBool(str)
	fmt.Printf("b type %T b = %v \n", b, b)
	var str2 string = "hello"
	var n int64 = 11
	n, _ = strconv.ParseInt(str2, 10, 64)
	fmt.Printf("n type %T n = %v \n", n, n)
}
