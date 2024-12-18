package day10

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

type point struct {
	x, y int
}

func (d Solver) GetLabel() string {
	return "Day 10"
}

func (d Solver) SolvePart1() string {
	input := d.GetInput()
	grid, trailheads := parseInput(input)
	sum := 0
	for _, trailhead := range trailheads {
		score := countDestinations(grid, trailhead)
		sum += score
	}
	return strconv.Itoa(sum)
}

func (d Solver) SolvePart2() string {
	input := d.GetInput()
	grid, trailheads := parseInput(input)
	sum := 0
	for _, trailhead := range trailheads {
		score := countPaths(grid, trailhead)
		sum += score
	}
	return strconv.Itoa(sum)
}

func parseInput(input string) (map[point]int, []point) {
	grid := make(map[point]int)
	var trailheads []point
	for y, line := range strings.Split(input, "\n") {
		for x, char := range line {
			height := int(char - '0')
			p := point{x, y}
			grid[p] = height
			if height == 0 {
				trailheads = append(trailheads, p)
			}
		}
	}
	return grid, trailheads
}

func countDestinations(grid map[point]int, trailhead point) int {
	height := 0
	neighbors := trailhead.neighbors()
	seen, _ := tryNeighbors(grid, neighbors, height, make(map[point]struct{}))
	return len(seen)
}

func countPaths(grid map[point]int, trailhead point) int {
	height := 0
	neighbors := trailhead.neighbors()
	_, score := tryNeighbors(grid, neighbors, height, make(map[point]struct{}))
	return score
}

func tryNeighbors(grid map[point]int, neighbors []point, height int, seen map[point]struct{}) (map[point]struct{}, int) {
	score := 0
	var s int
	for _, neighbor := range neighbors {
		if val, ok := grid[neighbor]; ok {
			if val == height+1 {
				if val == 9 {
					seen[neighbor] = struct{}{}
					score++
				} else {
					seen, s = tryNeighbors(grid, neighbor.neighbors(), height+1, seen)
					score += s
				}
			}
		}
	}
	return seen, score
}

func (p point) neighbors() []point {
	return []point{
		{p.x - 1, p.y},
		{p.x + 1, p.y},
		{p.x, p.y - 1},
		{p.x, p.y + 1},
	}
}

func (d Solver) GetFile() string {
	if d.File != "" {
		return d.File
	}
	return "day10/input.txt"
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
