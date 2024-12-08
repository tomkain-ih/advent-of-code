package day07

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

func (d Solver) GetLabel() string {
	return "Day 07"
}

func (d Solver) SolvePart1() int {
	input := d.GetInput()
	sum := solve(input, false)
	return sum
}

func (d Solver) SolvePart2() int {
	input := d.GetInput()
	sum := solve(input, true)
	return sum
}

func solve(input string, concat bool) int {
	sum := 0
	for _, equation := range strings.Split(input, "\n") {
		sides := strings.Split(equation, ":")
		value, _ := strconv.Atoi(sides[0])
		numbers := getNumbers(sides[1])
		results := generateStream(numbers, concat)
		for _, r := range results {
			if r == value {
				sum += value
				break
			}
		}
	}
	return sum
}

func getNumbers(arg string) []int {
	var numbers []int
	for _, n := range strings.Split(arg, " ") {
		v, _ := strconv.Atoi(n)
		numbers = append(numbers, v)
	}
	return numbers
}

func generateStream(numbers []int, concat bool) []int {
	results := []int{numbers[0]}
	for i := 1; i < len(numbers); i++ {
		results = generate(results, numbers[i], concat)
	}
	return results
}

func generate(left []int, right int, concat bool) []int {
	var result []int
	for _, l := range left {
		result = append(result, l+right)
		result = append(result, l*right)
		if concat {
			v, _ := strconv.Atoi(strconv.Itoa(l) + strconv.Itoa(right))
			result = append(result, v)
		}
	}
	return result
}

func (d Solver) GetFile() string {
	if d.File != "" {
		return d.File
	}
	return "day07/input.txt"
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
