package main

import (
	"fmt"
)

func main() {

	// 如果還有97天放假，請問幾個星期又幾天
	var weeks int = 97 / 7
	var days int = 97 % 7
	fmt.Printf("%d 個星期又 %d 天\n", weeks, days)
	// 攝氏溫度 = 5 / 9 * (華氏溫度 - 100)
	var fahrenheit float32 = 134.2
	var celsius float32 = 5.0 / 9 * (fahrenheit - 100)
	fmt.Printf("%v 度 C 等於 %v 度 F", celsius, fahrenheit)
}
