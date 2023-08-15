package main

import (
	"fmt"
)

func test(n1 int) {
	n1 = n1 + 1
	fmt.Println("test n1 =", n1)
}

func getSum(n1 int, n2 int) int {
	sum := n1 + n2
	fmt.Println("getSum sum =", sum)
	return sum
}
func getSumAndSun(n1 int, n2 int) (int, int) {
	sum := n1 + n2
	sub := n1 - n2
	return sum, sub
}

func main() {
	n1 := 10
	test(n1)
	fmt.Println("main n1 =", n1)
	sum := getSum(10, 20)
	fmt.Println("main sum =", sum)
	res1, res2 := getSumAndSun(1, 2)
	fmt.Printf("res1 = %v, res2 = %v\n", res1, res2)
	_, res3 := getSumAndSun(3, 9)
	fmt.Println("res3 =", res3)
}
