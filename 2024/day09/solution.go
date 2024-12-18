package day09

import (
	"log"
	"os"
	"strconv"
)

type Solver struct {
	Input string
	File  string
}

type File struct {
	id, span, start int
}

func (d Solver) GetLabel() string {
	return "Day 09"
}

func (d Solver) SolvePart1() string {
	input := d.GetInput()
	blocks, length, _, _, _ := decodeBlocks(input)
	frag := fragBlocks(blocks, length)
	return strconv.Itoa(checksum(frag))
}

func (d Solver) SolvePart2() string {
	//TODO fix this
	input := d.GetInput()
	blocks, length, spaces, maxFileId, files := decodeBlocks(input)
	blocks = fragFiles(blocks, length, spaces, maxFileId, files)
	return strconv.Itoa(checksum(blocks))
}

func decodeBlocks(input string) (map[int]int, int, map[int]int, int, map[int]File) {
	//12345 -> 0..111....22222
	blocks := make(map[int]int) //block id -> file id
	spaces := make(map[int]int) //block id -> span/length
	files := make(map[int]File) //file id -> span/length and start
	file := 0
	block := 0
	for i, c := range input {
		v := int(c - '0') //convert rune to int
		//v = block length
		if i%2 == 0 {
			//file block
			file = i / 2
			files[file] = File{file, v, block}
			for j := 0; j < v; j++ {
				blocks[block] = file
				block++
			}
		} else {
			//free space
			if v > 0 {
				spaces[block] = v
			}
			block += v
		}
	}
	//last block value = blocks length
	//last file value = max file id
	return blocks, block, spaces, file, files
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

func fragFiles(blocks map[int]int, length int, spaces map[int]int, maxFileId int, files map[int]File) map[int]int {
	for f := maxFileId; f >= 0; f-- {
		file := files[f]
		for block := 0; block < length; block++ {
			if spaceSpan, ok := spaces[block]; ok {
				if spaceSpan >= file.span {
					//remove file from file.start + file.span in blocks
					for i := file.start; i < file.start+file.span; i++ {
						delete(blocks, i)
					}
					// add file to block + file.span in blocks
					for i := block; i < block+file.span; i++ {
						blocks[i] = f
					}
					// update spaces
					delete(spaces, block)
					if spaceSpan > file.span {
						spaces[block+file.span] = spaceSpan - file.span
					}
					break
				}
			}
		}
	}
	return blocks
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
