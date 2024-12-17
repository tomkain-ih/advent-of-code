package day12

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Solver struct {
	Input string
	File  string
}

func (d Solver) GetLabel() string {
	return "Day 12"
}

func (d Solver) GetFile() string {
	if d.File != "" {
		return d.File
	}
	return "day12/input.txt"
}

func (d Solver) SolvePart1() int {
	input := d.GetInput()
	grid := readGrid(input)
	regions := findRegions(grid)

	// for each region, compute area, perimeter, and cost
	cost := 0
	for _, region := range regions {
		// for each region, multiply area by perimeter and sum products
		area := len(region)
		plant := grid[region[0]]
		perimeter := computePerimeter(region, plant, grid)
		fmt.Printf("%c: %d, %d\n", plant, area, perimeter)
		cost += area * perimeter
	}
	return cost
}

func (d Solver) SolvePart2() int {
	//input := d.GetInput()
	return 0
}

type Point struct {
	x, y int
}

func (p Point) Adjacent() []Point {
	return []Point{
		{p.x - 1, p.y},
		{p.x + 1, p.y},
		{p.x, p.y - 1},
		{p.x, p.y + 1},
	}
}

func readGrid(input string) map[Point]rune {
	grid := make(map[Point]rune)
	for y, line := range strings.Split(input, "\n") {
		for x, plant := range line {
			p := Point{x, y}
			grid[p] = plant
		}
	}
	return grid
}

func findRegions(grid map[Point]rune) [][]Point {
	seen := make(map[Point]struct{})
	regions := make([][]Point, 0)
	for p, plant := range grid {
		if _, ok := seen[p]; !ok {
			region := []Point{p}
			region, seen = addAdjacent(p, plant, region, seen, grid)
			regions = append(regions, region)
		}

	}
	return regions
}

func addAdjacent(p Point, plant rune, region []Point, seen map[Point]struct{}, grid map[Point]rune) ([]Point, map[Point]struct{}) {
	for _, adj := range p.Adjacent() {
		if _, ok := seen[adj]; !ok {
			if grid[adj] == plant {
				seen[adj] = struct{}{}
				region = append(region, adj)
				region, seen = addAdjacent(adj, plant, region, seen, grid)
			}
		}
	}
	return region, seen
}

func computePerimeter(region []Point, plant rune, grid map[Point]rune) int {
	result := 0
	for _, p := range region {
		perimeter := 4
		for _, adj := range p.Adjacent() {
			if grid[adj] == plant {
				perimeter--
			}
		}
		result += perimeter
	}
	return result
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
