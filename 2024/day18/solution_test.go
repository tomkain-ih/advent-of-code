package day18

import (
	"testing"
)

const input1 = `5,4
4,2
4,5
3,0
2,1
6,3
2,4
1,5
0,6
3,3
2,6
5,1
1,2
5,5
2,5
6,5
1,4
0,4
6,4
1,1
6,1
1,0
0,5
1,6
2,0`

func TestSolvePart1Example(t *testing.T) {
	solver := Solver{Input: input1, MaxIndex: 6, Bytes: 12}
	expected := 22
	actual := solver.SolvePart1()
	if expected != actual {
		t.Errorf("Expected %d, but got %d", expected, actual)
	}
}

func TestSolvePart1(t *testing.T) {
	solver := Solver{File: "input.txt", MaxIndex: 70, Bytes: 1024}
	expected := 446
	actual := solver.SolvePart1()
	if expected != actual {
		t.Errorf("Expected %d, but got %d", expected, actual)
	}
}

func TestSolvePart2Example(t *testing.T) {
	solver := Solver{Input: input1}
	expected := 55312
	actual := solver.SolvePart2()
	if expected != actual {
		t.Errorf("Expected %d, but got %d", expected, actual)
	}
}

func TestSolvePart2(t *testing.T) {
	solver := Solver{File: "input.txt"}
	expected := 0 // (39, 40)
	actual := solver.SolvePart2()
	if expected != actual {
		t.Errorf("Expected %d, but got %d", expected, actual)
	}
}
