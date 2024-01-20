package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age  int
}

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
func reflectTest02(b interface{}) {
	rType := reflect.TypeOf(b)
	fmt.Println("rType =", rType)
	rVal := reflect.ValueOf(b)
	fmt.Printf("rVal = %v, rVal type = %T\n", rVal, rVal)
	fmt.Printf("kind = %v, kind = %v\n", rType.Kind(), rVal.Kind())
	iV := rVal.Interface()
	fmt.Printf("iV = %v, iV type = %T\n", iV, iV)
	stu, ok := iV.(Student)
	if ok {
		fmt.Printf("stu.Name = %v\n", stu.Name)
	}
}
func main() {
	var num int = 100
	reflectTest01(num)
	stu := Student{
		Name: "tom",
		Age:  20,
	}
	reflectTest02(stu)
}
