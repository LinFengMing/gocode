package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num1 int = 99
	var num2 float64 = 23.456
	var b bool = true
	var myChar byte = 'h'
	var str string
	str = fmt.Sprintf("%d", num1)
	fmt.Printf("str type %T str = %q \n", str, str)
	str = fmt.Sprintf("%f", num2)
	fmt.Printf("str type %T str = %q \n", str, str)
	str = fmt.Sprintf("%t", b)
	fmt.Printf("str type %T str = %q \n", str, str)
	str = fmt.Sprintf("%c", myChar)
	fmt.Printf("str type %T str = %q \n", str, str)
	str = strconv.FormatInt(int64(num1), 10)
	fmt.Printf("str type %T str = %q \n", str, str)
	str = strconv.FormatFloat(num2, 'f', 10, 64)
	fmt.Printf("str type %T str = %q \n", str, str)
	str = strconv.FormatBool(b)
	fmt.Printf("str type %T str = %q \n", str, str)
	var num5 int64 = 4567
	str = strconv.Itoa(int(num5))
	fmt.Printf("str type %T str = %q \n", str, str)
}
