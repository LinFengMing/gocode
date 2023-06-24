package main

import (
	"fmt"
)

func main() {
	var num int = 9
	fmt.Printf("i address = %v \n", &num)
	var ptr *int = &num
	*ptr = 10
	fmt.Printf("num = %v \n", num)
	var a int = 300
	var b int = 400
	var ptr2 *int = &a
	*ptr2 = 100
	ptr2 = &b
	*ptr2 = 200
	fmt.Printf("a = %d, b = %d, ptr2 = %d \n", a, b, *ptr2)
}
