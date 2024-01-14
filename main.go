package main

import "fmt"

func main() {
	intChan := make(chan int, 3)
	intChan <- 100
	intChan <- 200
	close(intChan)
	// intChan <- 300 // panic: send on closed channel
	n1 := <-intChan
	fmt.Println("n1 =", n1)
	intChan2 := make(chan int, 100)
	for i := 0; i < 100; i++ {
		intChan2 <- i * 2
	}
	close(intChan2)
	for v := range intChan2 {
		fmt.Println("v =", v)
	}
}
