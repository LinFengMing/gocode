package main

import (
	"fmt"
)

type Student struct {
	name   string
	gender string
	age    int
	id     int
	score  float64
}

func (s *Student) say() string {
	infoStr := fmt.Sprintf("student 的資訊 name = [%v], gender = [%v], age = [%v], id = [%v], score = [%v]",
		s.name, s.gender, s.age, s.id, s.score)
	return infoStr
}

type Box struct {
	len    float64
	width  float64
	height float64
}

type Visitor struct {
	Name string
	Age  int
}

func (v *Visitor) showPrice() {
	if v.Age >= 90 || v.Age <= 8 {
		fmt.Println("考慮安全問題，不建議入場")
		return
	}
	if v.Age >= 18 {
		fmt.Printf("年齡為%v歲的%s，票價為20元\n", v.Age, v.Name)
	} else {
		fmt.Printf("年齡為%v歲的%s，免費\n", v.Age, v.Name)
	}
}

func (b *Box) getVolumn() float64 {
	return b.len * b.width * b.height
}

func main() {
	s := Student{"tom", "male", 18, 1000, 99.98}
	fmt.Println(s.say())
	b := Box{1.1, 2.2, 3.3}
	fmt.Printf("體積 = %.2f\n", b.getVolumn())
	var v Visitor
	for {
		fmt.Println("請輸入名字")
		fmt.Scanln(&v.Name)
		if v.Name == "n" {
			fmt.Println("退出")
			break
		}
		fmt.Println("請輸入年齡")
		fmt.Scanln(&v.Age)
		v.showPrice()
	}
}
