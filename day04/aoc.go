package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getSolutionPart1(sectionAssignmentsForEachPair []string) int {
	containsFullyWithinAnother := 0

	for _, sectionAssignmentsForPair := range sectionAssignmentsForEachPair {
		sectionAssignments := strings.Split(sectionAssignmentsForPair, ",")
		elf1Sections := strings.Split(sectionAssignments[0], "-")
		elf2Sections := strings.Split(sectionAssignments[1], "-")
		elf1FirstSection, _ := strconv.Atoi(elf1Sections[0])
		elf1SecondSection, _ := strconv.Atoi(elf1Sections[1])
		elf2FirstSection, _ := strconv.Atoi(elf2Sections[0])
		elf2SecondSection, _ := strconv.Atoi(elf2Sections[1])

		if (elf1FirstSection <= elf2FirstSection && elf1SecondSection >= elf2SecondSection) || (elf2FirstSection <= elf1FirstSection && elf2SecondSection >= elf1SecondSection) {
			containsFullyWithinAnother++
		}
	}

	return containsFullyWithinAnother
}

func getSolutionPart2(sectionAssignmentsForEachPair []string) int {
	containsWithinAnother := 0

	for _, sectionAssignmentsForPair := range sectionAssignmentsForEachPair {
		sectionAssignments := strings.Split(sectionAssignmentsForPair, ",")
		elf1Sections := strings.Split(sectionAssignments[0], "-")
		elf2Sections := strings.Split(sectionAssignments[1], "-")
		elf1FirstSection, _ := strconv.Atoi(elf1Sections[0])
		elf1SecondSection, _ := strconv.Atoi(elf1Sections[1])
		elf2FirstSection, _ := strconv.Atoi(elf2Sections[0])
		elf2SecondSection, _ := strconv.Atoi(elf2Sections[1])

		if (elf1FirstSection <= elf2FirstSection && elf1SecondSection >= elf2FirstSection) || (elf2FirstSection <= elf1FirstSection && elf2SecondSection >= elf1FirstSection) {
			containsWithinAnother++
		}
	}

	return containsWithinAnother
}

func parseInput(input string) []string {
	return strings.Split(strings.TrimSpace(input), "\r\n")
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
