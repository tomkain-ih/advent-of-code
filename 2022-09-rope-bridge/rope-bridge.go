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
	//input := largeExampleInput()
	input := input()

	// Part 1
	fmt.Println(countTailPositions(input))

	// Part 2
	fmt.Println(countMultiTailPositions(input, 9)) //2617 incorrect
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

func countMultiTailPositions(input []string, tailCount int) int {
	tailSet := make(map[string]bool)
	headPos := GridPos{0, 0}
	tails := generateMultiTailSlice(tailCount)
	for _, s := range input {
		for _, i := range explodedInstruction(s) {
			instr := string(i)
			headPos = headPos.move(instr)
			lastTail := multiTailChase(headPos, tails)
			tailSet[lastTail.string()] = true
		}
		fmt.Println(s)
		fmt.Print(headPos.string())
		for _, tail := range tails {
			fmt.Print("; " + tail.string())
		}
		fmt.Println()
	}
	return len(tailSet)
}

func generateMultiTailSlice(count int) (result []GridPos) {
	for i := 0; i < count; i++ {
		result = append(result, GridPos{0, 0})
	}
	return
}

func multiTailChase(head GridPos, tails []GridPos) (lastTail GridPos) {
	lastTail = head
	for i, tail := range tails {
		tail = tail.chase(lastTail)
		tails[i] = tail
		lastTail = tail
	}
	return
}

func countTailPositions(input []string) int {
	tailSet := make(map[string]bool)
	headPos := GridPos{0, 0}
	tailPos := GridPos{0, 0}
	tailSet[tailPos.string()] = true
	for _, s := range input {
		for _, i := range explodedInstruction(s) {
			instr := string(i)
			headPos = headPos.move(instr)
			tailPos = tailPos.chase(headPos)
			tailSet[tailPos.string()] = true
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

func largeExampleInput() []string {
	return strings.Split(`R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20`, "\n")
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
