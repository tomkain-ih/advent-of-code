package day01

import (
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Solver struct {
	Input string
	File  string
}

func (d Solver) GetLabel() string {
	return "Day 01"
}

func (d Solver) SolvePart1() int {
	input := d.GetInput()
	A, B := colwiseSplit(input)
	sort.Ints(A)
	sort.Ints(B)
	result := rowWiseAbsDiffSummed(A, B)
	return result
}

func (d Solver) SolvePart2() int {
	input := d.GetInput()
	A, B := colwiseSplit(input)
	sum := 0
	for _, a := range A {
		count := countOccurrences(B, a)
		sum += a * count
	}
	return sum
}

func colwiseSplit(data string) ([]int, []int) {
	parts := strings.Fields(data)

	// Convert the strings to integers
	numbers := make([]int, len(parts))
	for i, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			log.Fatal(err)
		}
		numbers[i] = num
	}

	// Separate the numbers into two slices
	var A, B []int
	for i, num := range numbers {
		if i%2 == 0 {
			A = append(A, num)
		} else {
			B = append(B, num)
		}
	}

	return A, B
}

func rowWiseAbsDiffSummed(a []int, b []int) int {
	sum := 0.0
	for i := 0; i < len(a); i++ {
		sum += math.Abs(float64(a[i] - b[i]))
	}
	return int(sum)
}

func countOccurrences(B []int, x int) int {
	count := 0
	for _, b := range B {
		if b == x {
			count++
		}
	}
	return count
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
	return "day01/input.txt"
}
