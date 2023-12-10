package main

import (
	"fmt"
)

type A struct {
	Name string
	Age  int
}
type B struct {
	Name  string
	Score float64
}
type C struct {
	A
	B
}
type D struct {
	a A
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
type TV2 struct {
	*Goods
	*Brand
}

func main() {
	var c C
	c.A.Name = "tom"
	c.B.Name = "jack"
	fmt.Println(c.A.Name) // tom
	fmt.Println(c.B.Name) // jack
	var d D
	d.a.Name = "john"
	fmt.Println(d.a.Name) // john
	tv := TV{Goods{"電視機001", 5000.99}, Brand{"海爾", "山東"}}
	fmt.Println(tv)
	tv2 := TV{
		Goods{
			Name:  "電視機002",
			Price: 6000.99,
		},
		Brand{
			Name:    "夏普",
			Address: "北京",
		},
	}
	fmt.Println(tv2)
	tv3 := TV2{&Goods{"電視機003", 7000.99}, &Brand{"東芝", "日本"}}
	fmt.Println(*tv3.Goods, *tv3.Brand)
	tv4 := TV2{
		&Goods{
			Name:  "電視機004",
			Price: 8000.99,
		},
		&Brand{
			Name:    "索尼",
			Address: "日本",
		},
	}
	fmt.Println(*tv4.Goods, *tv4.Brand)
}
