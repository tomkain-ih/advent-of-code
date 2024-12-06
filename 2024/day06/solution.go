package day06

import (
	"log"
	"os"
)

type Solver struct {
	Input string
	File  string
}

func (d Solver) GetLabel() string {
	return "Day 06"
}

func (d Solver) SolvePart1() string {
	//input := d.GetInput()
	return "Part 1 solve"
}

func (d Solver) SolvePart2() string {
	//input := d.GetInput()
	return "Part 2 solve"
}

func (d Solver) GetFile() string {
	if d.File != "" {
		return d.File
	}
	return "day06/input.txt"
}

func (d Solver) GetInput() string {
	if d.Input != "" {
		return d.Input
	}

	data, err := os.ReadFile(d.GetFile())
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}
