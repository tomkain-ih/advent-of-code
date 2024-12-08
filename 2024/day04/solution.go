package day04

import (
	"log"
	"os"
	"strconv"
	"strings"
)

type Solver struct {
	Input string
	File  string
}

func (d Solver) GetLabel() string {
	return "Day 04"
}

func (d Solver) SolvePart1() int {
	input := d.GetInput()
	grid, rows, cols := makeGrid(input)
	count := 0
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			key := key(row, col)
			if grid[key] == "X" {
				for _, adjKey := range adjacentKeys(key) {
					if grid[adjKey] == "M" {
						adjKey2 := nextKey(key, adjKey)
						if grid[adjKey2] == "A" {
							lastKey := nextKey(adjKey, adjKey2)
							if grid[lastKey] == "S" {
								count++
							}
						}
					}
				}
			}
		}
	}
	return count
}

func nextKey(first string, second string) string {
	r1, c1 := getCoords(first)
	r2, c2 := getCoords(second)
	row := r2
	col := c2
	if r1 < r2 {
		// move up
		row++
	} else if r1 > r2 {
		// move down
		row--
	}
	if c1 < c2 {
		// move right
		col++
	} else if c1 > c2 {
		// move left
		col--
	}
	return key(row, col)
}

func adjacentKeys(init string) []string {
	row, col := getCoords(init)
	return []string{
		key(row-1, col-1),
		key(row-1, col),
		key(row-1, col+1),
		key(row, col-1),
		key(row, col+1),
		key(row+1, col-1),
		key(row+1, col),
		key(row+1, col+1),
	}
}

func diagonalKeys(init string) []string {
	row, col := getCoords(init)
	return []string{
		key(row-1, col-1),
		key(row-1, col+1),
		key(row+1, col-1),
		key(row+1, col+1),
	}
}

func getCoords(key string) (int, int) {
	parts := strings.Split(key, "-")
	row, err := strconv.Atoi(parts[0])
	if err != nil {
		log.Fatal(err)
	}
	col, err := strconv.Atoi(parts[1])
	if err != nil {
		log.Fatal(err)
	}
	return row, col
}

func key(row int, col int) string {
	return strconv.Itoa(row) + "-" + strconv.Itoa(col)
}

func makeGrid(input string) (map[string]string, int, int) {
	grid := make(map[string]string)
	row := 0
	col := 0
	for _, char := range input {
		if char == '\n' {
			row++
			col = 0
		} else {
			grid[key(row, col)] = string(char)
			col++
		}
	}
	return grid, row + 1, col
}

func (d Solver) SolvePart2() int {
	input := d.GetInput()
	grid, rows, cols := makeGrid(input)
	count := 0
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			key := key(row, col)
			if grid[key] == "A" {
				if twoMsAndSs(grid, diagonalKeys(key)) {
					count++
				}
			}
		}
	}
	return count
}

func twoMsAndSs(grid map[string]string, keys []string) bool {
	var m []string
	var s []string
	for _, key := range keys {
		if grid[key] == "M" {
			m = append(m, key)
		} else if grid[key] == "S" {
			s = append(s, key)
		}
	}
	if len(m) == 2 && len(s) == 2 {
		r1, c1 := getCoords(m[0])
		r2, c2 := getCoords(m[1])
		if r1 == r2 || c1 == c2 {
			return true
		}
	}
	return false
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

func (d Solver) GetFile() string {
	if d.File != "" {
		return d.File
	}
	return "day04/input.txt"
}
