package day02

import (
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Solver struct {
	Input string
	File  string
}

func (d Solver) GetLabel() string {
	return "Day 02"
}

func (d Solver) SolvePart1() string {
	input := d.GetInput()
	safe := 0
	// split by line break
	for _, line := range strings.Split(input, "\n") {
		// process by line - boolean isSafe
		if isSafe(line) {
			// count safe
			safe++
		}
	}
	return strconv.Itoa(safe)
}

func (d Solver) SolvePart2() string {
	input := d.GetInput()
	safe := 0
	// split by line break
	for _, line := range strings.Split(input, "\n") {
		// process by line - boolean isSafe
		if isSafeWithTolerance(line) {
			// count safe
			safe++
		}
	}
	return strconv.Itoa(safe)
}

func isSafe(line string) bool {
	var prior *int
	decreasing := false
	increasing := false
	for _, num := range strings.Fields(line) {
		// convert to int
		n, err := strconv.Atoi(num)
		if err != nil {
			log.Fatal(err)
		}
		if prior != nil {
			if n < *prior {
				decreasing = true
			}
			if n > *prior {
				increasing = true
			}
			diff := math.Abs(float64(n - *prior))
			if diff < 1 || diff > 3 {
				return false
			}
		}
		// set prior to current
		prior = &n
	}
	if decreasing && increasing {
		return false
	}
	return true
}

func isSafeWithTolerance(line string) bool {
	levels := toIntSlice(line)
	for i := 0; i < len(levels); i++ {
		l := toIntSlice(line) // needed to prevent levels from being modified on append call
		s := append(l[:i], l[i+1:]...)
		if isSafe(toString(s)) {
			return true
		}
	}
	return false
}

func toString(s []int) string {
	strSlice := make([]string, len(s))
	for i, num := range s {
		strSlice[i] = strconv.Itoa(num)
	}
	return strings.Join(strSlice, " ")
}

func toIntSlice(line string) []int {
	parts := strings.Fields(line)
	levels := make([]int, len(parts))
	for i, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			log.Fatal(err)
		}
		levels[i] = num
	}
	return levels
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
	return "day02/input.txt"
}
