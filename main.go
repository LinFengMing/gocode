package main

import (
	"fmt"
)

func main() {
	var intArr [5]int = [...]int{1, 22, 33, 66, 99}
	// 使用 intArr[1:3] index 1 到 3 不包含 3
	slice := intArr[1:3]
	fmt.Println("intArr =", intArr)
	fmt.Println("slice =", slice)
	fmt.Println("slice 的元素數 =", len(slice))
	fmt.Println("slice 的容量 =", cap(slice))
}
