package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	//lines := exampleLines()
	//lines := debugLines()
	lines := readInputLines()

	// Part 1
	fmt.Println(countRedundantRangePairs(lines))

	// Part 2
	fmt.Println(countOverlapRangePairs(lines))
}

func countOverlapRangePairs(lines []string) (total int) {
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		fmt.Print(line)
		if isOverlap(convertToInts(line)) {
			total++
			fmt.Print(" -- true")
		} else {
			fmt.Print(" -- false")
		}
		fmt.Println()
	}
	return total
}

func countRedundantRangePairs(lines []string) (total int) {
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if isRedundant(convertToInts(line)) {
			total++
		}
	}
	return total
}

func isOverlap(ints [2][2]int) bool {
	if ints[1][0] <= ints[0][1] && ints[1][1] >= ints[0][0] {
		return true
	}
	if ints[0][0] <= ints[1][1] && ints[0][1] >= ints[1][0] {
		return true
	}
	return false
}

func isRedundant(ints [2][2]int) bool {
	if ints[0][0] <= ints[1][0] && ints[0][1] >= ints[1][1] {
		return true
	}
	if ints[0][0] >= ints[1][0] && ints[0][1] <= ints[1][1] {
		return true
	}
	return false
}

func convertToInts(line string) (result [2][2]int) {
	line = strings.ReplaceAll(line, "-", " ")
	line = strings.Replace(line, ",", " ", 1)
	split := strings.Fields(line)
	var ints [4]int
	for i, each := range split {
		value, _ := strconv.Atoi(each)
		ints[i] = value
	}
	result[0] = [2]int{ints[0], ints[1]}
	result[1] = [2]int{ints[2], ints[3]}
	return result
}

func readInputLines() []string {
	var results []string
	file, err := os.Open("2022-04-input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		results = append(results, scanner.Text())
	}
	return results
}

func exampleLines() []string {
	return []string{
		"2-4,6-8",
		"2-3,4-5",
		"5-7,7-9",
		"2-8,3-7",
		"6-6,4-6",
		"2-6,4-8",
	}
}

func debugLines() []string {
	return []string{
		"14-80,13-20",
		"39-78,40-40",
		"51-94,50-50",
		"27-84,27-85",
		"21-57,21-57",
		"80-87,87-90",
		"51-94,50-50", //false
		"6-55,4-5",    //false
	}
}
