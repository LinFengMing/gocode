package main

import (
	"fmt"
)

func main() {
	// 印出 1 ~ 100 之間所有9的倍數的整數的個數和總和
	var count uint64 = 0
	var sum uint64 = 0
	var i uint64 = 1
	for ; i < 101; i++ {
		if i%9 == 0 {
			count++
			sum += i
		}
	}
	fmt.Printf("個數 = %v, 總和 = %v\n", count, sum)
	var n int = 6
	for i := 0; i <= n; i++ {
		fmt.Printf("%v + %v = %v\n", i, n-i, n)
	}
}
