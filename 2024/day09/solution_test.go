package day09

import (
	"testing"
)

const input1 = "12345"
const input2 = `2333133121414131402`

// 0..111....22222
var length1 = 15
var blocks1 = map[int]int{
	0:  0,
	3:  1,
	4:  1,
	5:  1,
	10: 2,
	11: 2,
	12: 2,
	13: 2,
	14: 2,
}

// 022111222......
var frag1 = map[int]int{
	0: 0,
	1: 2,
	2: 2,
	3: 1,
	4: 1,
	5: 1,
	6: 2,
	7: 2,
	8: 2,
}

// 00...111...2...333.44.5555.6666.777.888899
var length2 = 42
var blocks2 = map[int]int{
	0:  0,
	1:  0,
	5:  1,
	6:  1,
	7:  1,
	11: 2,
	15: 3,
	16: 3,
	17: 3,
	19: 4,
	20: 4,
	22: 5,
	23: 5,
	24: 5,
	25: 5,
	27: 6,
	28: 6,
	29: 6,
	30: 6,
	32: 7,
	33: 7,
	34: 7,
	36: 8,
	37: 8,
	38: 8,
	39: 8,
	40: 9,
	41: 9,
}

// 0099811188827773336446555566..............
var frag2 = map[int]int{
	0:  0,
	1:  0,
	2:  9,
	3:  9,
	4:  8,
	5:  1,
	6:  1,
	7:  1,
	8:  8,
	9:  8,
	10: 8,
	11: 2,
	12: 7,
	13: 7,
	14: 7,
	15: 3,
	16: 3,
	17: 3,
	18: 6,
	19: 4,
	20: 4,
	21: 6,
	22: 5,
	23: 5,
	24: 5,
	25: 5,
	26: 6,
	27: 6,
}

func TestDecodeBlocks(t *testing.T) {
	b, l := decodeBlocks(input1)
	if length1 != l {
		t.Errorf("Expected %d, but got %d", length1, l)
	}
	for k, v := range blocks1 {
		if b[k] != v {
			t.Errorf("Expected %d, but got %d", v, b[k])
		}
	}

	b, l = decodeBlocks(input2)
	if length2 != l {
		t.Errorf("Expected %d, but got %d", length2, l)
	}
	for k, v := range blocks2 {
		if b[k] != v {
			t.Errorf("Expected %d, but got %d", v, b[k])
		}
	}
}

func TestFragBlocks(t *testing.T) {
	f := fragBlocks(blocks1, length1)
	for k, v := range frag1 {
		if f[k] != v {
			t.Errorf("Expected %d, but got %d", v, f[k])
		}
	}

	f = fragBlocks(blocks2, length2)
	for k, v := range f {
		if frag2[k] != v {
			t.Errorf("Expected %d, but got %d for key %d", v, frag2[k], k)
		}
	}
}

func TestChecksum(t *testing.T) {
	c := checksum(frag2)
	expected := 1928
	if c != expected {
		t.Errorf("Expected %d, but got %d", expected, c)
	}
}

func TestSolvePart1Example(t *testing.T) {
	solver := Solver{Input: input2}
	expected := 1928
	actual := solver.SolvePart1()
	if expected != actual {
		t.Errorf("Expected %d, but got %d", expected, actual)
	}
}

func TestSolvePart1(t *testing.T) {
	solver := Solver{File: "input.txt"}
	expected := 6299243228569
	actual := solver.SolvePart1()
	if expected != actual {
		t.Errorf("Expected %d, but got %d", expected, actual)
	}
}

func TestSolvePart2Example(t *testing.T) {
	solver := Solver{Input: input2}
	expected := 2858
	actual := solver.SolvePart2()
	if expected != actual {
		t.Errorf("Expected %d, but got %d", expected, actual)
	}
}

func TestSolvePart2(t *testing.T) {
	solver := Solver{File: "input.txt"}
	expected := 1182
	actual := solver.SolvePart2()
	if expected != actual {
		t.Errorf("Expected %d, but got %d", expected, actual)
	}
}
