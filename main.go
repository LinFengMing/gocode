package main

import "fmt"

func main() {
	// int8 -128 ~ 127
	var i int8 = 127
	fmt.Println("i =", i)
	// uint8 0 ~ 255
	var k uint8 = 255
	fmt.Println("k =", k)
	// int , unit , rune , byte
	var a int = 8900
	fmt.Println("a =", a)
	var b uint = 1
	fmt.Println("b =", b)
	var c rune = 999999999
	fmt.Println("c =", c)
	var d byte = 255
	fmt.Println("d =", d)
}
