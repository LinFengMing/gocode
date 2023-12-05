package main

import (
	"fmt"
	"gocode/model"
)

func main() {
	var stu = model.NewStudent("Tom", 88.8)
	fmt.Println(*stu)
	fmt.Println(stu.Name)
	fmt.Println(stu.GetScore())
}
