package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Feels good to be doin' stuff")

	tailSet := make(map[string]bool)
	headPos := GridPos{0, 0}
	tailPos := GridPos{0, 0}
	tailSet[tailPos.string()] = true
	for _, s := range exampleInput() {
		for _, i := range explodedInstruction(s) {
			instr := string(i)
			fmt.Print(headPos.string())
			fmt.Print(": ")
			fmt.Print(instr)

			headPos = headPos.move(instr)
			tailPos = tailPos.chase(headPos)
			tailSet[tailPos.string()] = true
			fmt.Print(" = ")
			fmt.Print(headPos.string())
			fmt.Println()
		}
	}

	fmt.Println(len(tailSet))

}

func explodedInstruction(instr string) string {
	d := string(instr[0])
	c, _ := strconv.Atoi(string(instr[2]))
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
