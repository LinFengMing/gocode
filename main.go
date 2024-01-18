package main

import (
	"fmt"
)

func main() {
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}
	stringsChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		stringsChan <- "hello" + fmt.Sprintf("%d", i)
	}
	for {
		select {
		case v := <-intChan:
			fmt.Printf("從 intChan 讀取的數據 %d\n", v)
		case v := <-stringsChan:
			fmt.Printf("從 stringsChan讀取的數據 %s\n", v)
		default:
			fmt.Printf("都取不到了\n")
			return
		}
	}
}
