package main

type Solver interface {
	GetLabel() string
	SolvePart1() int
	SolvePart2() int
	GetInput() string
	GetFile() string
}
