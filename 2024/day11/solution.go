package day11

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
	return "Day 11"
}

func (d Solver) GetFile() string {
	if d.File != "" {
		return d.File
	}
	return "day11/input.txt"
}

func (d Solver) SolvePart1() string {
	input := d.GetInput()
	stones := blink(input, 25)
	return strconv.Itoa(len(stones))
}

func (d Solver) SolvePart2() string {
	input := d.GetInput()
	stones := betterBlink(input, 75)
	return strconv.Itoa(sumValues(stones))
}

func sumValues(stones map[int]int) int {
	sum := 0
	for _, v := range stones {
		sum += v
	}
	return sum
}

func betterBlink(input string, times int) map[int]int {
	array := fromInput(input)
	stones := make(map[int]int)
	for _, stone := range array {
		if _, ok := stones[stone]; ok {
			stones[stone]++
		} else {
			stones[stone] = 1
		}
	}

	for i := 0; i < times; i++ {
		result := make(map[int]int)
		for oldStone, oldCount := range stones {
			for _, newStone := range applyRules(oldStone) {
				if _, ok := result[newStone]; ok {
					result[newStone] += oldCount
				} else {
					result[newStone] = oldCount
				}
			}
		}
		stones = result
	}
	return stones
}

func blink(input string, times int) []int {
	// convert input to int array
	stones := fromInput(input)
	// for number of iterations
	for i := 0; i < times; i++ {
		var result []int
		// for each int in array, apply rules, appending to new array
		for _, stone := range stones {
			result = append(result, applyRules(stone)...)
		}
		// replace old array with new array
		stones = result
	}
	return stones
}

func fromInput(input string) []int {
	var stones []int
	for _, v := range strings.Split(input, " ") {
		i, _ := strconv.Atoi(v)
		stones = append(stones, i)
	}
	return stones

}

func applyRules(stone int) []int {
	if stone == 0 {
		return []int{1}
	}
	str := strconv.Itoa(stone)
	l := len(str)
	if l%2 == 0 {
		s1, _ := strconv.Atoi(str[:l/2])
		s2, _ := strconv.Atoi(str[l/2:])
		return []int{s1, s2}
	}
	return []int{stone * 2024}

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
