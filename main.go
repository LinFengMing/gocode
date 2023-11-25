package main

import (
	"fmt"
)

type Person struct {
	Name   string
	Age    int
	Scores [5]float64
	ptr    *int
	slice  []int
	map1   map[string]string
}

type Monster struct {
	Name string
	Age  int
}

func main() {
	var p1 Person
	fmt.Println(p1)
	if p1.ptr == nil {
		fmt.Println("ok1")
	}
	if p1.slice == nil {
		fmt.Println("ok2")
	}
	if p1.map1 == nil {
		fmt.Println("ok3")
	}
	p1.ptr = new(int)
	p1.slice = make([]int, 10)
	p1.slice[0] = 100
	p1.map1 = make(map[string]string)
	p1.map1["key1"] = "Tom"
	fmt.Println(p1)
	var monster1 Monster
	monster1.Name = "牛魔王"
	monster1.Age = 500
	monster2 := monster1
	monster2.Name = "青牛精"
	fmt.Println(monster1)
	fmt.Println(monster2)
}
