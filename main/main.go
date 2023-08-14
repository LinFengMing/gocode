package main

import (
	"fmt"
	"gocode/utils"
)

func main() {
	var n1 float64 = 1.2
	var n2 float64 = 1.3
	var operator byte = '+'
	result := utils.Cal(n1, n2, operator)
	fmt.Println("res=", result)
	n1 = 1.6
	n2 = 1.5
	operator = '-'
	result = utils.Cal(n1, n2, operator)
	fmt.Printf("res=%.2f", result)
}
