package main

import (
	"fmt"
	"reflect"
)

func main() {
	var str string = "tom"
	fs := reflect.ValueOf(&str)
	fs.Elem().SetString("jack")
	fmt.Println(str)
}
