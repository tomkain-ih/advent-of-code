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
	/*//lines := exampleLines()
	lines := debugLines()
	//lines := readInputLines()
	total := countRedundantRangePairs(lines)
	fmt.Println(total)*/

	ints := convertToInts("14-80,13-20")
	doesContain()
}

func countRedundantRangePairs(lines []string) (total int) {
	for i := 0; i < len(lines); i++ {
		line := convertToInts(lines[i])
		redundant := "false"

		s1 := line[0][0]
		e1 := line[0][1]
		s2 := line[1][0]
		e2 := line[1][1]

		if doesContain(s1, e1, s2, e2) || doesContain(s2, e2, s1, e1) {
			total++
			redundant = "true"
		}

		fmt.Println(lines[i] + ": " + redundant)
	}
	return total
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

func doesContain(s1 int, e1 int, s2 int, e2 int) bool {
	fmt.Print(strconv.Itoa(s1))
	fmt.Print(strconv.Itoa(e1))
	fmt.Print(strconv.Itoa(s2))
	fmt.Print(strconv.Itoa(e2))
	result := s1 <= s2 && e1 >= e2
	fmt.Print(strconv.FormatBool(result))
	fmt.Println()
	return result
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
		"14-80,13-20", //false
		"39-78,40-40", //true
		"51-94,50-50", //false
		"27-84,27-85", //true
		"21-57,21-57", //true
		"80-87,87-90", //false
	}
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
	fmt.Println(ints)
	result[0] = [2]int{ints[0], ints[1]}
	result[1] = [2]int{ints[2], ints[3]}
	return result
}
