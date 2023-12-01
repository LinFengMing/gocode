package main

import (
	"fmt"
)

type Person struct {
	Name string
}

func (p Person) speak() {
	fmt.Println(p.Name, " is a good man")
}

func (p Person) calculate() {
	res := 0
	for i := 1; i <= 1000; i++ {
		res += i
	}
	fmt.Println(p.Name, "計算的結果是 =", res)
}

func (p Person) calculate2(n int) {
	res := 0
	for i := 1; i <= n; i++ {
		res += i
	}
	fmt.Println(p.Name, "計算的結果是 =", res)
}

func (p Person) test() {
	p.Name = "jack"
	fmt.Println("test() name =", p.Name) // 輸出 jack
}

func (p Person) getSum(n1 int, n2 int) int {
	return n1 + n2
}

func main() {
	var p Person
	p.Name = "tom"
	p.test()
	fmt.Println("main() name =", p.Name) // 輸出 tom
	p.speak()
	p.calculate()
	p.calculate2(10)
	res := p.getSum(10, 20)
	fmt.Println("res =", res)
}
