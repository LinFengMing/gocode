package main

import (
	"fmt"
)

func test() {
	defer func() {
		err := recover() // recover() 可以抓取異常
		if err != nil {
			fmt.Println("err =", err)
		}
	}()
	num1 := 10
	num2 := 0
	result := num1 / num2
	fmt.Println("result =", result)
}

func main() {
	test()
	fmt.Println("hihi")
}
