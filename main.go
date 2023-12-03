package main

import (
	"fmt"
)

type integer int

func (i integer) print() {
	fmt.Println("i =", i)
}

func (i *integer) change() {
	*i = *i + 1
}

type Student struct {
	Name string
	Age  int
}

func (s *Student) String() string {
	str := fmt.Sprintf("Name = [%v], Age = [%v]\n", s.Name, s.Age)
	return str
}

func main() {
	var i integer = 10
	i.change()
	i.print()
	stu := Student{
		Name: "Tom",
		Age:  20,
	}
	fmt.Println(&stu)
}
