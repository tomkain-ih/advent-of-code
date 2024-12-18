package day18

import (
	"log"
	"os"
	"strconv"
	"strings"
)

type Solver struct {
	Input    string
	File     string
	MaxIndex int
	Bytes    int
}

func (d Solver) GetLabel() string {
	return "Day 18"
}

func (d Solver) GetFile() string {
	if d.File != "" {
		return d.File
	}
	return "day18/input.txt"
}

func (d Solver) SolvePart1() int {
	// create grid
	grid := makeGrid(d.MaxIndex)

	// read input as grid points for Bytes count points
	corruptions := readInput(d.GetInput(), d.Bytes)

	// remove input points from grid
	for _, p := range corruptions {
		delete(grid, p)
	}

	// find first route, which should be shortest
	return route(grid, d.MaxIndex)
}

func (d Solver) SolvePart2() int {
	//input := d.GetInput()
	return 0
}

func makeGrid(size int) map[Point]struct{} {
	grid := make(map[Point]struct{})
	for x := 0; x <= size; x++ {
		for y := 0; y <= size; y++ {
			grid[Point{x, y}] = struct{}{}
		}
	}
	return grid
}

func readInput(input string, lines int) []Point {
	var points []Point
	split := strings.Split(input, "\n")
	for i := 0; i < lines; i++ {
		line := split[i]
		numerals := strings.Split(line, ",")
		x, _ := strconv.Atoi(numerals[0])
		y, _ := strconv.Atoi(numerals[1])
		points = append(points, Point{x, y})
	}
	return points

}

func route(grid map[Point]struct{}, size int) int {
	start := Point{0, 0}
	end := Point{size, size}
	queue := []PointDist{{start, 0}}
	visited := map[Point]struct{}{}
	for {
		if len(queue) == 0 {
			break
		}

		p := queue[0]
		queue = queue[1:]
		point := p.p
		dist := p.dist
		if point == end {
			return dist
		}

		if _, ok := visited[point]; !ok {
			for _, adj := range point.adjacent() {
				if _, ok := grid[adj]; ok {
					queue = append(queue, PointDist{adj, dist + 1})
				}
			}
		}
		visited[point] = struct{}{}
	}
	return 0
}

type Point struct {
	x, y int
}

func (p Point) adjacent() []Point {
	return []Point{
		{p.x - 1, p.y},
		{p.x + 1, p.y},
		{p.x, p.y - 1},
		{p.x, p.y + 1},
	}
}

type PointDist struct {
	p    Point
	dist int
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
