package main

import (
	"fmt"
)

type Cat struct {
	Name  string
	Age   int
	Color string
	Hobby string
}

func main() {
	var cat1 Cat
	fmt.Printf("cat1的位置 = %p\n", &cat1)
	cat1.Name = "小白"
	cat1.Age = 3
	cat1.Color = "白色"
	cat1.Hobby = "吃魚"
	fmt.Println("cat1 =", cat1)
	fmt.Println("貓的資料如下：")
	fmt.Println("name =", cat1.Name)
	fmt.Println("age =", cat1.Age)
	fmt.Println("color =", cat1.Color)
	fmt.Println("hobby =", cat1.Hobby)
}
