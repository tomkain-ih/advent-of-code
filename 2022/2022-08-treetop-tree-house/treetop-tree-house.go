package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	//grid := exampleGrid()
	grid := inputGrid()

	// Part 1
	countVisibleTreesFromEdges(grid)

	// Part 2
	fmt.Println(calculateMaxScenicScore(grid))
	//calculateScenicScore([]int{3, 3, 5, 4, 9}, []int{3, 5, 3, 5, 3}, 3, 2)
	//calculateScenicScore([]int{3, 3, 5, 4, 9}, []int{3, 2, 2, 9, 0}, 3, 4)
	//calculateScenicScore([]int{6, 5, 3, 3, 2}, []int{0, 5, 5, 3, 5}, 2, 1)
}

func calculateMaxScenicScore(grid []string) (maxScenicScore int) {
	_, dimensions := countEdgeTrees(grid)
	rmap, cmap := extractRowAndColMaps(grid)

	for r := 0; r < dimensions[0]; r++ {
		for c := 0; c < dimensions[1]; c++ {
			row := rmap[key("r", r)]
			col := cmap[key("c", c)]

			maxScenicScore = max(maxScenicScore, calculateScenicScore(row, col, r, c))
		}
	}
	return
}

func calculateScenicScore(row []int, col []int, r int, c int) int {
	height := row[c]
	if height != col[r] {
		log.Fatalln("index mismatch")
	}

	var scores [4]int
	onEdge := r == 0 || c == 0 || r == len(row)-1 || c == len(col)-1
	if onEdge {
		scores = [4]int{0, 0, 0, 0}
	} else {
		scores = [4]int{
			countVisibleTreesFromInternal(height, reverse(row[:c])),
			countVisibleTreesFromInternal(height, row[c+1:]),
			countVisibleTreesFromInternal(height, reverse(col[:r])),
			countVisibleTreesFromInternal(height, col[r+1:]),
		}
	}

	score := 1
	for _, s := range scores {
		score = score * s
	}
	/*if !onEdge {
		printScenicScoreDetails(row, col, r, c, height, scores, score)
	}*/
	return score

}

func reverse(ints []int) []int {
	result := make([]int, len(ints))
	copy(result, ints)
	for i, j := 0, len(ints)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = ints[j], ints[i]
	}
	return result
}

func printScenicScoreDetails(row []int, col []int, r int, c int, height int, scores [4]int, score int) {
	fmt.Println(row)
	fmt.Println(col)
	fmt.Print(r)
	fmt.Print("-")
	fmt.Print(c)
	fmt.Print(" = ")
	fmt.Print(height)
	fmt.Print(" : ")
	fmt.Print(scores)
	fmt.Print(" = ")
	fmt.Print(score)
	fmt.Println()

}

func countVisibleTreesFromInternal(height int, ints []int) int {
	result := len(ints)
	//fmt.Printf("Comparing %d and %v\n", height, ints)
	for i, v := range ints {
		if v >= height {
			result = i + 1
			break
		}

	}
	//fmt.Println(result)
	return result

}

func max(i1 int, i2 int) int {
	if i1 > i2 {
		return i1
	}
	return i2
}

func countVisibleTreesFromEdges(grid []string) int {
	edgeTrees, dimensions := countEdgeTrees(grid)

	//capture arrays for each column and row
	rmap, cmap := extractRowAndColMaps(grid)

	//evaluate each position except edges
	interiorPresent := 0
	for r := 1; r < dimensions[0]-1; r++ {
		for c := 1; c < dimensions[1]-1; c++ {
			row := rmap[key("r", r)]
			col := cmap[key("c", c)]

			if isVisible(row, col, r, c) {
				interiorPresent++
			}
		}
	}

	return edgeTrees + interiorPresent
}

func isVisible(row []int, col []int, r int, c int) bool {
	//printVisibilityDetails(row, col, r, c)
	height := row[c]
	if height != col[r] {
		log.Fatalln("index mismatch")
	}

	if checkVisibility(height, row[:c]) {
		return true
	}
	if checkVisibility(height, row[c+1:]) {
		return true
	}
	if checkVisibility(height, col[:r]) {
		return true
	}
	return checkVisibility(height, col[r+1:])
}

func printVisibilityDetails(row []int, col []int, r int, c int) {
	fmt.Println(row)
	fmt.Println(col)
	fmt.Print(r)
	fmt.Print("-")
	fmt.Print(c)
	fmt.Print(" = ")
	fmt.Print(row[c])
	fmt.Print(" (")
	fmt.Print(col[r])
	fmt.Print(")")
	fmt.Println()

	fmt.Println(row[:c])
	fmt.Println(row[c+1:])
	fmt.Println(col[:r])
	fmt.Println(col[r+1:])
}

func checkVisibility(cellValue int, array []int) bool {
	for _, i := range array {
		if i >= cellValue {
			return false
		}
	}
	return true
}

func extractRowAndColMaps(grid []string) (rmap map[string][]int, cmap map[string][]int) {
	rmap = make(map[string][]int)
	cmap = make(map[string][]int)
	for rownum, rowstring := range grid {
		var row []int
		for i, n := range rowstring {
			t, _ := strconv.Atoi(string(n))
			row = append(row, t)

			ckey := key("c", i)
			col := append(cmap[ckey], t)
			cmap[ckey] = col
		}
		rkey := key("r", rownum)
		rmap[rkey] = row
	}
	return
}

func key(s string, i int) string {
	return s + strconv.Itoa(i)
}

func countEdgeTrees(grid []string) (count int, dimensions [2]int) {
	rows := len(grid)
	cols := len(grid[0])
	count = (rows * 2) + ((cols - 2) * 2)
	dimensions = [2]int{rows, cols}
	return
}

func inputGrid() (result []string) {
	file, err := os.Open("2022-08-input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	return
}

func exampleGrid() []string {
	return strings.Split(`30373
25512
65332
33549
35390`, "\n")
}
