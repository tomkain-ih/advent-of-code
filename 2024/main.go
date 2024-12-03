package main

import (
	"2024/day01"
	"2024/day02"
	"2024/day03"
	"fmt"
)

func main() {
	solvers := []Solver{
		day01.Solver{},
		day02.Solver{},
		day03.Solver{},
		// Add more solvers as needed, e.g., day02.Day02Solver{}
	}

	solver := solvers[len(solvers)-1]
	fmt.Println(solver.GetLabel())
	fmt.Println(solver.SolvePart1())
	fmt.Println(solver.SolvePart2())
}
