package main

import (
	"fmt"
)

func main() {
	var age int = 10
	switch {
	case age == 10:
		fmt.Println("age == 10")
	case age == 20:
		fmt.Println("age == 20")
	default:
		fmt.Println("沒有配對到")
	}
	var score int = 90
	switch {
	case score > 90:
		fmt.Println("成績優秀")
	case score >= 70 && score <= 90:
		fmt.Println("成績優良")
	case score >= 60 && score < 70:
		fmt.Println("成績及格")
	default:
		fmt.Println("成績不及格")
	}
	switch grade := 60; {
	case grade > 90:
		fmt.Println("成績優秀")
	case grade >= 70 && grade <= 90:
		fmt.Println("成績優良")
	case grade >= 60 && grade < 70:
		fmt.Println("成績及格")
	default:
		fmt.Println("成績不及格")
	}
	var num int = 10
	switch num {
	case 10:
		fmt.Println("ok1")
		fallthrough
	case 20:
		fmt.Println("ok2")
		fallthrough
	case 30:
		fmt.Println("ok3")
	default:
		fmt.Println("沒有配對到")
	}
	var x interface{}
	var y = 10.0
	x = y
	switch i := x.(type) {
	case nil:
		fmt.Printf("x 的類型是 %T", i)
	case int:
		fmt.Printf("x 的類型是 int")
	case float64:
		fmt.Printf("x 的類型是 float64")
	case func(int) float64:
		fmt.Printf("x 的類型是 fun(int)")
	case bool, string:
		fmt.Printf("x 的類型是 bool 或 string")
	default:
		fmt.Printf("x 的類型是未知")
	}
}
