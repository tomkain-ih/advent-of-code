package day05

import (
	"testing"
)

const input1 = `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`

func TestCreateRulesMap(t *testing.T) {
	input := `53|29
53|13`
	expected := map[int][]int{
		53: {29, 13},
	}
	actual := createRulesMap(input)
	if expected[53][0] != actual[53][0] || expected[53][1] != actual[53][1] {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}

func slicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestCreateUpdatesSlice(t *testing.T) {
	updates := createUpdatesSlice(input1)
	expected1 := []int{75, 47, 61, 53, 29}
	expected2 := []int{97, 61, 53, 29, 13}

	if !slicesEqual(updates[0], expected1) {
		t.Errorf("Expected %v, but got %v", expected1, updates[0])
	}
	if !slicesEqual(updates[1], expected2) {
		t.Errorf("Expected %v, but got %v", expected2, updates[1])
	}
}

func TestVerifySequence(t *testing.T) {
	updates1 := []int{75, 47, 61, 53, 29}
	rules := createRulesMap(input1)
	expected := true
	actual := verifySequence(updates1, rules)
	if expected != actual {
		t.Errorf("Expected %t, but got %t", expected, actual)
	}

	updates2 := []int{75, 97, 47, 61, 53}
	expected = false
	actual = verifySequence(updates2, rules)
	if expected != actual {
		t.Errorf("Expected %t, but got %t", expected, actual)
	}
}

func TestMiddleElem(t *testing.T) {
	updates := []int{75, 47, 61, 53, 29}
	expected := 61
	actual := middleElem(updates)
	if expected != actual {
		t.Errorf("Expected %d, but got %d", expected, actual)
	}

	updates = []int{61, 13, 29}
	expected = 13
	actual = middleElem(updates)
	if expected != actual {
		t.Errorf("Expected %d, but got %d", expected, actual)
	}
}

func TestSolvePart1Example(t *testing.T) {
	solver := Solver{Input: input1}
	expected := 143
	actual := solver.SolvePart1()
	if expected != actual {
		t.Errorf("Expected %d, but got %d", expected, actual)
	}
}

func TestSolvePart1(t *testing.T) {
	solver := Solver{File: "input.txt"}
	expected := 5639
	actual := solver.SolvePart1()
	if expected != actual {
		t.Errorf("Expected %d, but got %d", expected, actual)
	}
}

func TestCorrectSequence(t *testing.T) {
	rules := createRulesMap(input1)

	update := []int{97, 13, 75, 29, 47}
	expected := []int{97, 75, 47, 29, 13}
	actual := correctSequence(update, rules)
	if !slicesEqual(expected, actual) {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
	update = []int{61, 13, 29}
	expected = []int{61, 29, 13}
	actual = correctSequence(update, rules)
	if !slicesEqual(expected, actual) {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}

func TestSolvePart2Example(t *testing.T) {
	solver := Solver{Input: input1}
	expected := 123
	actual := solver.SolvePart2()
	if expected != actual {
		t.Errorf("Expected %d, but got %d", expected, actual)
	}
}

func TestSolvePart2(t *testing.T) {
	solver := Solver{File: "input.txt"}
	expected := 5273
	actual := solver.SolvePart2()
	if expected != actual {
		t.Errorf("Expected %d, but got %d", expected, actual)
	}
}
