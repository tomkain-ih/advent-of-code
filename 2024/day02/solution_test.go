package day02

import (
	"testing"
)

const input1 = `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

func TestIsSafe(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Example 1", args{"7 6 4 2 1"}, true},
		{"Example 2", args{"1 2 7 8 9"}, false},
		{"Example 3", args{"9 7 6 2 1"}, false},
		{"Example 4", args{"1 3 2 4 5"}, false},
		{"Example 5", args{"8 6 4 4 1"}, false},
		{"Example 6", args{"1 3 6 7 9"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSafe(tt.args.line); got != tt.want {
				t.Errorf("isSafe() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolvePart1Example(t *testing.T) {
	solver := Solver{Input: input1}
	expected := "2"
	actual := solver.SolvePart1()
	if expected != actual {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}

func TestSolvePart1(t *testing.T) {
	solver := Solver{File: "input.txt"}
	expected := "686"
	actual := solver.SolvePart1()
	if expected != actual {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}

func TestIsSafeWithTolerance(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Example 1", args{"7 6 4 2 1"}, true},
		{"Example 2", args{"1 2 7 8 9"}, false},
		{"Example 3", args{"9 7 6 2 1"}, false},
		{"Example 4", args{"1 3 2 4 5"}, true},
		{"Example 5", args{"8 6 4 4 1"}, true},
		{"Example 6", args{"1 3 6 7 9"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSafeWithTolerance(tt.args.line); got != tt.want {
				t.Errorf("isSafe() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolvePart2Example(t *testing.T) {
	solver := Solver{Input: input1}
	expected := "4"
	actual := solver.SolvePart2()
	if expected != actual {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}

func TestSolvePart2(t *testing.T) {
	solver := Solver{File: "input.txt"}
	expected := "717"
	actual := solver.SolvePart2()
	if expected != actual {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}
