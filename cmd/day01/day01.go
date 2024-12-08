package main

import (
	"aoc/internal/utils"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input, err := utils.ReadFile("resources/day01.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	leftList := []int{}
	rightList := []int{}
	// parse input into the two lists
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		numbersStrings := strings.Split(line, "   ")

		if len(numbersStrings) < 2 {
			break
		}

		numLeft, err := strconv.Atoi(numbersStrings[0])
		numRight, err := strconv.Atoi(numbersStrings[1])
		if err != nil {
			// handle error
		}

		leftList = append(leftList, numLeft)
		rightList = append(rightList, numRight)
	}

	sort.Ints(leftList)
	sort.Ints(rightList)

	fmt.Println("--- Part One ---")
	fmt.Println("Result:", part1(leftList, rightList))
	fmt.Println("--- Part Two ---")
	fmt.Println("Result:", part2(leftList, rightList))

	os.Exit(0)
}

// part one
func part1(leftList []int, rightList []int) int {
	// calculate distances bwtween the two lists
	distances := []int{}
	for i := 0; i < len(leftList); i++ {
		if leftList[i] > rightList[i] {
			distances = append(distances, leftList[i]-rightList[i])
		} else {
			distances = append(distances, rightList[i]-leftList[i])
		}
	}

	return sum(distances)
}

// part two
func part2(leftList []int, rightList []int) int {
	similarityScores := []int{}
	for _, v := range leftList {
		count := 0
		for _, w := range rightList {
			if v == w {
				count++
			} else if w > v {
				break
			}
		}

		similarityScores = append(similarityScores, v*count)
	}

	return sum(similarityScores)
}

func sum(elements []int) int {
	sum := 0
	for _, element := range elements {
		sum += element
	}

	return sum
}
