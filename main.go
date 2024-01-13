package main

import (
	"fmt"
)

var (
	myMap = make(map[int]int, 10)
)

func test(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	myMap[n] = res // 併發寫入，fatal error: concurrent map writes
}

func main() {
	for i := 1; i <= 200; i++ {
		go test(i)
	}
	for i, v := range myMap {
		fmt.Printf("map[%d] = %d\n", i, v)
	}
}
