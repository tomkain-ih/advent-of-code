package main

import (
	"log"
	"strconv"
)

type GridPos struct {
	Row int
	Col int
}

func (g GridPos) string() string {
	return strconv.Itoa(g.Row) + "," + strconv.Itoa(g.Col)
}

func (g GridPos) move(d string) GridPos {
	if d == "U" {
		return GridPos{g.Row + 1, g.Col}
	}
	if d == "D" {
		return GridPos{g.Row - 1, g.Col}
	}
	if d == "L" {
		return GridPos{g.Row, g.Col - 1}
	}
	if d == "R" {
		return GridPos{g.Row, g.Col + 1}
	}
	log.Fatalln("invalid instruction")
	return g
}

func (g GridPos) chase(h GridPos) GridPos {
	return h
}
