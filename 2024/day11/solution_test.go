package day11

import (
	"testing"
)

const input1 = `125 17`

// 6 blinks
const output6 = `2097446912 14168 4048 2 0 2 4 40 48 2024 40 48 80 96 2 8 6 7 6 0 3 2`

//const result6 = 22

// 25 blinks
const result25 = 55312

func TestBlink(t *testing.T) {
	expected := fromInput(output6)
	actual := blink(input1, 6)
	if len(expected) != len(actual) {
		t.Errorf("Expected %d, but got %d", len(expected), len(actual))
	}

	betterActual := betterBlink(input1, 6)
	sumBetterActual := sumValues(betterActual)
	if len(expected) != sumBetterActual {
		t.Errorf("Expected %d, but got %d", len(expected), sumBetterActual)
	}

}

func TestSolvePart1Example(t *testing.T) {
	solver := Solver{Input: input1}
	expected := "55312"
	actual := solver.SolvePart1()
	if expected != actual {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}

func TestSolvePart1(t *testing.T) {
	solver := Solver{File: "input.txt"}
	expected := "213625"
	actual := solver.SolvePart1()
	if expected != actual {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}

func TestSolvePart2(t *testing.T) {
	solver := Solver{File: "input.txt"}
	expected := "252442982856820"
	actual := solver.SolvePart2()
	if expected != actual {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}
