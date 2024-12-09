package day09

import (
	"log"
	"os"
)

type Solver struct {
	Input string
	File  string
}

func (d Solver) GetLabel() string {
	return "Day 09"
}

func (d Solver) SolvePart1() int {
	input := d.GetInput()
	blocks, length := decodeBlocks(input)
	frag := fragBlocks(blocks, length)
	return checksum(frag)
}

func (d Solver) SolvePart2() int {
	//input := d.GetInput()
	return 0
}

func decodeBlocks(input string) (map[int]int, int) {
	blocks := make(map[int]int)
	spaces := make(map[int]int)
	block := 0
	for i, c := range input {
		v := int(c - '0')
		//v = block length
		if i%2 == 0 {
			//file block
			//i/2 = file id
			for j := 0; j < v; j++ {
				blocks[block] = i / 2
				block++
			}
		} else {
			//free space
			spaces[block] = v
			block += v
		}
	}
	//last block value = blocks length
	return blocks, block
}

func fragBlocks(blocks map[int]int, length int) map[int]int {
	frag := make(map[int]int)
	cursor := length
	for i := 0; i < cursor; i++ {
		if _, ok := blocks[i]; ok {
			frag[i] = blocks[i]
		} else {
			for {
				cursor--
				if _, ok := blocks[cursor]; ok {
					frag[i] = blocks[cursor]
					break
				}
			}
		}
	}
	return frag
}

func fragFiles(blocks map[int]int, length int, spaces map[int]int) map[int]int {
	frag := make(map[int]int)
	cursor := length
	for i := 0; i < cursor; i++ {
		if _, ok := blocks[i]; ok {
			frag[i] = blocks[i]
		} else {
			for {
				cursor--
				if _, ok := blocks[cursor]; ok {
					frag[i] = blocks[cursor]
					//delete(blocks, cursor)
					break
				}
			}
		}
	}
	return frag
}

func checksum(frag map[int]int) int {
	sum := 0
	for i, v := range frag {
		sum += i * v
	}
	return sum
}

func (d Solver) GetFile() string {
	if d.File != "" {
		return d.File
	}
	return "day09/input.txt"
}

func (d Solver) GetInput() string {
	if d.Input != "" {
		return d.Input
	}

	data, err := os.ReadFile(d.GetFile())
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}
