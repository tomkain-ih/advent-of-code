package main

type Solver interface {
	GetLabel() string
	Solve() string
	GetInput() string
}
