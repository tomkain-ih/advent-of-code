package day06

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

type Vector struct {
	point     complex128
	direction complex128
}

func (d Solver) GetLabel() string {
	return "Day 06"
}

func (d Solver) SolvePart1() string {
	input := d.GetInput()
	grid, start := readGrid(input)
	path, _ := walk(grid, start)
	return strconv.Itoa(len(path))
}

func (d Solver) SolvePart2() string {
	input := d.GetInput()
	grid, start := readGrid(input)
	path, _ := walk(grid, start)
	points := 0
	for _, p := range path {
		grid[p] = '#'
		_, looped := walk(grid, start)
		if looped {
			points++
		}
		grid[p] = '.'
	}
	return strconv.Itoa(points)
}

func readGrid(input string) (map[complex128]rune, complex128) {
	//grid as map: key = x/y point, value = char
	grid := make(map[complex128]rune)
	//starting point '^'
	var start complex128
	//iterate over split by \n
	for y, line := range strings.Split(input, "\n") {
		//iterate over chars
		for x, char := range line {
			p := complex(float64(x), float64(y))
			grid[p] = char
			if char == '^' {
				start = p
			}
		}
	}
	return grid, start
}

func walk(grid map[complex128]rune, start complex128) ([]complex128, bool) {
	vectors := make(map[Vector]struct{})
	seen := make(map[complex128]struct{})
	pos := start
	dir := complex(0, -1) // up
	looped := false
	for {
		v := Vector{pos, dir}
		if _, ok := vectors[v]; ok {
			//if already visited, stop
			looped = true
			break
		}
		if _, ok := grid[pos]; !ok {
			//if off grid, stop
			break
		}
		for {
			next := pos + dir
			if val, _ := grid[next]; val != '#' {
				break
			}
			// if hit wall, turn clockwise
			dir *= complex(0, 1)
		}
		seen[pos] = struct{}{}
		vectors[v] = struct{}{}
		pos += dir
	}
	path := make([]complex128, 0, len(seen))
	for p := range seen {
		path = append(path, p)
	}
	return path, looped
}

func (d Solver) GetFile() string {
	if d.File != "" {
		return d.File
	}
	return "day06/input.txt"
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
