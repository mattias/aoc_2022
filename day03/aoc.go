package main

import (
	"fmt"
	"os"
	"strings"
)

func getSolutionPart1(input []string) int {
	return len(input)
}

func getSolutionPart2(input []string) int {
	return len(input)
}

func parseInput(input string) []string {
	var rucksacks []string

	lines := strings.Split(strings.TrimSpace(input), "\r\n")
	rucksacks = append(rucksacks, lines...)

	return rucksacks
}

func main() {
	inputBytes, err := os.ReadFile("input.txt")
	if err != nil {
		panic("couldn't read input")
	}

	input := parseInput(string(inputBytes))

	fmt.Println("Go")
	part := os.Getenv("part")

	if part == "part2" {
		fmt.Println(getSolutionPart2(input))
	} else {
		fmt.Println(getSolutionPart1(input))
	}
}
