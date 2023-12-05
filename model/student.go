package model

type student struct {
	Name  string
	score float64
}

// factory pattern
func NewStudent(n string, s float64) *student {
	return &student{
		Name:  n,
		score: s,
	}
}

func (s *student) GetScore() float64 {
	return s.score
}
