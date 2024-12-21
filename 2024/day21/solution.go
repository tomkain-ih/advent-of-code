package day21

import (
	"log"
	"os"
)

type Solver struct {
	Input string
	File  string
}

func (d Solver) GetLabel() string {
	return "Day 21"
}

func (d Solver) GetFile() string {
	if d.File != "" {
		return d.File
	}
	return "day21/input.txt"
}

func (d Solver) SolvePart1() string {
	return "Part 1 Solve"
}

func (d Solver) SolvePart2() string {
	return "Part 2 Solve"
}

func numericToDirectional(code string) string {
	numericKeypad := map[string]point{
		"0": {1, 0},
		"A": {2, 0},
		"1": {0, 1},
		"2": {1, 1},
		"3": {2, 1},
		"4": {0, 2},
		"5": {1, 2},
		"6": {2, 2},
		"7": {0, 3},
		"8": {1, 3},
		"9": {2, 3},
	}
	var keyPresses []point
	for _, c := range code {
		keyPresses = append(keyPresses, numericKeypad[string(c)])
	}

	directions := ""
	for i := 0; i < len(keyPresses)-1; i++ {
		directions += traverse(keyPresses[i], keyPresses[i+1])
	}
}

func traverse(start, end point) string {
	
}

type point struct {
	x, y int
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
