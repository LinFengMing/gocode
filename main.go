package main

import (
	"fmt"
)

func test(slice []int) {
	slice[0] = 100
}

func main() {
	var arr []int = []int{1, 2, 3, 4, 5}
	slice := make([]int, 1)
	fmt.Println(slice)
	copy(slice, arr)
	fmt.Println(slice)
	var slice2 []int
	var arr2 [5]int = [...]int{1, 2, 3, 4, 5}
	slice2 = arr2[:]
	var slice3 = slice2
	slice3[0] = 10
	fmt.Println("slice3", slice3)
	fmt.Println("slice2", slice2)
	fmt.Println("arr2", arr2)
	slice4 := []int{1, 2, 3, 4}
	fmt.Println("slice4", slice4)
	test(slice4)
	fmt.Println("slice4", slice4)
}
