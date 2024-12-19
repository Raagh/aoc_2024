package main

import (
	"aoc/internal/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// read form file
	input, err := utils.ReadFile("resources/day02.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	fmt.Println("--- Part One ---")
	fmt.Println("Result:", part1(input))
	fmt.Println("--- Part Two ---")
	fmt.Println("Result:", part2(input))

	os.Exit(0)
}

// part one
func part1(input string) int {
	totalSafeReports := 0
	line := strings.Split(input, "\n")
	for _, l := range line {
		numbersStrings := strings.Split(l, " ")
		if isSafe(numbersStrings) {
			totalSafeReports++
		}
	}

	return totalSafeReports
}

func isSafe(numbersStrings []string) bool {
	if len(numbersStrings) < 2 {
		return false
	}

	firstNumber, _ := strconv.Atoi(numbersStrings[0])
	secondNumber, _ := strconv.Atoi(numbersStrings[1])
	shouldDecrease := firstNumber > secondNumber

	for i := 1; i < len(numbersStrings); i++ {
		firstNumber, _ := strconv.Atoi(numbersStrings[i-1])
		secondNumber, _ := strconv.Atoi(numbersStrings[i])
		isDecreasing := firstNumber > secondNumber

		if isDecreasing != shouldDecrease {
			return false
		}

		if abs(firstNumber-secondNumber) < 1 || abs(firstNumber-secondNumber) > 3 {
			return false
		}
	}

	return true
}

func isSafeByRemoving(list []string) bool {
	for i := 0; i < len(list); i++ {
		newArray := append([]string{}, list[:i]...)
		newArray = append(newArray, list[i+1:]...)

		if isSafe(newArray) {
			return true
		}
	}

	return false
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// part two
func part2(input string) int {
	totalSafeReports := 0
	line := strings.Split(input, "\n")
	for _, l := range line {
		numbersStrings := strings.Split(l, " ")
		if isSafe(numbersStrings) || isSafeByRemoving(numbersStrings) {
			totalSafeReports++
		}
	}

	return totalSafeReports
}
