package day12

import (
	"fmt"
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
	return "Day 12"
}

func (d Solver) GetFile() string {
	if d.File != "" {
		return d.File
	}
	return "day12/input.txt"
}

func (d Solver) SolvePart1() string {
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
		cost += area * perimeter
	}
	return strconv.Itoa(cost)
}

func (d Solver) SolvePart2() string {
	input := d.GetInput()
	grid := readGrid(input)
	regions := findRegions(grid)

	cost := 0
	for _, region := range regions {
		// for each region, multiply area by sides and sum products
		area := len(region)
		plant := grid[region[0]]
		sides := computeSides(region)
		fmt.Printf("%s: %d * %d\n", string(plant), area, sides)
		cost += area * sides
	}
	return strconv.Itoa(cost)
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

type ByRowThenCol []Point

func (a ByRowThenCol) Len() int      { return len(a) }
func (a ByRowThenCol) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByRowThenCol) Less(i, j int) bool {
	if a[i].y == a[j].y {
		return a[i].x < a[j].x
	}
	return a[i].y < a[j].y
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
			var region []Point
			region, seen = addAdjacent(p, plant, region, seen, grid)
			if len(region) == 0 {
				region = []Point{p}
			}
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

func computeSides(region []Point) int {
	//TODO finish
	/*sorted := make([]Point, len(region))
	copy(sorted, region)
	sort.Sort(ByRowThenCol(sorted))
	x := sorted[0].x
	y := sorted[0].y
	sides := 4
	for i := 1; i < len(sorted); i++ {

	}*/
	return 0
}

/*func computeSides(region []Point) int {
	sides := 4
	bounds := make(map[int][]int)
	rowSet := make(map[int]struct{})
	for _, p := range region {
		rowSet[p.y] = struct{}{}
		if _, ok := bounds[p.y]; !ok {
			bounds[p.y] = []int{p.x, p.x}
		} else {
			if p.x < bounds[p.y][0] {
				bounds[p.y][0] = p.x
			}
			if p.x > bounds[p.y][1] {
				bounds[p.y][1] = p.x
			}
		}
	}
	var rows []int
	for row := range rowSet {
		rows = append(rows, row)
	}
	sort.Ints(rows)
	a, b := bounds[rows[0]][0], bounds[rows[0]][1]
	for i := 1; i < len(rows); i++ {
		l := bounds[rows[i]][0]
		r := bounds[rows[i]][1]
		if l != a {
			sides += 2
		}
		if r != b {
			sides += 2
		}
		a, b = l, r
	}
	return sides
}*/

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
