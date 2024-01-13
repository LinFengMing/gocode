package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	myMap = make(map[int64]int64, 10)
	// lock 是一個全域互斥鎖
	lock sync.Mutex
)

func test(n int) {
	res := int64(1)
	for i := 1; i <= n; i++ {
		res *= int64(i)
	}
	lock.Lock()
	myMap[int64(n)] = res
	lock.Unlock()
}

func main() {
	for i := 1; i <= 20; i++ {
		go test(i)
	}
	time.Sleep(time.Second * 5)
	lock.Lock()
	for i, v := range myMap {
		fmt.Printf("map[%d] = %d\n", i, v)
	}
	lock.Unlock()
}
