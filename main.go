package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func test01(p Person) {
	fmt.Println(p.Name)
}

func test02(p *Person) {
	fmt.Println(p.Name)
}

func (p Person) test03() { // 這裡是值拷貝
	p.Name = "jack"
	fmt.Println(p.Name) // jack
}

func (p *Person) test04() { // 這裡是位址拷貝
	p.Name = "mary"
	fmt.Println(p.Name) // mary
}

func main() {
	p := Person{"tom", 20}
	test01(p)
	test02(&p)
	p.test03()
	fmt.Println(p.Name) // tom
	(&p).test03()       // 這裡值拷貝
	fmt.Println(p.Name) // tom
	(&p).test04()       // 這裡是位址拷貝
	fmt.Println(p.Name) // mary
}
