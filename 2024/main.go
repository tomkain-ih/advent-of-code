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
	"2024/day19"
	"2024/day20"
	"2024/day21"
	"2024/day22"
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
		day18.Solver{},
		day19.Solver{},
		day20.Solver{},
		day21.Solver{},
		day22.Solver{},
	}

	solver := solvers[len(solvers)-1]
	fmt.Println(solver.GetLabel())
	fmt.Println(solver.SolvePart1())
	fmt.Println(solver.SolvePart2())
}
