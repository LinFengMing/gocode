package main

import (
	"fmt"
)

func main() {
	var myChars [26]byte
	for i := 0; i < 26; i++ {
		myChars[i] = 'A' + byte(i)
	}
	for i := 0; i < 26; i++ {
		fmt.Printf("%c\n", myChars[i])
	}

	var intArr [5]int = [...]int{1, -1, 9, 90, 11}
	maxVal := intArr[0]
	maxIndex := 0
	for i := 1; i < len(intArr); i++ {
		if maxVal < intArr[i] {
			maxVal = intArr[i]
			maxIndex = i
		}
	}
	fmt.Printf("maxVal = %v, maxIndex = %v", maxVal, maxIndex)
}
