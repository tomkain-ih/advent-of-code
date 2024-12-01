package day01

import (
	"log"
	"os"
)

type Day01Solver struct{}

func (d Day01Solver) GetLabel() string {
	return "Day 01"
}

func (d Day01Solver) Solve() string {
	input := d.GetInput()
	// Implement your solution logic here using the input
	return "Solution for Day 01: " + input
}

func (d Day01Solver) GetInput() string {
	data, err := os.ReadFile("day01/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}
