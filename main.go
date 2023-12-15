package main

import (
	"fmt"
)

type Monkey struct {
	Name string
}
type BirdAble interface {
	flying()
}
type FishAble interface {
	swimming()
}

func (m *Monkey) climbing() {
	fmt.Println(m.Name, "生來會爬樹")
}

type LittleMonkey struct {
	Monkey
}

func (l *LittleMonkey) flying() {
	fmt.Println(l.Name, "學會飛翔")
}
func (l *LittleMonkey) swimming() {
	fmt.Println(l.Name, "學會游泳")
}
func main() {
	monkey := LittleMonkey{Monkey{"悟空"}}
	monkey.climbing()
	monkey.flying()
	monkey.swimming()
}
