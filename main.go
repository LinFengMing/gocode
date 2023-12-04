package main

import (
	"fmt"
)

type Student struct {
	Name string
	Age  int
}

func main() {
	var stu1 = Student{"小明", 18}
	fmt.Println(stu1)
	stud2 := Student{"小红", 20}
	fmt.Println(stud2)
	var stu3 = Student{
		Name: "小剛",
		Age:  22,
	}
	fmt.Println(stu3)
	std4 := Student{
		Age:  24,
		Name: "小花",
	}
	fmt.Println(std4)
	// 返迴結構體指針
	var stu5 *Student = &Student{"小強", 26}
	fmt.Println(*stu5)
	stu6 := &Student{"小王", 24}
	fmt.Println(*stu6)
	var stud7 *Student = &Student{
		Name: "小張",
		Age:  28,
	}
	fmt.Println(*stud7)
	std8 := &Student{
		Age:  30,
		Name: "小李",
	}
	fmt.Println(*std8)
}
