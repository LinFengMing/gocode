package main

import (
	"fmt"
)

func main() {
	var intArr [3]int // 預設值為 0, int64 間隔 8 個字元
	fmt.Println(intArr)
	intArr[0] = 10
	intArr[1] = 20
	intArr[2] = 30
	fmt.Println(intArr)
	fmt.Printf("intArr 的位置 = %p, intArr[0] 的位置 = %p, intArr[1] 的位置 = %p, intArr[2] 的位置 = %p\n",
		&intArr, &intArr[0], &intArr[1], &intArr[2])
}
