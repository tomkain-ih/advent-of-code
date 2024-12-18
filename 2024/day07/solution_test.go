package day07

import (
	"fmt"
	"testing"
)

const input1 = `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`

func TestGenerateValues(t *testing.T) {
	expected := []int{29, 190}
	actual := generate([]int{10}, 19, false)
	if expected[0] != actual[0] || expected[1] != actual[1] {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}

func TestStream(t *testing.T) {
	input := []int{81, 40, 27}
	left := []int{input[0]}
	for i := 1; i < len(input); i++ {
		left = generate(left, input[i], false)
	}
	fmt.Println(left)
}

func TestSolvePart1Example(t *testing.T) {
	solver := Solver{Input: input1}
	expected := "3749"
	actual := solver.SolvePart1()
	if expected != actual {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}

func TestSolvePart1(t *testing.T) {
	solver := Solver{File: "input.txt"}
	expected := "6392012777720"
	actual := solver.SolvePart1()
	if expected != actual {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}

func TestSolvePart2Example(t *testing.T) {
	solver := Solver{Input: input1}
	expected := "11387"
	actual := solver.SolvePart2()
	if expected != actual {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}

func TestSolvePart2(t *testing.T) {
	solver := Solver{File: "input.txt"}
	expected := "61561126043536"
	actual := solver.SolvePart2()
	if expected != actual {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}
