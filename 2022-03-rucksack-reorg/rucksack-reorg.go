package main

import (
	"bufio"
	"errors"
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
	// Part 1
	fmt.Println(calculateDuplicativeCharacterValue(readInputLines()))

	// Part 2
	fmt.Println(calculateCommonCharacterValueAcrossGroupsOfRucksacks(readInputLines(), 3))
}

func calculateCommonCharacterValueAcrossGroupsOfRucksacks(lines []string, groupSize int) (total int) {
	for i := 0; i < len(lines); i += groupSize {
		end := i + groupSize
		if end > len(lines) {
			end = len(lines)
		}

		character, err := findCommonCharacter(lines[i:end])
		if err != nil {
			fmt.Println(err)
		} else {
			total = total + characterValue(character)
		}
	}
	return total
}

func findCommonCharacter(group []string) (rune, error) {
	var common []rune
	for _, rucksack := range group {
		r := []rune(rucksack)
		if common == nil {
			common = r
		} else {
			common = Intersection(common, r)
		}
	}
	if len(common) != 1 {
		return 'a', errors.New("problem finding single common character")
	}
	return common[0], nil

}

func Intersection(a, b []rune) (c []rune) {
	m := make(map[rune]int)

	for _, item := range a {
		m[item] = 1
	}

	for _, item := range b {
		if _, ok := m[item]; ok {
			i := m[item]
			m[item] = i + 1
		}
	}
	for k, e := range m {
		if e > 1 {
			c = append(c, k)
		}
	}
	return
}

func calculateDuplicativeCharacterValue(lines []string) int {
	total := 0
	for _, line := range lines {
		chars := []rune(line)
		midpoint := len(chars) / 2
		firstHalf := string(chars[0:midpoint])
		secondHalf := chars[midpoint:]

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
