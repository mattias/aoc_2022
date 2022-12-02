package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAOC_parseInput(t *testing.T) {
	input := "A Y\r\nB X\r\nC Z"
	expectedParsedInput := []string{"A Y", "B X", "C Z"}

	actualParsedInput, err := parseInput(input)
	assert.NoError(t, err)
	assert.Equal(t, expectedParsedInput, actualParsedInput)
}

func TestAOC_getSolutionPart1(t *testing.T) {
	input := []string{"A Y", "B X", "C Z"}
	expectedSolution := 15

	actualSolution := getSolutionPart1(input)
	assert.Equal(t, expectedSolution, actualSolution)
}

func TestAOC_getSolutionPart2(t *testing.T) {
	input := []string{"A Y", "B X", "C Z"}
	expectedSolution := 12

	actualSolution := getSolutionPart2(input)
	assert.Equal(t, expectedSolution, actualSolution)
}
