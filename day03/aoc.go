package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

type rucksack struct {
	firstCompartment  string
	secondCompartment string
}

func newRucksuck(rawRucksack string) *rucksack {
	r := rucksack{}
	halfRucksack := (len(rawRucksack) / 2)
	r.firstCompartment = rawRucksack[:halfRucksack]
	r.secondCompartment = rawRucksack[halfRucksack:]
	return &r
}

type rucksackOrganizer struct {
	priorities map[rune]int
	rucksacks  []rucksack
}

func newRucksackOrganizer() *rucksackOrganizer {
	ro := rucksackOrganizer{}
	ro.priorities = make(map[rune]int)
	ro.rucksacks = make([]rucksack, 0)

	lowerCounter := 1
	upperCounter := 27
	for r := 'a'; r <= 'z'; r++ {
		ro.priorities[r] = lowerCounter
		lowerCounter++
		ro.priorities[unicode.ToUpper(r)] = upperCounter
		upperCounter++
	}

	return &ro
}

func (ro *rucksackOrganizer) addRucksacks(rucksacks []string) {
	for _, rawRucksack := range rucksacks {
		r := newRucksuck(rawRucksack)
		ro.rucksacks = append(ro.rucksacks, *r)
	}
}

func (ro *rucksackOrganizer) priorityOfItemAppearingInBothCompartmentsOfEachRucksack() []int {
	var priorities []int

	for _, r := range ro.rucksacks {
		for _, firstType := range r.firstCompartment {
			for _, secondType := range r.secondCompartment {
				if firstType == secondType {
					points := ro.priorities[rune(firstType)]
					priorities = append(priorities, points)
					goto nextRucksack
				}
			}
		}
	nextRucksack:
	}

	return priorities
}

func getSolutionPart1(input []string) int {
	ro := newRucksackOrganizer()
	ro.addRucksacks(input)

	sum := 0

	for _, priority := range ro.priorityOfItemAppearingInBothCompartmentsOfEachRucksack() {
		sum += priority
	}

	return sum
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
