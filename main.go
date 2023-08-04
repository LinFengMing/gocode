package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var count int = 0
	rand.New(rand.NewSource(time.Now().UnixNano()))
	for {
		n := rand.Intn(100) + 1
		fmt.Println(n)
		count++
		if n == 99 {
			break
		}
	}
	fmt.Println("生成 99 一共使用了", count, "次")
}
