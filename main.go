package main

import (
	"fmt"
)

func main() {
	// while
	var i int = 1
	for {
		if i > 10 {
			break
		}
		fmt.Println("hello,world!", i)
		i++
	}
	// do...while
	var j int = 1
	for {
		fmt.Println("hello,world!", j)
		j++
		if j > 10 {
			break
		}
	}
}
