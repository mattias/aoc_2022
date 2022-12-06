package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAOC_getSolutionPart1(t *testing.T) {
	input := "mjqjpqmgbljsphdztnvjfqwrcgsmlb"
	expectedSolution := 7

	actualSolution := getSolutionPart1And2(input, 4)
	assert.Equal(t, expectedSolution, actualSolution)
}

func TestAOC_getSolutionPart2(t *testing.T) {
	input := "mjqjpqmgbljsphdztnvjfqwrcgsmlb"
	expectedSolution := 19

	actualSolution := getSolutionPart1And2(input, 14)
	assert.Equal(t, expectedSolution, actualSolution)
}
