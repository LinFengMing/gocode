package main

import (
	"fmt"
)

func main() {
	var intArr [5]int = [...]int{1, -1, 9, 90, 12}
	sum := 0
	for _, v := range intArr {
		sum += v
	}
	fmt.Printf("sum = %v, 平均值 = %v", sum, float64(sum)/float64(len(intArr)))
}
