package day05

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
	return "Day 05"
}

func (d Solver) SolvePart1() int {
	input := d.GetInput()
	rules := createRulesMap(input)
	updates := createUpdatesSlice(input)
	sum := 0
	for _, update := range updates {
		if verifySequence(update, rules) {
			sum += middleElem(update)
		}
	}
	return sum
}

func (d Solver) SolvePart2() int {
	input := d.GetInput()
	rules := createRulesMap(input)
	updates := createUpdatesSlice(input)
	sum := 0
	for _, update := range updates {
		if !verifySequence(update, rules) {
			corrected := correctSequence(update, rules)
			sum += middleElem(corrected)
		}
	}
	return sum
}

func correctSequence(update []int, rules map[int][]int) []int {
	for _, _ = range update {
		for j := 0; j < len(update)-1; j++ {
			if contains(rules[update[j+1]], update[j]) {
				update[j], update[j+1] = update[j+1], update[j]
			}
		}
	}
	return update
}

func contains(update []int, elem int) bool {
	for _, u := range update {
		if u == elem {
			return true
		}
	}
	return false
}

func middleElem(update []int) int {
	if len(update)%2 != 1 {
		fmt.Printf("expected odd number of elements, but got %d", len(update))
		fmt.Println(update)
		panic("expected odd number of elements")
	}
	return update[len(update)/2]
}

func createUpdatesSlice(input string) [][]int {
	result := [][]int{}
	for _, line := range strings.Split(input, "\n") {
		if strings.Contains(line, ",") {
			parts := strings.Split(line, ",")
			updates := make([]int, len(parts))
			for i, part := range parts {
				update, _ := strconv.Atoi(part)
				updates[i] = update
			}
			result = append(result, updates)
		}
	}
	return result
}

func (d Solver) GetFile() string {
	if d.File != "" {
		return d.File
	}
	return "day05/input.txt"
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

func createRulesMap(input string) map[int][]int {
	rules := make(map[int][]int)
	for _, line := range strings.Split(input, "\n") {
		if strings.Contains(line, "|") {
			parts := strings.Split(line, "|")
			left, right := parts[0], parts[1]
			leftInt, _ := strconv.Atoi(left)
			rightInt, _ := strconv.Atoi(right)
			rules[leftInt] = append(rules[leftInt], rightInt)
		}
	}
	return rules
}

func verifySequence(updates []int, rules map[int][]int) bool {
	for i, update := range updates {
		for _, next := range updates[i+1:] {
			for _, prior := range rules[next] {
				if prior == update {
					return false
				}
			}
		}

	}
	return true
}
