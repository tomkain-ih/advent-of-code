package day04

import (
	"testing"
)

const input1 = `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

func TestMakeGrid(t *testing.T) {
	grid, rows, cols := makeGrid(input1)
	if grid["0-0"] != "M" {
		t.Errorf("Expected M, but got %s", grid["0-0"])
	}
	if grid["4-0"] != "X" {
		t.Errorf("Expected X, but got %s", grid["4-0"])
	}
	if grid["6-9"] != "S" {
		t.Errorf("Expected S, but got %s", grid["6-9"])
	}
	if rows != 10 {
		t.Errorf("Expected 10, but got %d", rows)
	}
	if cols != 10 {
		t.Errorf("Expected 10, but got %d", cols)
	}
}

func TestNextKey(t *testing.T) {
	if nextKey("0-0", "1-1") != "2-2" {
		t.Errorf("Expected 2-2, but got %s", nextKey("0-0", "1-1"))
	}
	if nextKey("0-0", "1-0") != "2-0" {
		t.Errorf("Expected 2-0, but got %s", nextKey("0-0", "1-0"))
	}
	if nextKey("0-0", "0-1") != "0-2" {
		t.Errorf("Expected 0-2, but got %s", nextKey("0-0", "0-1"))
	}
}

func TestSolvePart1Example(t *testing.T) {
	solver := Solver{Input: input1}
	expected := 18
	actual := solver.SolvePart1()
	if expected != actual {
		t.Errorf("Expected %d, but got %d", expected, actual)
	}
}

func TestSolvePart1(t *testing.T) {
	solver := Solver{File: "input.txt"}
	expected := 2578
	actual := solver.SolvePart1()
	if expected != actual {
		t.Errorf("Expected %d, but got %d", expected, actual)
	}
}

func TestSolvePart2Example(t *testing.T) {
	solver := Solver{Input: input1}
	expected := 9
	actual := solver.SolvePart2()
	if expected != actual {
		t.Errorf("Expected %d, but got %d", expected, actual)
	}
}

func TestSolvePart2(t *testing.T) {
	solver := Solver{File: "input.txt"}
	expected := 1972
	actual := solver.SolvePart2()
	if expected != actual {
		t.Errorf("Expected %d, but got %d", expected, actual)
	}
}
