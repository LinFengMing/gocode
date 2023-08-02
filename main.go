package main

import (
	"fmt"
)

func main() {
	var layer int
	fmt.Println("輸入層數")
	fmt.Scanln(&layer)
	for i := 1; i <= layer; i++ {
		for k := 1; k <= layer-i; k++ {
			fmt.Print(" ")
		}
		for j := 1; j <= 2*i-1; j++ {
			if j == 1 || j == 2*i-1 || i == layer {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
