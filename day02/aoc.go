package main

import (
	"fmt"
	"os"
	"strings"
)

type rockPaperScissorCalculator struct {
	points map[string][]int
}

func newRockProckPaperScissorCalculator() *rockPaperScissorCalculator {
	rpsc := rockPaperScissorCalculator{
		points: map[string][]int{
			"A X": {4, 3},
			"B X": {1, 1},
			"C X": {7, 2},

			"A Y": {8, 4},
			"B Y": {5, 5},
			"C Y": {2, 6},

			"A Z": {3, 8},
			"B Z": {9, 9},
			"C Z": {6, 7},
		},
	}

	return &rpsc
}

func (rpsc *rockPaperScissorCalculator) roundScore(round string) []int {
	return rpsc.points[round]
}

func getSolutionPart1(rounds []string) int {
	rpsc := newRockProckPaperScissorCalculator()

	totalScore := 0

	for _, rawRound := range rounds {
		round := strings.Split(rawRound, "\r\n")
		roundScore := rpsc.roundScore(round[0])
		totalScore += roundScore[0]
	}

	return totalScore
}

func getSolutionPart2(rounds []string) int {
	rpsc := newRockProckPaperScissorCalculator()

	totalScore := 0

	for _, rawRound := range rounds {
		round := strings.Split(rawRound, "\r\n")
		roundScore := rpsc.roundScore(round[0])
		totalScore += roundScore[1]
	}

	return totalScore
}

func parseInput(input string) ([]string, error) {
	rounds := strings.Split(strings.TrimSpace(input), "\r\n")

	return rounds, nil
}

func main() {
	inputBytes, err := os.ReadFile("input.txt")
	if err != nil {
		panic("couldn't read input")
	}

	input, err := parseInput(string(inputBytes))
	if err != nil {
		panic("couldn't parse input")
	}

	fmt.Println("Go")
	part := os.Getenv("part")

	if part == "part2" {
		fmt.Println(getSolutionPart2(input))
	} else {
		fmt.Println(getSolutionPart1(input))
	}
}
