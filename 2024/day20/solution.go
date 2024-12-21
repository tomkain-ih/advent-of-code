package day20

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Solver struct {
	Input     string
	File      string
	Threshold int
}

func (d Solver) GetLabel() string {
	return "Day 20"
}

func (d Solver) GetFile() string {
	if d.File != "" {
		return d.File
	}
	return "day20/input.txt"
}

func (d Solver) GetThreshold() int {
	if d.Threshold != 0 {
		return d.Threshold
	}
	return 100
}

func (d Solver) SolvePart1() string {
	// read grid, start, end, track, and wall (map[point]rune)
	grid, begin := readGrid(d.GetInput())
	// solve puzzle start -> end: ordered path of points
	path := findPath(grid, false, begin) // false = don't ignore first wall
	// for each path point, for each adjacent wall point
	steps := len(path)
	cheats := make(map[int]int) // steps saved -> count
	starts := make(map[point]struct{})
	for i, p := range path {
		for _, a := range p.adjacent() {
			if _, ok := starts[a]; !ok {
				if grid[a] == wall {
					x := point{10, 7}
					if a == x {
						fmt.Println("here")
					}
					// solve wall point -> end
					alt := findPath(grid, true, a) // true = ignore first wall
					// number of steps compared to init len(path)
					altSteps := len(alt) + i - 1
					diff := steps - altSteps
					if diff > 0 {
						cheats[diff]++
					}
				}
				starts[a] = struct{}{}
			}
		}
	}

	count := 0
	for k, v := range cheats {
		fmt.Println(v, k)
		if k >= d.GetThreshold() {
			count += v
		}
	}
	return strconv.Itoa(count)
}

func (d Solver) SolvePart2() string {
	return "Part 2 Solve"
}

func readGrid(input string) (map[point]rune, point) {
	grid := make(map[point]rune)
	var start_ point
	for y, line := range strings.Split(input, "\n") {
		for x, r := range line {
			p := point{x, y}
			grid[p] = r
			if r == start {
				start_ = p
			}
		}
	}
	return grid, start_
}

func findPath(grid map[point]rune, ignoreFirstWall bool, begin point) []point {
	var queue []trail
	path := trail{visited: make(map[point]struct{}), last: begin}
	queue = append(queue, path)
	var fastest trail
	for len(queue) > 0 {
		path = queue[0]
		queue = queue[1:]
		for _, a := range path.last.adjacent() {
			if _, ok := path.visited[a]; !ok {
				if grid[a] == end {
					newTrail := path.add(a)
					if len(fastest.points) == 0 || len(newTrail.points) < len(fastest.points) {
						fastest = newTrail
					}
				}

				if grid[a] == track {
					queue = append(queue, path.add(a))
				}

				if grid[a] == wall && ignoreFirstWall && len(path.points) == 0 {
					queue = append(queue, path.add(a))
				}

			}
		}
	}
	return fastest.points
}

type point struct {
	x, y int
}

type trail struct {
	points  []point
	visited map[point]struct{}
	last    point
}

func (t trail) add(p point) trail {
	newTrail := t.copy()
	newTrail.points = append(newTrail.points, p)
	newTrail.visited[p] = struct{}{}
	newTrail.last = p
	return newTrail
}

func (t trail) copy() trail {
	visited := make(map[point]struct{})
	for k := range t.visited {
		visited[k] = struct{}{}
	}
	return trail{
		points:  append([]point{}, t.points...),
		visited: visited,
		last:    t.last,
	}
}

func (p point) adjacent() []point {
	return []point{
		{p.x, p.y - 1},
		{p.x, p.y + 1},
		{p.x - 1, p.y},
		{p.x + 1, p.y},
	}
}

const wall = '#'
const track = '.'
const start = 'S'
const end = 'E'

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
