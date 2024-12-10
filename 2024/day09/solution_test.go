package day09

import (
	"testing"
)

const input1 = "12345"

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
var files1 = map[int]File{
	0: {0, 1, 0},
	1: {1, 3, 3},
	2: {2, 5, 10},
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

const input2 = `2333133121414131402`

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
var files2 = map[int]File{
	0: {0, 2, 0},
	1: {1, 3, 5},
	2: {2, 1, 11},
	3: {3, 3, 15},
	4: {4, 2, 19},
	5: {5, 4, 22},
	6: {6, 4, 27},
	7: {7, 3, 32},
	8: {8, 4, 36},
	9: {9, 2, 40},
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

// 00992111777.44.333....5555.6666.....8888..
var fileFrag2 = map[int]int{
	0:  0,
	1:  0,
	2:  9,
	3:  9,
	4:  2,
	5:  1,
	6:  1,
	7:  1,
	8:  7,
	9:  7,
	10: 7,
	12: 4,
	13: 4,
	15: 3,
	16: 3,
	17: 3,
	22: 5,
	23: 5,
	24: 5,
	25: 5,
	27: 6,
	28: 6,
	29: 6,
	30: 6,
	36: 8,
	37: 8,
	38: 8,
	39: 8,
}

func TestDecodeBlocks(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		length int
		blocks map[int]int
		id     int
		files  map[int]File
	}{
		{"1", input1, length1, blocks1, 2, files1},
		{"2", input2, length2, blocks2, 9, files2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, l, s, id, f := decodeBlocks(tt.input)
			if tt.length != l {
				t.Errorf("Expected %d, but got %d", tt.length, l)
			}
			for k, v := range tt.blocks {
				if b[k] != v {
					t.Errorf("Expected %d, but got %d", v, b[k])
				}
			}
			for k, _ := range s {
				if _, ok := tt.blocks[k]; ok {
					t.Errorf("%d key occupies space and file", k)
				}
			}
			if tt.id != id {
				t.Errorf("Expected %d, but got %d", tt.id, id)
			}
			for k, v := range f {
				if tt.files[k] != v {
					t.Errorf("Expected %v, but got %v", tt.files[k], v)
				}
			}
		})
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

func TestFragFiles(t *testing.T) {
	blocks, length, spaces, maxFileId, files := decodeBlocks(input2)
	f := fragFiles(blocks, length, spaces, maxFileId, files)
	for k, v := range fileFrag2 {
		if f[k] != v {
			t.Errorf("Expected %d, but got %d for key %d", v, f[k], k)
		}
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
	expected := 6326952672104 //failing, too high
	actual := solver.SolvePart2()
	if expected != actual {
		t.Errorf("Expected %d, but got %d", expected, actual)
	}
}
