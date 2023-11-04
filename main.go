package main

import (
	"fmt"
)

func main() {
	var arr [5]int = [...]int{10, 20, 30, 40, 50}
	slice := arr[:] // arr[0:len(arr)]
	for i, v := range slice {
		fmt.Printf("i = %v, v = %v\n", i, v)
	}
	slice2 := slice[1:2]
	fmt.Println("slice2", slice2)
}
