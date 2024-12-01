package main

import (
	"2024/day01"
	"fmt"
)

func main() {
	solvers := []Solver{
		day01.Day01Solver{},
		// Add more solvers as needed, e.g., day02.Day02Solver{}
	}

	solver := solvers[len(solvers)-1]
	fmt.Println(solver.GetLabel())
	fmt.Println(solver.Solve())
}
