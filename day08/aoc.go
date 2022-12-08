package main

import (
	"fmt"
	"os"
	"strings"
)

func isTreeVisibleFromTop(current int, trees *[][]int, startRow int, startCol int) bool {
	for row := startRow - 1; row >= 0; row-- {
		check := (*trees)[row][startCol]
		if current <= check {
			return false
		}
	}

	return true
}

func isTreeVisibleFromRight(current int, trees *[][]int, startRow int, startCol int) bool {
	cols := len((*trees)[0]) - 1

	for col := startCol + 1; col <= cols; col++ {
		check := (*trees)[startRow][col]
		if current <= check {
			return false
		}
	}

	return true
}

func isTreeVisibleFromDown(current int, trees *[][]int, startRow int, startCol int) bool {
	rows := len(*trees) - 1

	for row := startRow + 1; row <= rows; row++ {
		check := (*trees)[row][startCol]
		if current <= check {
			return false
		}
	}

	return true
}

func isTreeVisibleFromLeft(current int, trees *[][]int, startRow int, startCol int) bool {
	for col := startCol - 1; col >= 0; col-- {
		check := (*trees)[startRow][col]
		if current <= check {
			return false
		}
	}

	return true
}

func isTreeVisible(current int, trees *[][]int, startRow int, startCol int) bool {
	return isTreeVisibleFromTop(current, trees, startRow, startCol) ||
		isTreeVisibleFromRight(current, trees, startRow, startCol) ||
		isTreeVisibleFromDown(current, trees, startRow, startCol) ||
		isTreeVisibleFromLeft(current, trees, startRow, startCol)
}

func getSolutionPart1(tallTrees [][]int) int {
	rows := len(tallTrees)
	cols := len(tallTrees[0])
	edgeTrees := (cols * 2) + (rows * 2) - 4

	visibleTreeCount := 0

	for row := 1; row <= rows-2; row++ {
		for col := 1; col <= cols-2; col++ {
			currentTree := tallTrees[row][col]
			if isTreeVisible(currentTree, &tallTrees, row, col) {
				visibleTreeCount++
			}
		}
	}

	return edgeTrees + visibleTreeCount
}

func scenicScoreTop(current int, trees *[][]int, startRow int, startCol int) int {
	treesToAdmire := 0
	for row := startRow - 1; row >= 0; row-- {
		check := (*trees)[row][startCol]
		treesToAdmire++
		if current <= check {
			return treesToAdmire
		}
	}

	return treesToAdmire
}

func scenicScoreRight(current int, trees *[][]int, startRow int, startCol int) int {
	cols := len((*trees)[0]) - 1
	treesToAdmire := 0

	for col := startCol + 1; col <= cols; col++ {
		check := (*trees)[startRow][col]
		treesToAdmire++
		if current <= check {
			return treesToAdmire
		}
	}

	return treesToAdmire
}

func scenicScoreDown(current int, trees *[][]int, startRow int, startCol int) int {
	rows := len(*trees) - 1
	treesToAdmire := 0

	for row := startRow + 1; row <= rows; row++ {
		check := (*trees)[row][startCol]
		treesToAdmire++
		if current <= check {
			return treesToAdmire
		}
	}

	return treesToAdmire
}

func scenicScoreLeft(current int, trees *[][]int, startRow int, startCol int) int {
	treesToAdmire := 0

	for col := startCol - 1; col >= 0; col-- {
		check := (*trees)[startRow][col]
		treesToAdmire++
		if current <= check {
			return treesToAdmire
		}
	}

	return treesToAdmire
}

func getSolutionPart2(trees [][]int) int {
	rows := len(trees)
	cols := len(trees[0])

	bestScenicScore := 0

	var current int
	var topScore int
	var rightScore int
	var downScore int
	var leftScore int

	for row := 1; row <= rows-2; row++ {
		for col := 1; col <= cols-2; col++ {
			current = trees[row][col]
			topScore = scenicScoreTop(current, &trees, row, col)
			rightScore = scenicScoreRight(current, &trees, row, col)
			downScore = scenicScoreDown(current, &trees, row, col)
			leftScore = scenicScoreLeft(current, &trees, row, col)

			scenicScore := topScore * rightScore * downScore * leftScore

			if scenicScore > bestScenicScore {
				bestScenicScore = scenicScore
			}
		}
	}

	return bestScenicScore
}

func parseInput(input string) [][]int {
	var ints [][]int

	lines := strings.Split(strings.TrimSpace(input), "\r\n")

	for _, line := range lines {
		var intRow []int
		for _, char := range line {
			i := int(char) - '0'

			intRow = append(intRow, i)
		}

		ints = append(ints, intRow)
	}

	return ints
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
