package main

import (
	"fmt"
)

type AInterface interface {
	Say()
}
type Stu struct {
	Name string
}

func (s Stu) Say() {
	fmt.Println("Stu Say()...")
}

type integer int

func (i integer) Say() {
	fmt.Println("integer Say i = ", i)
}
func main() {
	var s Stu
	var a AInterface = s
	a.Say()
	var i integer = 10
	var a2 AInterface = i
	a2.Say()
}
