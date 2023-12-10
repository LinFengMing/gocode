package main

import (
	"fmt"
)

type Monster struct {
	Name string
	Age  int
}
type E struct {
	Monster
	int
}
type Goods struct {
	Name  string
	Price float64
}
type Brand struct {
	Name    string
	Address string
}
type TV struct {
	Goods
	Brand
}

func main() {
	var e E
	e.Name = "牛魔王"
	e.Age = 500
	e.int = 20
	fmt.Println(e)
	tv := TV{Goods{"電視機001", 5000.99}, Brand{"索尼", "日本"}}
	fmt.Println(tv.Goods.Name)
}
