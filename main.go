package main

import (
	"fmt"
)

func cal(n1 int, n2 int) (sum int, sub int) {
	sum = n1 + n2
	sub = n1 - n2
	return
}

func sum(n1 int, args ...int) int {
	sum := n1
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum
}

func main() {
	result1, _ := cal(10, 20)
	fmt.Printf("result1 = %d\n", result1)
	result2 := sum(10, 0, -1, 90, 10, 100)
	fmt.Printf("result2 = %d\n", result2)
}
