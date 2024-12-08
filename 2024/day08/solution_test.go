package day08

import (
	"testing"
)

const input1 = `............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............`

// top left is 0,0, bottom right is
// x,y
var antennas = map[rune][]Point{
	'0': {{4, -4}, {7, -3}, {5, -2}, {8, -1}},
	'A': {{9, -9}, {8, -8}, {6, -5}},
}

var expectedPairs = [][]Point{
	{{4, -4}, {7, -3}},
	{{4, -4}, {5, -2}},
	{{4, -4}, {8, -1}},
	{{7, -3}, {5, -2}},
	{{7, -3}, {8, -1}},
	{{5, -2}, {8, -1}},
	{{9, -9}, {8, -8}},
	{{9, -9}, {6, -5}},
	{{8, -8}, {6, -5}},
}

func TestReadGrid(t *testing.T) {
	antennas, last := readGrid(input1)
	if len(antennas) != 2 {
		t.Errorf("Expected 2 antennas, but got %d", len(antennas))
	}
	p := Point{11, -11}
	if last != p {
		t.Errorf("Expected last to be 11,-11, but got %v", last)
	}
	for k, v := range antennas {
		if len(v) != len(antennas[k]) {
			t.Errorf("Expected %d antennas for %v, but got %d", len(antennas[k]), k, len(v))
		}
	}
}

func TestGetPairs(t *testing.T) {
	pairs := getPairs(antennas)
	if len(pairs) != len(expectedPairs) {
		t.Errorf("Expected %d pairs, but got %d", len(expectedPairs), len(pairs))
	}
}

func TestGetAntiNodes(t *testing.T) {
	pair := []Point{{4, -4}, {7, -3}}
	antinodes := getAntiNodes(pair)
	if len(antinodes) != 2 {
		t.Errorf("Expected 2 antinodes, but got %d", len(antinodes))
	}
	expected := []Point{{10, -2}, {1, -5}}
	if antinodes[0] != expected[0] && antinodes[0] != expected[1] {
		t.Errorf("Expected %v, but got %v", expected, antinodes)
	}
	if antinodes[1] != expected[0] && antinodes[1] != expected[1] {
		t.Errorf("Expected %v, but got %v", expected, antinodes)
	}
}

func TestSolvePart1Example(t *testing.T) {
	solver := Solver{Input: input1}
	expected := "14"
	actual := solver.SolvePart1()
	if expected != actual {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}

func TestSolvePart1(t *testing.T) {
	solver := Solver{File: "input.txt"}
	expected := "344"
	actual := solver.SolvePart1()
	if expected != actual {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}

func TestSolvePart2Example(t *testing.T) {
	solver := Solver{Input: input1}
	expected := "34"
	actual := solver.SolvePart2()
	if expected != actual {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}

func TestSolvePart2(t *testing.T) {
	solver := Solver{File: "input.txt"}
	expected := "1182"
	actual := solver.SolvePart2()
	if expected != actual {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}
