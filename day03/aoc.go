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
	rawRucksack       string
}

func newRucksuck(rawRucksack string) *rucksack {
	r := rucksack{}
	r.rawRucksack = rawRucksack
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
		priorities = append(priorities, ro.findPriorityPoint(r))
	}

	return priorities
}

func (ro *rucksackOrganizer) findPriorityPoint(r rucksack) int {
	for _, firstType := range r.firstCompartment {
		for _, secondType := range r.secondCompartment {
			if firstType == secondType {
				points := ro.priorities[rune(firstType)]
				return points
			}
		}
	}

	return 0
}

func (ro *rucksackOrganizer) priorityPointsOfBadgesFromEachGroup() []int {
	var priorities []int

	for i := 0; i < len(ro.rucksacks); i += 3 {

		priorities = append(priorities, ro.findPriorityPointForBadge(
			ro.rucksacks[i],
			ro.rucksacks[i+1],
			ro.rucksacks[i+2],
		))
	}

	return priorities
}

func (ro *rucksackOrganizer) findPriorityPointForBadge(r1 rucksack, r2 rucksack, r3 rucksack) int {
	for _, firstType := range r1.rawRucksack {
		for _, secondType := range r2.rawRucksack {
			for _, thirdType := range r3.rawRucksack {
				if firstType == secondType && firstType == thirdType {
					points := ro.priorities[rune(firstType)]
					return points
				}
			}
		}
	}

	return 0
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
	ro := newRucksackOrganizer()
	ro.addRucksacks(input)

	sum := 0

	for _, priority := range ro.priorityPointsOfBadgesFromEachGroup() {
		sum += priority
	}

	return sum
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
