package main

import (
	"fmt"
)

type MenthodUtils struct {
}

func (m *MenthodUtils) JudgnNum(num int) {
	if num%2 == 0 {
		fmt.Println(num, "是偶數")
	} else {
		fmt.Println(num, "是奇數")
	}
}

func (mu *MenthodUtils) Print(n int, m int, key string) {
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			fmt.Print(key)
		}
		fmt.Println()
	}
}

type Calcuator struct {
	Num1 float64
	Num2 float64
}

func (c *Calcuator) GetSum() float64 {
	return c.Num1 + c.Num2
}

func (c *Calcuator) GetSub() float64 {
	return c.Num1 - c.Num2
}

func (c *Calcuator) GetMul() float64 {
	return c.Num1 * c.Num2
}

func (c *Calcuator) GetDiv() float64 {
	return c.Num1 / c.Num2
}

func (c *Calcuator) GetRes(operator byte) float64 {
	res := 0.0
	switch operator {
	case '+':
		res = c.Num1 + c.Num2
	case '-':
		res = c.Num1 - c.Num2
	case '*':
		res = c.Num1 * c.Num2
	case '/':
		res = c.Num1 / c.Num2
	default:
		fmt.Println("運算符號錯誤")
	}
	return res
}

func main() {
	var mu MenthodUtils
	mu.JudgnNum(10)
	fmt.Println()
	mu.Print(5, 5, "#")
	var c Calcuator
	c.Num1 = 1.2
	c.Num2 = 2.2
	fmt.Printf("sum = %v\n", fmt.Sprintf("%.2f", c.GetSum()))
	fmt.Printf("sub = %v\n", fmt.Sprintf("%.2f", c.GetSub()))
	fmt.Printf("mul = %v\n", fmt.Sprintf("%.2f", c.GetMul()))
	fmt.Printf("sub = %v\n", fmt.Sprintf("%.2f", c.GetDiv()))
	fmt.Printf("res = %v\n", fmt.Sprintf("%.2f", c.GetRes('+')))
}
