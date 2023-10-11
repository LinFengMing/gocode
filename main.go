package main

import (
	"fmt"
)

type myFunType func(int, int) int

func getSum(n1 int, n2 int) int {
	return n1 + n2
}

func myFun(funvar func(int, int) int, num1 int, num2 int) int {
	return funvar(num1, num2)
}

func myFun2(funvar myFunType, num1 int, num2 int) int {
	return funvar(num1, num2)
}

func getSumAndSub(n1 int, n2 int) (sum int, sub int) {
	sum = n1 + n2
	sub = n1 - n2
	return
}

func main() {
	a := getSum
	fmt.Printf("a 的類型 %T, getSum 是 %T\n", a, getSum)
	result := a(10, 40)
	fmt.Println("result =", result)
	result2 := myFun(getSum, 50, 60)
	fmt.Println("result2 =", result2)
	type myInt int // 在 go 中 myInt 和 int 雖然都是 int 類型，但是 go 認為 myInt 和 int 是兩個類型
	var num1 myInt
	var num2 int
	num1 = 40
	fmt.Println("num1 =", num1)
	num2 = int(num1)
	fmt.Println("num2 =", num2)
	result3 := myFun2(getSum, 50, 60)
	fmt.Println("result3 =", result3)
	c, d := getSumAndSub(1, 2)
	fmt.Printf("a = %v, b = %v\n", c, d)
}
