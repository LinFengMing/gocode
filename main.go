package main

import (
	"fmt"
)

func main() {
	var number int = 9
	for i := 1; i <= number; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v * %v = %v \t", j, i, j*i)
		}
		fmt.Println()
	}
}
