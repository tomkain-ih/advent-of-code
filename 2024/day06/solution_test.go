package day06

import (
	"fmt"
	"testing"
)

const input1 = `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

const part1Grid = `....#.....
....XXXXX#
....X...X.
..#.X...X.
..XXXXX#X.
..X.X.X.X.
.#XXXXXXX.
.XXXXXXX#.
#XXXXXXX..
......#X..`

func TestReadGrid(t *testing.T) {
	solve, _ := readGrid(part1Grid)
	for k, v := range solve {
		if v == 'X' {
			fmt.Println(k)
		}
	}
}

func TestSolvePart1Example(t *testing.T) {
	solver := Solver{Input: input1}
	expected := "41"
	actual := solver.SolvePart1()
	if expected != actual {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}

func TestSolvePart1(t *testing.T) {
	solver := Solver{File: "input.txt"}
	expected := "5444"
	actual := solver.SolvePart1()
	if expected != actual {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}

func TestSolvePart2Example(t *testing.T) {
	solver := Solver{Input: input1}
	expected := "6"
	actual := solver.SolvePart2()
	if expected != actual {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}

func TestSolvePart2(t *testing.T) {
	solver := Solver{File: "input.txt"}
	expected := "1946"
	actual := solver.SolvePart2()
	if expected != actual {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}
