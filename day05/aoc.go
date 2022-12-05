package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type supplyStacks struct {
	stacksOfCrates         map[int][]rune
	rearrangementProcedure [][]int
}

func newSupplyStacks(stacksOfCrates map[int][]rune, rearrangementProcedure [][]int) *supplyStacks {
	ss := supplyStacks{
		stacksOfCrates,
		rearrangementProcedure,
	}

	return &ss
}

func getSolutionPart1(input *supplyStacks) int {
	return int(input.stacksOfCrates[0][0])
}

func getSolutionPart2(input *supplyStacks) int {
	return int(input.stacksOfCrates[0][0])
}

func parseStacksOfCrates(input string) *map[int][]rune {
	stacksOfCrates := make(map[int][]rune)

	lines := strings.Split(input, "\r\n")
	columns := (len(lines[0]) + 1) / 4

	for _, line := range lines {
		if strings.ContainsAny("1", line) {
			break
		}
		for column := 1; column <= columns; column++ {
			if line[(column*4)-3] != ' ' {
				stacksOfCrates[column] = append(stacksOfCrates[column], (rune)(line[(column*4)-3]))
			}
		}
	}

	return &stacksOfCrates
}

func parseRearrangementProcedure(input string) (*[][]int, error) {
	rearrangementProcedure := [][]int{}
	lines := strings.Split(input, "\r\n")
	re := regexp.MustCompile(`\d+`)

	for _, line := range lines {
		procedureRow := []int{}
		for _, num := range re.FindAllString(line, 3) {
			i, err := strconv.Atoi(num)
			if err != nil {
				return nil, err
			}
			procedureRow = append(procedureRow, i)
		}
		rearrangementProcedure = append(rearrangementProcedure, procedureRow)
	}

	return &rearrangementProcedure, nil
}

func parseInput(input string) (*supplyStacks, error) {
	parts := strings.Split(input, "\r\n\r\n")

	stacksOfCrates := parseStacksOfCrates(parts[0])
	rearrangementProcedure, err := parseRearrangementProcedure(parts[1])
	if err != nil {
		return nil, err
	}

	return newSupplyStacks(*stacksOfCrates, *rearrangementProcedure), nil
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
