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

func (d Solver) SolvePart1() string {
	// create grid
	grid := makeGrid(d.getMaxIndex())

	// read input as grid points
	corruptions := readInput(d.GetInput())

	// remove bytes count of input points from grid
	for _, p := range corruptions[:d.getBytes()] {
		delete(grid, p)
	}

	// find first route, which should be shortest
	return strconv.Itoa(route(grid, d.getMaxIndex()))
}

func (d Solver) SolvePart2() string {
	grid := makeGrid(d.MaxIndex)
	corruptions := readInput(d.GetInput())
	for _, p := range corruptions[:d.getBytes()] {
		delete(grid, p)
	}
	var result Point
	for i := d.getBytes(); i < len(corruptions); i++ {
		corrupt := corruptions[i]
		delete(grid, corrupt)
		if s := route(grid, d.MaxIndex); s == -1 {
			result = corrupt
		}
	}
	return result.String()
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

func readInput(input string) []Point {
	var points []Point
	for _, line := range strings.Split(input, "\n") {
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
	return -1
}

type Point struct {
	x, y int
}

func (p Point) String() string {
	return strconv.Itoa(p.x) + "," + strconv.Itoa(p.y)
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

func (d Solver) getBytes() int {
	if d.Bytes != 0 {
		return d.Bytes
	}
	return 1024
}

func (d Solver) getMaxIndex() int {
	if d.MaxIndex != 0 {
		return d.MaxIndex
	}
	return 70
}
