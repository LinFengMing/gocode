package main

import (
	"fmt"
)

func main() {
	var n int = 10
	fmt.Println("ok1")
	if n > 20 {
		goto label
	}
	fmt.Println("ok2")
	fmt.Println("ok3")
	fmt.Println("ok4")
label:
	fmt.Println("ok5")
	fmt.Println("ok6")
	fmt.Println("ok7")
	var m int = 30
	fmt.Println("ok1")
	if m > 20 {
		return
	}
	fmt.Println("ok2")
	fmt.Println("ok3")
	fmt.Println("ok4")
	fmt.Println("ok5")
	fmt.Println("ok6")
	fmt.Println("ok7")
}
