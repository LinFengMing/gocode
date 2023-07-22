package main

import (
	"fmt"
)

func main() {
	for i := 1; i < 11; i++ {
		fmt.Println("你好", i)
	}
	j := 1
	for j < 11 {
		fmt.Println("你好~", j)
		j++
	}
	k := 1
	for {
		if k < 11 {
			fmt.Println("ok")
		} else {
			break
		}
		k++
	}
}
