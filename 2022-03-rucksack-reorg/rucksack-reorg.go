package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// for each line
// split input in half
// find single character seen in both halves
// calculate value of duplicative character
// sum values
func main() {
	//inputLines := readInputLines()
	inputLines := exampleInputLines()
	total := calculateDuplicativeCharacterValue(inputLines)
	fmt.Println(total)
}

func calculateDuplicativeCharacterValue(lines []string) int {
	total := 0
	for _, line := range lines {
		chars := []rune(line)
		midpoint := len(chars) / 2
		firstHalf := string(chars[0 : midpoint+1])
		secondHalf := chars[midpoint : len(chars)-1]

		for _, c := range secondHalf {
			if strings.Contains(firstHalf, string(c)) {
				total = total + characterValue(c)
				break
			}
		}
	}
	return total
}

func characterValue(c rune) int {
	//https://stackoverflow.com/questions/21322173/convert-rune-to-int
	//https://www.asciitable.com/
	value := int(c)
	if value < 91 {
		return value - 38
	}
	if value >= 97 {
		return value - 96
	}
	return 0
}

func exampleInputLines() []string {
	return []string{
		"vJrwpWtwJgWrhcsFMMfFFhFp",
		"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
		"PmmdzqPrVvPwwTWBwg",
		"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
		"ttgJtRGJQctTZtZT",
		"CrZsJsPPZsGzwwsLwLmpwMDw",
	}
}

func readInputLines() []string {
	var results []string
	file, err := os.Open("2022-03-input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		results = append(results, scanner.Text())
	}
	return results
}
