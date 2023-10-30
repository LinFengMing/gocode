package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var intArr [5]int
	len := len(intArr)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < len; i++ {
		intArr[i] = r.Intn(100)
	}
	fmt.Println("交換前", intArr)
	temp := 0
	for i := 0; i < len/2; i++ {
		temp = intArr[len-1-i]
		intArr[len-1-i] = intArr[i]
		intArr[i] = temp
	}
	fmt.Println("交換後", intArr)
}
