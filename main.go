package main

import (
	"fmt"
)

func main() {
	var n1 int32 = 30
	var n2 int32 = 20
	if n1+n2 >= 50 {
		fmt.Println("Hello World")
	}
	var n3 float64 = 10.1
	var n4 float64 = 19.9
	if n3 > 10.0 && n4 < 20.0 {
		fmt.Println("和 =", (n3 + n4))
	}
	var n5 int32 = 7
	var n6 int32 = 8
	if (n5+n6)%3 == 0 && (n5+n6)%5 == 0 {
		fmt.Println("能被3跟5整除")
	}
	var year = 2020
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		fmt.Println(year, "是閏年")
	}
}
