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
	if g.touches(h) {
		return g
	}
	// diagonal
	if g.Row != h.Row && g.Col != h.Col {
		return moveDiagonal(g, h)
	}
	//stepwise
	return moveStepwise(g, h)
}

func moveDiagonal(g GridPos, h GridPos) GridPos {
	rowDiff := g.Row - h.Row
	colDiff := g.Col - h.Col
	//absolute value hack
	rowDiff = rowDiff * rowDiff
	colDiff = colDiff * colDiff

	if colDiff > rowDiff {
		//move to adjacent col, same row
		newCol := h.Col
		if g.Col > h.Col {
			newCol++
		} else {
			newCol--
		}
		return GridPos{h.Row, newCol}
	}
	newRow := h.Row
	if g.Row > h.Row {
		newRow++
	} else {
		newRow--
	}
	return GridPos{newRow, h.Col}

}

func moveStepwise(g GridPos, h GridPos) GridPos {
	if g.Row == h.Row {
		newCol := h.Col
		if g.Col > h.Col {
			newCol++
		} else {
			newCol--
		}
		return GridPos{h.Row, newCol}
	}
	newRow := h.Row
	if g.Row > h.Row {
		newRow++
	} else {
		newRow--
	}
	return GridPos{newRow, h.Col}
}

func (g GridPos) touches(h GridPos) bool {
	if g == h {
		return true
	}
	rowDiff := g.Row - h.Row
	colDiff := g.Col - h.Col
	// absolute value hack
	rowDiff = rowDiff * rowDiff
	colDiff = colDiff * colDiff

	if g.Row == h.Row && colDiff == 1 {
		return true
	}
	if g.Col == h.Col && rowDiff == 1 {
		return true
	}
	if rowDiff == 1 && rowDiff == colDiff {
		return true
	}
	return false
}
