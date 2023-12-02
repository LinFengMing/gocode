package main

import (
	"fmt"
)

type Circle struct {
	radius float64
}

func (c *Circle) area() float64 {
	fmt.Printf("c 指向的位址 = %p\n", c)
	c.radius = 10.0
	return 3.14 * (*c).radius * (*c).radius
}

func main() {
	var c Circle
	fmt.Printf("main c 位址 = %p\n", &c)
	c.radius = 5.0
	res := (&c).area()
	fmt.Println("Area of Circle(c) is", res)
	fmt.Println("c.radius =", c.radius)
}
