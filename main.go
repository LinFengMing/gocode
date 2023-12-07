package main

import (
	"fmt"
)

type Pupil struct {
	Name  string
	Age   int
	Score int
}

func (p *Pupil) ShowInfo() {
	fmt.Printf("Name = %v Age = %v Score = %v\n", p.Name, p.Age, p.Score)
}

func (p *Pupil) SetScore(score int) {
	p.Score = score
}

func (p *Pupil) Testing() {
	fmt.Println("PrimarySchoolStydent is testing...")
}

type Graduate struct {
	Name  string
	Age   int
	Score int
}

func (g *Graduate) ShowInfo() {
	fmt.Printf("Name = %v Age = %v Score = %v\n", g.Name, g.Age, g.Score)
}

func (g *Graduate) SetScore(score int) {
	g.Score = score
}

func (g *Graduate) Testing() {
	fmt.Println("Graduate is testing...")
}

func main() {
	var pupil = &Pupil{
		Name: "Tom",
		Age:  10,
	}
	pupil.Testing()
	pupil.SetScore(90)
	pupil.ShowInfo()
	var graduate = &Graduate{
		Name: "Jerry",
		Age:  20,
	}
	graduate.Testing()
	graduate.SetScore(100)
	graduate.ShowInfo()
}
