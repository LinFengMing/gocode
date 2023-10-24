package main

import (
	"fmt"
	"time"
)

func main() {
	// 每 1 秒列印一個數字
	i := 0
	for {
		i++
		fmt.Println(i)
		time.Sleep(time.Second)
		if i == 10 {
			break
		}
	}
	// 每 0.1 秒列印一個數字
	j := 0
	for {
		j++
		fmt.Println(j)
		time.Sleep(time.Millisecond * 100)
		if j == 10 {
			break
		}
	}
	now := time.Now()
	fmt.Printf("unix = %v, unixnano = %v", now.Unix(), now.UnixNano())
}
