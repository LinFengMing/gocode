package main

import (
	"fmt"
)

func main() {
	// map 定義不會分配記憶體空間，必須使用 make 來分配記憶體空間
	var m map[string]string
	m = make(map[string]string, 10)
	m["no1"] = "jiro"
	m["no2"] = "teru"
	fmt.Println(m)
}
