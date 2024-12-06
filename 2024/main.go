package main

import (
	"2024/day01"
	"2024/day02"
	"2024/day03"
	"2024/day04"
	"2024/day05"
	"2024/day06"
	"fmt"
)

func main() {
	solvers := []Solver{
		day01.Solver{},
		day02.Solver{},
		day03.Solver{},
		day04.Solver{},
		day05.Solver{},
		day06.Solver{},
	}

	solver := solvers[len(solvers)-1]
	fmt.Println(solver.GetLabel())
	fmt.Println(solver.SolvePart1())
	fmt.Println(solver.SolvePart2())
}
