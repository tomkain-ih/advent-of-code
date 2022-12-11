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
	//input := exampleInput()
	input := input()
	fmt.Println(countTailPositions(input))
	//singleMove()

}

func singleMove() {
	instr := "L"
	headPos := GridPos{4, 3}
	tailPos := GridPos{3, 4}

	headPos = headPos.move(instr)
	tailPos = tailPos.chase(headPos)
	fmt.Print(instr)
	fmt.Print(": ")
	fmt.Print(headPos.string())
	fmt.Print(" - ")
	fmt.Print(tailPos.string())
	fmt.Println()
}

func countTailPositions(input []string) int {
	tailSet := make(map[string]bool)
	headPos := GridPos{0, 0}
	tailPos := GridPos{0, 0}
	tailSet[tailPos.string()] = true
	for _, s := range input {
		fmt.Println(s)
		for _, i := range explodedInstruction(s) {
			instr := string(i)
			headPos = headPos.move(instr)
			tailPos = tailPos.chase(headPos)
			tailSet[tailPos.string()] = true
			fmt.Print(instr)
			fmt.Print(": ")
			fmt.Print(headPos.string())
			fmt.Print(" - ")
			fmt.Print(tailPos.string())
			fmt.Println()
		}
	}

	return len(tailSet)
}

func explodedInstruction(instr string) string {
	d := string(instr[0])
	c, _ := strconv.Atoi(instr[2:])
	repeat := strings.Repeat(d, c)
	return repeat
}

func exampleInput() []string {
	return strings.Split(`R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`, "\n")
}

func input() (result []string) {
	file, err := os.Open("2022-09-input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	return
}
