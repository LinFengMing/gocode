package main

import (
	"fmt"
)

type A struct {
	Num int
}

type B struct {
	Num int
}

func main() {
	var a A
	var b B
	a = A(b) // 可以轉換，但 struct Fields 必須一樣
	fmt.Println(a, b)
}
