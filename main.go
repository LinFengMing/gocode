package main

import (
	"fmt"
)

func main() {
	num1 := 100
	fmt.Printf("num1的類型 = %T, num1的值 = %v, num1的位置 = %v\n", num1, num1, &num1)
	num2 := new(int)
	fmt.Printf("num2的類型 = %T, num2的值 = %v, num2的位置 = %v, num2 指向的值 = %v\n",
		num2, num2, &num2, *num2)
}
