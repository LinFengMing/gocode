package main

import (
	"fmt"
)

func main() {
	var slice []float64 = make([]float64, 5, 10)
	slice[1] = 10
	slice[3] = 20
	fmt.Println("slice =", slice)
	fmt.Println("slice size =", len(slice))
	fmt.Println("slice cap =", cap(slice))
	var slice2 []string = []string{"Tom", "Jack", "Marry"}
	fmt.Println("slice2 =", slice2)
	fmt.Println("slice2 size =", len(slice2))
	fmt.Println("slice2 cap =", cap(slice2))
}
