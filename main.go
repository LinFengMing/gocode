package main

import "fmt"

func main() {
	var intChan chan int
	intChan = make(chan int, 3)
	fmt.Printf("intChan 的值 = %v, intChan本身的位址 = %p\n", intChan, &intChan)
	intChan <- 10
	num := 211
	intChan <- num
	intChan <- 50
	// intChan <- 666 // deadlock
	fmt.Printf("intChan len = %v, cap = %v\n", len(intChan), cap(intChan)) // 3, 3
	var num2 int
	num2 = <-intChan
	fmt.Println("num2 = ", num2)
	fmt.Printf("intChan len = %v, cap = %v\n", len(intChan), cap(intChan)) // 2, 3
	num3 := <-intChan
	num4 := <-intChan
	fmt.Println("num3 = ", num3)
	fmt.Println("num4 = ", num4)
	fmt.Printf("intChan len = %v, cap = %v\n", len(intChan), cap(intChan)) // 0, 3
	// num5 := <-intChan // deadlock
}
