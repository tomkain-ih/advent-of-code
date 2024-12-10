package day10

import (
	"testing"
)

const input1 = `0123
1234
8765
9876`

// 1 trailhead, score of 1

const input2 = `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`

// 9 trailheads
// scores of 5, 6, 5, 3, 1, 3, 5, 3, and 5
// sum = 36

func TestSolvePart1Example(t *testing.T) {
	solver := Solver{Input: input2}
	expected := 36
	actual := solver.SolvePart1()
	if expected != actual {
		t.Errorf("Expected %d, but got %d", expected, actual)
	}
}

func TestSolvePart1(t *testing.T) {
	solver := Solver{File: "input.txt"}
	expected := 531
	actual := solver.SolvePart1()
	if expected != actual {
		t.Errorf("Expected %d, but got %d", expected, actual)
	}
}

func TestSolvePart2Example(t *testing.T) {
	solver := Solver{Input: input2}
	expected := 81
	actual := solver.SolvePart2()
	if expected != actual {
		t.Errorf("Expected %d, but got %d", expected, actual)
	}
}

func TestSolvePart2(t *testing.T) {
	solver := Solver{File: "input.txt"}
	expected := 1210
	actual := solver.SolvePart2()
	if expected != actual {
		t.Errorf("Expected %d, but got %d", expected, actual)
	}
}
