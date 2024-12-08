package day03

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Solver struct {
	Input string
	File  string
}

func getPart1RegexPattern() string {
	return `mul[(]\d+,\d+[)]`
}

func getPart2RegexPattern() string {
	return `don't[(][)]|do[(][)]|` + getPart1RegexPattern()
}

func getMatches(input string, pattern string) []string {
	re, err := regexp.Compile(pattern)
	if err != nil {
		log.Fatal("Error compiling regex pattern")
	}
	return re.FindAllString(input, -1)
}

func (d Solver) GetLabel() string {
	return "Day 03"
}

func (d Solver) SolvePart1() int {
	input := d.GetInput()
	matches := getMatches(input, getPart1RegexPattern())
	sum := 0
	for _, match := range matches {
		result := multiply(match)
		sum += result
	}
	return sum
}

func multiply(input string) int {
	substring := input[4 : len(input)-1]
	parts := strings.Split(substring, ",")
	num1, err := strconv.Atoi(parts[0])
	if err != nil {
		log.Fatal(err)
	}
	num2, err := strconv.Atoi(parts[1])
	if err != nil {
		log.Fatal(err)
	}
	return num1 * num2
}

func (d Solver) SolvePart2() int {
	input := d.GetInput()
	matches := getMatches(input, getPart2RegexPattern())
	sum := 0
	do := true
	for _, match := range matches {
		if strings.Contains(match, "don't") {
			do = false
		} else if strings.Contains(match, "do") {
			do = true
		} else if do {
			result := multiply(match)
			sum += result
		}
	}
	return sum
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
	return "day03/input.txt"
}
