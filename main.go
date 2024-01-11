package main

import (
	"fmt"
	"strconv"
	"time"
)

func goroutine() {
	for i := 1; i <= 10; i++ {
		fmt.Println("hello, word " + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
	}
}
func main() {
	go goroutine() // 開啟一個協程
	for i := 1; i <= 10; i++ {
		fmt.Println("hello, golang " + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
	}
}
