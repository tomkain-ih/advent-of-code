package day19

import (
	"fmt"
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
	return "Day 19"
}

func (d Solver) GetFile() string {
	if d.File != "" {
		return d.File
	}
	return "day19/input.txt"
}

func (d Solver) SolvePart1() string {
	towels, patterns := parseInput(d.GetInput())
	possible := 0
	for _, pattern := range patterns {
		if makes(pattern, towels) > 0 {
			possible++
		}
	}
	return strconv.Itoa(possible)
}

func (d Solver) SolvePart2() string {
	towels, patterns := parseInput(d.GetInput())
	possible := 0
	for _, pattern := range patterns {
		m := makes(pattern, towels)
		fmt.Println(pattern, m)
		possible += m
	}
	return strconv.Itoa(possible)
}

func parseInput(input string) ([]string, []string) {
	var towels []string
	var patterns []string

	lines := strings.Split(input, "\n")
	for _, towel := range strings.Split(lines[0], ", ") {
		towels = append(towels, towel)
	}

	for _, pattern := range lines[2:] {
		patterns = append(patterns, pattern)
	}
	return towels, patterns
}

func makes(pattern string, towels []string) int {
	var queue []string
	queue = append(queue, pattern)
	seen := make(map[string]int)
	makes := 0
	for len(queue) > 0 {
		sub := queue[0]
		queue = queue[1:]
		count := 1
		if _, ok := seen[sub]; ok {
			count = seen[sub]
		}
		for _, towel := range towels {
			// subpattern matches towel
			if sub == towel {
				makes += count
			}
			//subpattern starts with towel
			if len(sub) > len(towel) {
				if towel == sub[:len(towel)] {
					s := sub[len(towel):]
					if _, ok := seen[s]; !ok {
						queue = append(queue, s)
					}
					seen[s] += count
				}
			}
		}
		seen[sub] = count + 1
	}
	return makes
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
