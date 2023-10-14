package main

import (
	"fmt"
)

type mySum func(int, int) int

func sumOne(n1, n2 float32) float32 {
	fmt.Printf("n1 type = %T\n", n1)
	return n1 + n2
}

func sum(n1, n2 int) int {
	return n1 + n2
}

// func sum2(n1, n2, n3 int) int {
// 	return n1 + n2
// }

func myFunc(funcVar mySum, num1 int, num2 int) int {
	return funcVar(num1, num2)
}

func swap(n1 *int, n2 *int) {
	t := *n1
	*n1 = *n2
	*n2 = t
}

func main() {
	fmt.Println("sum =", sumOne(1, 2))
	a := sum
	// b := sum2
	fmt.Println(myFunc(a, 1, 2))
	// fmt.Println(myFunc(b, 1, 2)) // 錯誤，原因是類型不符合，因為不能將 sum2 賦於 mySum
	c := 10
	d := 20
	swap(&c, &d)
	fmt.Printf("c = %v, b = %v", c, d)
}
