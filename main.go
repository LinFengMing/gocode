package main

import (
	"fmt"
)

func main() {
lable1:
	for i := 0; i < 4; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				break lable1
			}
			fmt.Println("j =", j)
		}
	}
}
