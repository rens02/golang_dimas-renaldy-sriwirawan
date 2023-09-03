package main

import (
	"fmt"
)

type Student struct {
	Name   string
	Scores []int
}

func (s *Student) CalculateAverageScore() float64 {
	// the process of calculating the average student score
	total := 0
	for _, score := range s.Scores {
		total += score
	}
	return float64(total) / float64(len(s.Scores))
}

func (s *Student) FindScore(f func(a, b int) bool) (score int, name string) {
	// the process to find the minimum or maximum score of students
	score = s.Scores[0]
	name = s.Name
	for _, sc := range s.Scores {
		if f(sc, score) {
			score = sc
			name = s.Name
		}
	}
	return score, name
}

func main() {
	// initialize array of Student with static data
	students := []Student{
		{Name: "Rizky", Scores: []int{80}},
		{Name: "Kobar", Scores: []int{75}},
		{Name: "Ismal", Scores: []int{70}},
		{Name: "Uman", Scores: []int{75}},
		{Name: "Yopan", Scores: []int{60}},
	}

	// find the average, minimum, and maximum values
	totalScore := 0
	minScore := students[0].Scores[0]
	maxScore := students[0].Scores[0]
	var minName, maxName string
	for _, s := range students {
		totalScore += s.Scores[0]
		if min, name := s.FindScore(func(a, b int) bool { return a < b }); min < minScore {
			minScore = min
			minName = name
		}
		if max, name := s.FindScore(func(a, b int) bool { return a > b }); max > maxScore {
			maxScore = max
			maxName = name
		}
	}

	averageScore := float64(totalScore) / float64(len(students))
	fmt.Printf("Average Score : %.2f\n", averageScore)
	fmt.Printf("Min Score of students : %s(%d)\n", minName, minScore)
	fmt.Printf("Max Score of students : %s(%d)\n", maxName, maxScore)
}
