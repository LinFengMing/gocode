package main

import (
	"fmt"
)

func modify(map1 map[int]int) {
	map1[10] = 900
}

type Stu struct {
	Name string
	Age  int
}

func main() {
	map1 := make(map[int]int)
	map1[1] = 90
	map1[2] = 80
	map1[10] = 1
	map1[20] = 2
	modify(map1)
	fmt.Println(map1)
	// map value use struct
	students := make(map[string]Stu, 10)
	stu1 := Stu{"tom", 18}
	stu2 := Stu{"mary", 19}
	students["no1"] = stu1
	students["no2"] = stu2
	fmt.Println(students)
	for k, v := range students {
		fmt.Printf("學生的編號是 %v\n", k)
		fmt.Printf("學生的名字是 %v\n", v.Name)
		fmt.Printf("學生的年齡是 %v\n", v.Age)
		fmt.Println()
	}
}
