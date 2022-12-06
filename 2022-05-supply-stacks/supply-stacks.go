package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// Part 1
	//stacks := readExampleStacks()
	stacks := readInputStacks()
	//rawInstrs := readExampleInstructions()
	rawInstrs := readInstructionsFromFile()
	executeInstructions(stacks, rawInstrs)
	fmt.Println(getTopCrates(stacks))

}

func getTopCrates(stacks []*Stack) (message string) {
	for _, stack := range stacks {
		value := stack.Pop().Value
		message = message + charValue(value)
	}
	return
}

func executeInstructions(stacks []*Stack, instrs []string) {
	for _, r := range instrs {
		instr := new(StackInstruction)
		instr.NewStackInstruction(r)
		//fmt.Println(instr.String())
		instr.Execute(stacks)
	}
}

func readInputStacks() []*Stack {
	//cheated, prepped these "by hand"
	return []*Stack{
		CreateStack([]rune{'D', 'T', 'W', 'F', 'J', 'S', 'H', 'N'}),
		CreateStack([]rune{'H', 'R', 'P', 'Q', 'T', 'N', 'B', 'G'}),
		CreateStack([]rune{'L', 'Q', 'V'}),
		CreateStack([]rune{'N', 'B', 'S', 'W', 'R', 'Q'}),
		CreateStack([]rune{'N', 'D', 'F', 'T', 'V', 'M', 'B'}),
		CreateStack([]rune{'M', 'D', 'B', 'V', 'H', 'T', 'R'}),
		CreateStack([]rune{'D', 'B', 'Q', 'J'}),
		CreateStack([]rune{'D', 'N', 'J', 'V', 'R', 'Z', 'H', 'Q'}),
		CreateStack([]rune{'B', 'N', 'H', 'M', 'S'}),
	}
}

func charValue(i int) string {
	return string(rune(i))
}

func readExampleStacks() []*Stack {
	return []*Stack{
		CreateStack([]rune{'Z', 'N'}),
		CreateStack([]rune{'M', 'C', 'D'}),
		CreateStack([]rune{'P'}),
	}
}

func CreateStack(chars []rune) *Stack {
	s := NewStack()
	for _, char := range chars {
		s.Add(int(char))
	}
	return s
}

func readExampleInstructions() []string {
	return []string{
		"move 1 from 2 to 1",
		"move 3 from 1 to 3",
		"move 2 from 2 to 1",
		"move 1 from 1 to 2",
	}
}

func stackExample() {
	s := NewStack()
	s.Push(&Node{1})
	s.Push(&Node{2})
	s.Push(&Node{3})
	fmt.Println(s.Pop(), s.Pop(), s.Pop())

	q := NewQueue(1)
	q.Push(&Node{4})
	q.Push(&Node{5})
	q.Push(&Node{6})
	fmt.Println(q.Pop(), q.Pop(), q.Pop())
}

func readInstructionsFromFile() []string {
	var results []string
	file, err := os.Open("2022-05-input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "move") {
			results = append(results, line)
		}
	}
	return results
}
