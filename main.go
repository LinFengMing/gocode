package main

import (
	"fmt"
)

func main() {
	var c4 int = 22269 // 国
	fmt.Printf("c4 = %c\n", c4)
	// 字符類型是可以進行運算的，等於是一個整數，運算時是按照碼值運行
	var n1 = 10 + 'a' // 10 + 97 = 107
	fmt.Println("n1 =", n1)
}
