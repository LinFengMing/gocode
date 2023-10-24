package main

import (
	"fmt"
	"strconv"
	"time"
)

func test() {
	str := ""
	for i := 0; i < 100000; i++ {
		str += "hello" + strconv.Itoa(i)
	}
}

func main() {
	start := time.Now().Unix()
	test()
	end := time.Now().Unix()
	fmt.Printf("執行 test() 使用時間為 %v 秒\n", end-start)
}
