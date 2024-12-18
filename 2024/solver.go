package main

type Solver interface {
	GetLabel() string
	SolvePart1() string
	SolvePart2() string
	GetInput() string
	GetFile() string
}
