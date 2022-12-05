package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAOC_parseInput(t *testing.T) {
	input := "    [D]    \r\n[N] [C]    \r\n[Z] [M] [P]\r\n 1   2   3 \r\n\r\nmove 1 from 2 to 1\r\nmove 3 from 1 to 3\r\nmove 2 from 2 to 1\r\nmove 1 from 1 to 2"

	expectedStacksOfCrates := map[int]string{
		1: "NZ",
		2: "DCM",
		3: "P",
	}

	expectedRearrangementProcedure := [][]int{
		{1, 2, 1},
		{3, 1, 3},
		{2, 2, 1},
		{1, 1, 2},
	}

	expectedSupplyStacks := newSupplyStacks(expectedStacksOfCrates, expectedRearrangementProcedure)

	actualParsedInput, err := parseInput(input)
	assert.NoError(t, err)
	assert.Equal(t, expectedSupplyStacks, actualParsedInput)
}

func TestAOC_getSolutionPart1(t *testing.T) {
	stacksOfCrates := map[int]string{
		1: "NZ",
		2: "DCM",
		3: "P",
	}

	rearrangementProcedure := [][]int{
		{1, 2, 1},
		{3, 1, 3},
		{2, 2, 1},
		{1, 1, 2},
	}

	input := newSupplyStacks(stacksOfCrates, rearrangementProcedure)
	input.stackColumns = 3
	expectedSolution := "CMZ"

	actualSolution := getSolutionPart1(input)
	assert.Equal(t, expectedSolution, actualSolution)
}

func TestAOC_getSolutionPart2(t *testing.T) {
	stacksOfCrates := map[int]string{
		1: "NZ",
		2: "DCM",
		3: "P",
	}

	rearrangementProcedure := [][]int{
		{1, 2, 1},
		{3, 1, 3},
		{2, 2, 1},
		{1, 1, 2},
	}

	input := newSupplyStacks(stacksOfCrates, rearrangementProcedure)
	input.stackColumns = 3
	expectedSolution := "MCD"

	actualSolution := getSolutionPart2(input)
	assert.Equal(t, expectedSolution, actualSolution)
}
