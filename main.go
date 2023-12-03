package main

import (
	"fmt"
)

type MenthodUtils struct {
}

func (m MenthodUtils) Print() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 8; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func (m MenthodUtils) Print2(m1 int, m2 int) {
	for i := 0; i < m1; i++ {
		for j := 0; j < m2; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func (m MenthodUtils) Area(len float64, width float64) float64 {
	return len * width
}

func main() {
	var m = MenthodUtils{}
	m.Print()
	fmt.Println()
	m.Print2(2, 3)
	fmt.Println()
	fmt.Println(m.Area(2.0, 3.0))
}
