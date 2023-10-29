package main

import (
	"fmt"
)

// 長度是 array 類型的一部分
func test01(arr [3]int) {
	arr[0] = 44
	fmt.Println("test01 arr =", arr)
}

func test02(arr *[3]int) {
	fmt.Printf("test02 arr 指標的位置 = %p\n", &arr)
	(*arr)[0] = 88
	fmt.Println("test02 arr =", arr)
}

func main() {
	arr := [3]int{11, 22, 33}
	fmt.Printf("main arr 的位置 = %p\n", &arr)
	test01(arr)
	test02(&arr)
	fmt.Println("main arr =", arr)
}
