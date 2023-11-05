package main

import (
	"fmt"
)

func main() {
	str := "hello@world"
	slice := str[6:]
	fmt.Println("slice =", slice)
	arr := []byte(str)
	arr[0] = 'H'
	str = string(arr)
	fmt.Println("str =", str)
	// []rune 是按字元處理，可以使用中文
	arr1 := []rune(str)
	arr1[0] = '哈'
	str = string(arr1)
	fmt.Println("str =", str)
}
