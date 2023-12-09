package main

import (
	"fmt"
)

type Student struct {
	Name  string
	Age   int
	Score int
}

func (s *Student) ShowInfo() {
	fmt.Printf("Name = %v Age = %v Score = %v\n", s.Name, s.Age, s.Score)
}

func (s *Student) SetScore(score int) {
	s.Score = score
}

func (s *Student) getSum(n1 int, n2 int) int {
	return n1 + n2
}

type Pupil struct {
	Student
}

func (p *Pupil) Testing() {
	fmt.Println("PrimarySchoolStydent is testing...")
}

type Graduate struct {
	Student
}

func (g *Graduate) Testing() {
	fmt.Println("Graduate is testing...")
}

func main() {
	pupil := &Pupil{}
	pupil.Student.Name = "tom"
	pupil.Student.Age = 8
	pupil.Testing()
	pupil.Student.SetScore(80)
	pupil.Student.ShowInfo()
	fmt.Println("res =", pupil.Student.getSum(1, 2))
	graduate := &Graduate{}
	graduate.Student.Name = "mary"
	graduate.Student.Age = 28
	graduate.Testing()
	graduate.Student.SetScore(90)
	graduate.Student.ShowInfo()
	fmt.Println("res =", graduate.Student.getSum(1, 2))
}
