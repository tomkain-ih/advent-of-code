package day21

import "testing"

const input1 = `029A
980A
179A
456A
379A`

func TestSolvePart1Example(t *testing.T) {
	solver := Solver{Input: input1}
	expected := "126384"
	actual := solver.SolvePart1()
	if expected != actual {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}

func TestSolvePart1(t *testing.T) {
	solver := Solver{File: "input.txt"}
	expected := "269"
	actual := solver.SolvePart1()
	if expected != actual {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}

func TestSolvePart2Example(t *testing.T) {
	solver := Solver{Input: input1}
	expected := "16"
	actual := solver.SolvePart2()
	if expected != actual {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}

func TestSolvePart2(t *testing.T) {
	solver := Solver{File: "input.txt"}
	expected := "758839075658876"
	actual := solver.SolvePart2()
	if expected != actual {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}
