package main

import (
	"fmt"
)

func fbn(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return fbn(n-1) + fbn(n-2)
	}
}

func f(n int) int {
	if n == 1 {
		return 3
	} else {
		return 2*f(n-1) + 1
	}
}

func main() {
	fmt.Println("result =", fbn(3))
	fmt.Println("result =", fbn(4))
	fmt.Println("result =", fbn(5))
	fmt.Println("result =", fbn(6))
	fmt.Println("result =", f(1))
	fmt.Println("result =", f(5))
}
