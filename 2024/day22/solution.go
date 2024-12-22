package day22

import (
	"log"
	"os"
	"strconv"
	"strings"
)

type Solver struct {
	Input string
	File  string
}

func (d Solver) GetLabel() string {
	return "Day 22"
}

func (d Solver) GetFile() string {
	if d.File != "" {
		return d.File
	}
	return "day22/input.txt"
}

func (d Solver) SolvePart1() string {
	result := 0
	for _, line := range strings.Split(d.GetInput(), "\n") {
		secret, _ := strconv.Atoi(line)
		result += computeSecretNumber(secret, 2000)
	}
	return strconv.Itoa(result)
}

func (d Solver) SolvePart2() string {
	return "Part 2 Solve"
}

func (d Solver) GetInput() string {
	if d.Input != "" {
		return d.Input
	}

	data, err := os.ReadFile(d.GetFile())
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}

func computeSecretNumber(init int, gen int) int {
	result := init
	for i := 0; i < gen; i++ {
		result = evolve(result)
	}
	return result
}

func evolve(secret int) int {
	// Calculate the result of multiplying the secret number by 64.
	// Then, mix this result into the secret number.
	// Finally, prune the secret number.
	mixin := secret * 64
	secret = mix(secret, mixin)
	secret = prune(secret)

	// Calculate the result of dividing the secret number by 32.
	// Round the result down to the nearest integer.
	// Then, mix this result into the secret number.
	// Finally, prune the secret number.
	mixin = secret / 32
	secret = mix(secret, mixin)
	secret = prune(secret)

	// Calculate the result of multiplying the secret number by 2048.
	// Then, mix this result into the secret number.
	// Finally, prune the secret number.
	mixin = secret * 2048
	secret = mix(secret, mixin)
	secret = prune(secret)

	return secret
}

func mix(secret int, mixin int) int {
	// calculate the bitwise XOR of the given value and the secret number.
	// Then, the secret number becomes the result of that operation.
	// (If the secret number is 42 and you were to mix 15 into the secret number,
	// the secret number would become 37.)
	return secret ^ mixin
}

func prune(secret int) int {
	// calculate the value of the secret number modulo 16777216.
	// Then, the secret number becomes the result of that operation.
	// (If the secret number is 100000000 and you were to prune the secret number,
	// the secret number would become 16113920.)
	return secret % 16777216
}
