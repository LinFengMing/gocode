package main

import (
	"fmt"
)

func BinarySearch(arr *[6]int, leftIndex int, rightIndex int, findVal int) {
	if leftIndex > rightIndex {
		fmt.Println("找不到")
		return
	}

	middleIndex := (leftIndex + rightIndex) / 2

	if (*arr)[middleIndex] > findVal {
		BinarySearch(arr, leftIndex, middleIndex-1, findVal)
	} else if (*arr)[middleIndex] < findVal {
		BinarySearch(arr, middleIndex+1, rightIndex, findVal)
	} else {
		fmt.Printf("找到了，index = %v\n", middleIndex)
	}
}

func main() {
	arr := [6]int{1, 8, 10, 89, 1000, 1234}
	BinarySearch(&arr, 0, len(arr)-1, 1000)
}
