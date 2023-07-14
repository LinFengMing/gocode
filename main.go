package main

import (
	"fmt"
	"math"
)

func main() {
	var z bool = true
	if z == false {
		fmt.Println("a")
	} else if z {
		fmt.Println("b")
	} else if !z {
		fmt.Println("c")
	} else {
		fmt.Println("d")
	}
	var a float64 = 3.0
	var b float64 = 100.0
	var c float64 = 6.0
	m := b*b - 4*a*c
	if m > 0 {
		x1 := (-b + math.Sqrt(m)) / 2 * a
		x2 := (-b - math.Sqrt(m)) / 2 * a
		fmt.Printf("x1 = %v, X2 = %v\n", x1, x2)
	} else if m == 0 {
		x1 := (-b + math.Sqrt(m)) / 2 * a
		fmt.Printf("x1 = %v\n", x1)
	} else {
		fmt.Println("無解")
	}
	var height int32
	var money float32
	var handsome bool
	fmt.Println("請輸入身高：")
	fmt.Scanln(&height)
	fmt.Println("請輸入存款：")
	fmt.Scanln(&money)
	fmt.Println("請輸入帥嗎？")
	fmt.Scanln(&handsome)
	if height > 180 && money > 1.0 && handsome {
		fmt.Println("我一定嫁給他")
	} else if height > 180 || money > 1.0 || handsome {
		fmt.Println("嫁吧，比上不足，比下有餘")
	} else {
		fmt.Println("不嫁")
	}
}
