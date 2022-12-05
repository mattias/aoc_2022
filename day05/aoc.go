package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type supplyStacks struct {
	stackColumns           int
	stacksOfCrates         map[int]string
	rearrangementProcedure [][]int
}

func newSupplyStacks(stacksOfCrates map[int]string, rearrangementProcedure [][]int) *supplyStacks {
	ss := supplyStacks{}
	ss.stacksOfCrates = stacksOfCrates
	ss.rearrangementProcedure = rearrangementProcedure
	ss.stackColumns = 0

	return &ss
}

func reverseString(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}

	return
}

func (ss *supplyStacks) runRearrangementProcedure() {
	var sliceSize int
	for _, procedure := range ss.rearrangementProcedure {
		sliceSize = procedure[0]
		from := reverseString(ss.stacksOfCrates[procedure[1]][:sliceSize])
		ss.stacksOfCrates[procedure[1]] = ss.stacksOfCrates[procedure[1]][sliceSize:]
		ss.stacksOfCrates[procedure[2]] = from + ss.stacksOfCrates[procedure[2]]
	}
}

func (ss *supplyStacks) runRearrangementProcedureV2() {
	var sliceSize int
	for _, procedure := range ss.rearrangementProcedure {
		sliceSize = procedure[0]
		from := ss.stacksOfCrates[procedure[1]][:sliceSize]
		ss.stacksOfCrates[procedure[1]] = ss.stacksOfCrates[procedure[1]][sliceSize:]
		ss.stacksOfCrates[procedure[2]] = from + ss.stacksOfCrates[procedure[2]]
	}
}

func (ss *supplyStacks) cratesAtTop() string {
	cratesAtTop := ""
	for i := 1; i <= ss.stackColumns; i++ {
		cratesAtTop += (string)(ss.stacksOfCrates[i][0])
	}
	return cratesAtTop
}

func getSolutionPart1(input *supplyStacks) string {
	input.runRearrangementProcedure()
	return input.cratesAtTop()
}

func getSolutionPart2(input *supplyStacks) string {
	input.runRearrangementProcedureV2()
	return input.cratesAtTop()
}

func parseStacksOfCrates(input string) (*map[int]string, int) {
	stacksOfCrates := make(map[int]string)

	lines := strings.Split(input, "\r\n")
	columns := (len(lines[0]) + 1) / 4

	for _, line := range lines {
		if strings.ContainsAny("1", line) {
			break
		}
		for column := 1; column <= columns; column++ {
			if line[(column*4)-3] != ' ' {
				stacksOfCrates[column] = stacksOfCrates[column] + (string)(line[(column*4)-3])
			}
		}
	}

	return &stacksOfCrates, columns
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

	stacksOfCrates, columns := parseStacksOfCrates(parts[0])
	rearrangementProcedure, err := parseRearrangementProcedure(parts[1])
	if err != nil {
		return nil, err
	}

	ss := newSupplyStacks(*stacksOfCrates, *rearrangementProcedure)
	ss.stackColumns = columns

	return ss, nil
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
