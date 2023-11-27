package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	var p1 Person
	p1.Name = "John"
	p1.Age = 25
	fmt.Println(p1)
	var p2 = Person{"John", 25}
	fmt.Println(p2)
	var p3 *Person = new(Person)
	(*p3).Name = "John"
	(*p3).Age = 25
	fmt.Println(*p3)
	var p4 *Person = &Person{"John", 25}
	fmt.Println(*p4)
}
