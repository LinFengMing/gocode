package main

import (
	"fmt"
)

type Person struct {
	Name string
}

func (p Person) test() {
	p.Name = "jack"
	fmt.Println("test() name =", p.Name) // 輸出 jack
}

func main() {
	var p Person
	p.Name = "tom"
	p.test()
	fmt.Println("main() name =", p.Name) // 輸出 tom
}
