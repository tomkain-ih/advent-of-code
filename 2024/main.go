package main

import (
	"2024/day01"
	"2024/day02"
	"2024/day03"
	"2024/day04"
	"2024/day05"
	"2024/day06"
	"2024/day07"
	"2024/day08"
	"2024/day09"
	"2024/day10"
	"2024/day11"
	"2024/day12"
	"2024/day18"
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
		day07.Solver{},
		day08.Solver{},
		day09.Solver{},
		day10.Solver{},
		day11.Solver{},
		day12.Solver{},
		day18.Solver{MaxIndex: 70, Bytes: 1024},
	}

	solver := solvers[len(solvers)-1]
	fmt.Println(solver.GetLabel())
	if solver.SolvePart1() != 0 {
		fmt.Println(solver.SolvePart1())
	} else {
		fmt.Println("Part 1 solve")
	}
	if solver.SolvePart2() != 0 {
		fmt.Println(solver.SolvePart2())
	} else {
		fmt.Println("Part 2 solve")
	}
}
