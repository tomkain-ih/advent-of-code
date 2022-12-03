package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func readFile() []string {
	var results []string
	file, err := os.Open("2022-02-input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		results = append(results, scanner.Text())
	}
	return results
}

func testSpecific(opponent string, my string) int {
	opponentPlay := new(Play)
	opponentPlay.NewPlay(opponent)

	myPlay := new(Play)
	myPlay.NewPlay(my)

	return myPlay.Versus(opponentPlay)
}

func testExample() int {
	result := testSpecific("A", "Y")
	result = result + testSpecific("B", "X")
	return result + testSpecific("C", "Z")
}

func computeFromFile() int {
	input := readFile()
	total := 0
	for i := 0; i < len(input); i++ {
		abbrevs := strings.Fields(input[i])
		opponent := new(Play)
		opponent.NewPlay(abbrevs[0])
		myPlay := new(Play)
		myPlay.NewPlay(abbrevs[1])

		result := myPlay.Versus(opponent)
		//fmt.Println(result)
		total = total + result
	}
	return total
}

func testPart2Specific(opponentAbrev string, result string) int {
	instruction := new(Instruction)
	instruction.NewInstruction(opponentAbrev, result)
	return instruction.getScore()
}

func testPart2Example() int {
	result := testPart2Specific("A", "Y")
	result = result + testPart2Specific("B", "X")
	return result + testPart2Specific("C", "Z")
}

func computePart2FromFile() int {
	input := readFile()
	total := 0
	for i := 0; i < len(input); i++ {
		abbrevs := strings.Fields(input[i])
		instruction := new(Instruction)
		instruction.NewInstruction(abbrevs[0], abbrevs[1])
		total = total + instruction.getScore()
	}
	return total
}

func testAllCombos() {
	fmt.Println(testSpecific("A", "X")) //Rock ties rock - 4
	fmt.Println(testSpecific("A", "Y")) //paper beats rock - 8
	fmt.Println(testSpecific("A", "Z")) //scissors loses to rock - 3

	fmt.Println(testSpecific("B", "X")) //rock loses to paper - 1
	fmt.Println(testSpecific("B", "Y")) //paper ties paper - 5
	fmt.Println(testSpecific("B", "Z")) //scissors beats paper - 9

	fmt.Println(testSpecific("C", "X")) //rock beats scissors - 7
	fmt.Println(testSpecific("C", "Y")) //paper loses to scissors - 2
	fmt.Println(testSpecific("C", "Z")) //scissors ties scissors - 6
}

func main() {
	// Part 1
	//fmt.Println(testExample())
	//testAllCombos()
	fmt.Println(computeFromFile())

	// Part 2
	//fmt.Println(testPart2Example())
	fmt.Println(computePart2FromFile())
}
