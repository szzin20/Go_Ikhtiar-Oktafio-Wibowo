package main

import (
	"fmt"
)

type Student struct {
	name  []string
	score []int
}

func (s Student) Average() float64 {
	totalScore := 0
	for _, score := range s.score {
		totalScore += score
	}
	return float64(totalScore) / float64(len(s.score))
}

func (s Student) Min() (min int, name string) {
	min = s.score[0]
	name = s.name[0]
	for i := 1; i < len(s.score); i++ {
		if s.score[i] < min {
			min = s.score[i]
			name = s.name[i]
		}
	}
	return min, name
}

func (s Student) Max() (max int, name string) {
	max = s.score[0]
	name = s.name[0]
	for i := 1; i < len(s.score); i++ {
		if s.score[i] > max {
			max = s.score[i]
			name = s.name[i]
		}
	}
	return max, name
}

func main() {
	var a Student

	names := []string{"tiar", "okta", "queen", "zachra", "dora"}
	scores := []int{60, 70, 75, 75, 80}

	for i := 0; i < 5; i++ {
		a.name = append(a.name, names[i])
		a.score = append(a.score, scores[i])
	}

	fmt.Println("Student Data:")
	for i, name := range a.name {
		fmt.Printf("%s: %d\n", name, a.score[i])
	}

	fmt.Println("\nAverage Score Students is", a.Average())

	scoreMax, nameMax := a.Max()
	fmt.Println("Max Score Students is:", nameMax, "(", scoreMax, ")")

	scoreMin, nameMin := a.Min()
	fmt.Println("Min Score Students is:", nameMin, "(", scoreMin, ")")
}