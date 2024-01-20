package main

import (
	"fmt"
	"reflect"
)

func reflectTest01(b interface{}) {
	rType := reflect.TypeOf(b)
	fmt.Println("rType =", rType)
	rVal := reflect.ValueOf(b)
	fmt.Printf("rVal = %v, rVal type = %T\n", rVal, rVal)
	n2 := 2 + rVal.Int()
	fmt.Println("n2 =", n2)
	iV := rVal.Interface()
	num2 := iV.(int)
	fmt.Println("num2 =", num2)
}
func main() {
	var num int = 100
	reflectTest01(num)
}
