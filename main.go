package main

import (
	"fmt"
)

type AInterface interface {
	Say()
}
type BInterface interface {
	Hello()
}
type CInterface interface {
	test01()
}
type DInterface interface {
	test02()
}
type EInterface interface {
	CInterface
	DInterface
	test03()
}
type Stu struct {
}

func (stu Stu) test01() {
	fmt.Println("test01()...")
}
func (stu Stu) test02() {
	fmt.Println("test02()...")
}
func (stu Stu) test03() {
	fmt.Println("test03()...")
}

type Monster struct {
}

func (m Monster) Hello() {
	fmt.Println("Monster Hello()...")
}
func (m Monster) Say() {
	fmt.Println("Monster Say()...")
}

type T interface {
}

func main() {
	var monster Monster
	var a AInterface = monster
	var b BInterface = monster
	a.Say()
	b.Hello()
	var stu Stu
	var e EInterface = stu
	e.test01()
	e.test02()
	e.test03()
	var t T = stu
	fmt.Println(t)
	var t2 interface{} = stu
	fmt.Println(t2)
	var num1 float32 = 8.8
	t2 = num1
	fmt.Println(t2)
}
