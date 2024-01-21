package main

import (
	"fmt"
	"reflect"
)

func reflect01(b interface{}) {
	rVal := reflect.ValueOf(b)
	fmt.Println("rVal kind =", rVal.Kind())
	rVal.Elem().SetInt(20)
}
func main() {
	var num int = 100
	reflect01(&num)
	fmt.Println("num =", num)
}
