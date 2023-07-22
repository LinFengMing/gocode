package main

import (
	"fmt"
)

func main() {
	var str string = "hello,world!台北"
	str2 := []rune(str)
	for i := 0; i < len(str2); i++ {
		fmt.Printf("%c\n", str2[i])
	}
	str = "abc~ok新北"
	for index, val := range str {
		fmt.Printf("index = %d, val= %c\n", index, val)
	}
}
