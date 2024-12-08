package day03

import (
	"testing"
)

const input1 = `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`
const input2 = `xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`

func TestRegexPatterns(t *testing.T) {
	matches := getMatches(input1, getPart1RegexPattern())
	if len(matches) != 4 {
		t.Errorf("Expected 4 matches, but got %d", len(matches))
	}

	matches = getMatches(input2, getPart2RegexPattern())
	if len(matches) != 6 {
		t.Errorf("Expected 6 matches, but got %d", len(matches))
	}
}

func TestMultiply(t *testing.T) {
	expected := 8
	actual := multiply("mul(2,4)")
	if expected != actual {
		t.Errorf("Expected %d, but got %d", expected, actual)
	}
}

func TestSolvePart1Example(t *testing.T) {
	solver := Solver{Input: input1}
	expected := 161
	actual := solver.SolvePart1()
	if expected != actual {
		t.Errorf("Expected %d, but got %d", expected, actual)
	}
}

func TestSolvePart1(t *testing.T) {
	solver := Solver{File: "input.txt"}
	expected := 174960292
	actual := solver.SolvePart1()
	if expected != actual {
		t.Errorf("Expected %d, but got %d", expected, actual)
	}
}

func TestSolvePart2Example(t *testing.T) {
	solver := Solver{Input: input2}
	expected := 48
	actual := solver.SolvePart2()
	if expected != actual {
		t.Errorf("Expected %d, but got %d", expected, actual)
	}
}

func TestSolvePart2(t *testing.T) {
	solver := Solver{File: "input.txt"}
	expected := 56275602
	actual := solver.SolvePart2()
	if expected != actual {
		t.Errorf("Expected %d, but got %d", expected, actual)
	}
}
