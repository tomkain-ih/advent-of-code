package day08

import (
	"log"
	"os"
	"strings"
)

type Solver struct {
	Input string
	File  string
}

type Point struct {
	X int
	Y int
}

func (d Solver) GetLabel() string {
	return "Day 08"
}

func (d Solver) SolvePart1() int {
	input := d.GetInput()
	antennas, bounds := readGrid(input)
	antinodes := make(map[Point]struct{})
	//for each value, capture the unique pairs of points
	for _, pairs := range getPairs(antennas) {
		//for each pair of points, calculate their two antinodes
		for _, antinode := range getAntiNodes(pairs) {
			//for each antinode if within bounds, increment antinode count
			if antinode.X >= 0 && antinode.X <= bounds.X && antinode.Y <= 0 && antinode.Y >= bounds.Y {
				antinodes[antinode] = struct{}{}
			}
		}
	}
	return len(antinodes)
}

func (d Solver) SolvePart2() int {
	input := d.GetInput()
	antennas, bounds := readGrid(input)
	antinodes := make(map[Point]struct{})
	for _, pairs := range getPairs(antennas) {
		antinodes[pairs[0]] = struct{}{}
		antinodes[pairs[1]] = struct{}{}
		for _, antinode := range getAntiNodesAndHarmonics(pairs, bounds) {
			antinodes[antinode] = struct{}{}
		}
	}
	return len(antinodes)
}

func (p Point) withinBounds(bounds Point) bool {
	return p.X >= 0 && p.X <= bounds.X && p.Y <= 0 && p.Y >= bounds.Y
}

func readGrid(input string) (map[rune][]Point, Point) {
	antennas := make(map[rune][]Point)
	var last Point
	for y, line := range strings.Split(input, "\n") {
		for x, val := range line {
			if val != '.' {
				antennas[val] = append(antennas[val], Point{x, -y})
			}
		}
		last = Point{len(line) - 1, -y}
	}
	return antennas, last
}

func getPairs(antennas map[rune][]Point) [][]Point {
	var pairs [][]Point
	for _, v := range antennas {
		for i := 0; i < len(v); i++ {
			for j := i + 1; j < len(v); j++ {
				pair := []Point{v[i], v[j]}
				pairs = append(pairs, pair)
			}
		}
	}
	return pairs
}

func getAntiNodesAndHarmonics(pair []Point, bounds Point) []Point {
	if len(pair) != 2 {
		log.Fatal("Expected 2 values in pair")
	}
	var result []Point
	node1 := pair[0]
	node2 := pair[1]
	yDist := node1.Y - node2.Y
	xDist := node1.X - node2.X

	for {
		node1 = Point{node1.X + xDist, node1.Y + yDist}
		node2 = Point{node2.X - xDist, node2.Y - yDist}
		if !node1.withinBounds(bounds) && !node2.withinBounds(bounds) {
			break
		}

		if node1.withinBounds(bounds) {
			result = append(result, node1)
		}
		if node2.withinBounds(bounds) {
			result = append(result, node2)
		}
	}

	return result
}

func getAntiNodes(pair []Point) []Point {
	if len(pair) != 2 {
		log.Fatal("Expected 2 values in pair")
	}
	var result []Point
	yDist := pair[0].Y - pair[1].Y
	xDist := pair[0].X - pair[1].X

	result = append(result, Point{pair[0].X + xDist, pair[0].Y + yDist})
	result = append(result, Point{pair[1].X - xDist, pair[1].Y - yDist})
	return result
}

func (d Solver) GetFile() string {
	if d.File != "" {
		return d.File
	}
	return "day08/input.txt"
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
