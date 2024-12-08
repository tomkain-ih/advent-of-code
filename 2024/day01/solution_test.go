package day01

import (
	"testing"
)

const input1 = `3   4
4   3
2   5
1   3
3   9
3   3`

func TestSolvePart1Example(t *testing.T) {
	solver := Solver{Input: input1}
	expected := 11
	actual := solver.SolvePart1()
	if expected != actual {
		t.Errorf("Expected %d, but got %d", expected, actual)
	}
}

func TestSolvePart1(t *testing.T) {
	solver := Solver{File: "input.txt"}
	expected := 2742123
	actual := solver.SolvePart1()
	if expected != actual {
		t.Errorf("Expected %d, but got %d", expected, actual)
	}
}

func TestSolvePart2Example(t *testing.T) {
	solver := Solver{Input: input1}
	expected := 31
	actual := solver.SolvePart2()
	if expected != actual {
		t.Errorf("Expected %d, but got %d", expected, actual)
	}
}

func TestSolvePart2(t *testing.T) {
	solver := Solver{File: "input.txt"}
	expected := 21328497
	actual := solver.SolvePart2()
	if expected != actual {
		t.Errorf("Expected %d, but got %d", expected, actual)
	}
}
