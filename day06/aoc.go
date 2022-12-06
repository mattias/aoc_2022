package main

import (
	"fmt"
	"os"
)

func isAllDifferent(input string) bool {
	checked := make(map[rune]bool)

	for _, char := range input {
		if checked[char] {
			return false
		}

		checked[char] = true
	}

	return true
}

func getSolutionPart1And2(input string, amountOfDistinctCharacters int) int {
	for i := 0; i < len(input); i++ {
		if isAllDifferent(input[i : i+amountOfDistinctCharacters]) {
			return i + amountOfDistinctCharacters
		}
	}

	return 0
}

func main() {
	inputBytes, err := os.ReadFile("input.txt")
	if err != nil {
		panic("couldn't read input")
	}

	input := string(inputBytes)

	fmt.Println("Go")
	part := os.Getenv("part")

	if part == "part2" {
		fmt.Println(getSolutionPart1And2(input, 14))
	} else {
		fmt.Println(getSolutionPart1And2(input, 4))
	}
}
