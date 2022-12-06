package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// Part 1
	/*for _, input := range exampleInputs() {
		fmt.Println(NewBuffer(input))
	}*/
	//fmt.Println(NewBuffer(fileInput(), 4))

	// Part 2
	/*for _, input := range exampleInputs() {
		fmt.Println(NewBuffer(input, 14))
	}*/
	fmt.Println(NewBuffer(fileInput(), 14))

}

func fileInput() (result string) {
	file, err := os.Open("2022-06-input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = scanner.Text()
	}
	return
}

func countCharsOfFirstUniqueConsecutiveNChars(input string, markerLength int) int {
	for i := markerLength; i < len(input); i++ {
		sbst := input[i-markerLength : i]
		if isUnique(sbst, markerLength) {
			return i
		}
	}
	return len(input)
}

func isUnique(sbst string, markerLength int) bool {
	set := make(map[int32]bool)
	for _, s := range sbst {
		set[s] = true
	}
	return len(set) == markerLength
}

func exampleInputs() []string {
	return []string{
		"mjqjpqmgbljsphdztnvjfqwrcgsmlb",    //7, 19
		"bvwbjplbgvbhsrlpgdmjqwftvncz",      //5, 23
		"nppdvjthqldpwncqszvftbrmjlhg",      //6, 23
		"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", //10, 29
		"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",  //11, 26
	}
}

type Buffer struct {
	Input        string
	MarkerLength int
	Marker       int
}

func (b *Buffer) String() string {
	return fmt.Sprint(b.MarkerLength, " - ", b.Marker, " - ", b.Input[0:8])
}

func NewBuffer(input string, markerLength int) *Buffer {
	return &Buffer{
		Input:        input,
		MarkerLength: markerLength,
		Marker:       countCharsOfFirstUniqueConsecutiveNChars(input, markerLength),
	}
}
