package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func calculateCaloriesPerElf(input [][]int) *[]int {
	caloriesPerElf := make([]int, 0)
	var sum int

	for _, calorieGroup := range input {
		sum = 0
		for _, calorie := range calorieGroup {
			sum += calorie
		}

		caloriesPerElf = append(caloriesPerElf, sum)
	}

	return &caloriesPerElf
}

func sortCaloriesByDescending(caloriesPerElf *[]int) {
	sort.Sort(sort.Reverse(sort.IntSlice(*caloriesPerElf)))
}

func top3Highest(caloriesPerElf *[]int) int {
	return (*caloriesPerElf)[0] + (*caloriesPerElf)[1] + (*caloriesPerElf)[2]
}

func getSolutionPart1(input [][]int) int {
	caloriesPerElf := calculateCaloriesPerElf(input)
	sortCaloriesByDescending(caloriesPerElf)

	return (*caloriesPerElf)[0]
}

func getSolutionPart2(input [][]int) int {
	caloriesPerElf := calculateCaloriesPerElf(input)
	sortCaloriesByDescending(caloriesPerElf)

	return top3Highest(caloriesPerElf)
}

func parseInput(input string) ([][]int, error) {
	var ints [][]int

	lines := strings.Split(strings.TrimSpace(input), "\r\n")
	j := 0
	temp := make([]int, 0)

	for _, line := range lines {
		if line == "" {
			j++
			ints = append(ints, [][]int{temp}...)
			temp = make([]int, 0)
			continue
		}
		i, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		temp = append(temp, i)
	}

	ints = append(ints, [][]int{temp}...)

	return ints, nil
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
