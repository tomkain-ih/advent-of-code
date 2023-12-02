package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

// Part 1
func findMaxElfCalories() int {
	file, err := os.Open("2022-01-input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	maxSeen := 0
	current := 0
	elfNumber := 0

	for scanner.Scan() {
		textLine := scanner.Text()
		if len(textLine) > 0 {
			i, err := strconv.Atoi(textLine)
			if err != nil {
				log.Fatal(err)
			}
			current = current + i
		} else {
			elfNumber++
			fmt.Println("----- Elf Number " + strconv.Itoa(elfNumber) + "-----")
			fmt.Println("current = " + strconv.Itoa(current))
			maxSeen = max(maxSeen, current)
			fmt.Println("max seen = " + strconv.Itoa(maxSeen))
			current = 0
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return maxSeen
}

// Part 2
func findAndSumTopThreeElfCalories() int {
	file, err := os.Open("2022-01-input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	var elfTotals []int
	current := 0

	for scanner.Scan() {
		textLine := scanner.Text()
		if len(textLine) > 0 {
			i, err := strconv.Atoi(textLine)
			if err != nil {
				log.Fatal(err)
			}
			current = current + i
		} else {
			fmt.Println("Appending " + strconv.Itoa(current))
			elfTotals = append(elfTotals, current)
			fmt.Println(strconv.Itoa(len(elfTotals)) + " elves seen")
			current = 0
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Ints(elfTotals)
	return sum(elfTotals)
}

func sum(slice []int) int {
	result := 0
	size := len(slice) - 1
	for i := size; i > size-3; i-- {
		value := slice[i]
		fmt.Println("Top elf value = " + strconv.Itoa(value))
		result = result + value
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	//fmt.Println(findMaxElfCalories())
	fmt.Println(findAndSumTopThreeElfCalories())
}
