package day22

import "testing"

const input1 = `1
10
100
2024`

func Test_Mix(t *testing.T) {
	expected := 37
	actual := mix(42, 15)
	if expected != actual {
		t.Errorf("Expected %d, but got %d", expected, actual)
	}
}

func Test_Prune(t *testing.T) {
	expected := 16113920
	actual := prune(100000000)
	if expected != actual {
		t.Errorf("Expected %d, but got %d", expected, actual)
	}
}

func Test_Evolve(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{123, 15887950},
		{704524, 1553684},
		{7753432, 5908254},
	}
	for _, tt := range tests {
		actual := evolve(tt.input)
		if tt.expected != actual {
			t.Errorf("Expected %d, but got %d", tt.expected, actual)
		}
	}
}

func Test_ComputeSecretNumber(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{1, 8685429},
		{10, 4700978},
		{100, 15273692},
		{2024, 8667524},
	}
	for _, tt := range tests {
		actual := computeSecretNumber(tt.input, 2000)
		if tt.expected != actual {
			t.Errorf("Expected %d, but got %d", tt.expected, actual)
		}
	}

}

func TestSolvePart1Example(t *testing.T) {
	solver := Solver{Input: input1}
	expected := "37327623"
	actual := solver.SolvePart1()
	if expected != actual {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}

func TestSolvePart1(t *testing.T) {
	solver := Solver{File: "input.txt"}
	expected := "18525593556"
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
