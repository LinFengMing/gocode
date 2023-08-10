package main

import (
	"fmt"
)

func main() {
	var n1 float64 = 1.2
	var n2 float64 = 1.3
	var operator byte = '+'
	var res float64
	switch operator {
	case '+':
		res = n1 + n2
	case '-':
		res = n1 - n2
	case '*':
		res = n1 * n2
	case '/':
		res = n1 / n2
	default:
		fmt.Println("操作符號錯誤")
	}
	fmt.Println("res=", res)
	n1 = 1.6
	n2 = 1.5
	operator = '-'
	switch operator {
	case '+':
		res = n1 + n2
	case '-':
		res = n1 - n2
	case '*':
		res = n1 * n2
	case '/':
		res = n1 / n2
	default:
		fmt.Println("操作符號錯誤")
	}
	fmt.Println("res=", res)
}
