package main

import (
	"aoc/internal/utils"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	// read form file
	input, err := utils.ReadFile("resources/day03.txt")
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
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := re.FindAllStringSubmatch(input, -1)

	totalNumber := 0
	for _, match := range matches {
		number1, _ := strconv.Atoi(match[1])
		number2, _ := strconv.Atoi(match[2])
		totalNumber += number1 * number2
	}

	return totalNumber
}

// part two
func part2(input string) int {
	re := regexp.MustCompile(`(mul\((\d+),(\d+)\)|do\(\)|don't\(\))`)
	matches := re.FindAllStringSubmatch(input, -1)

	isEnabled := true

	totalNumber := 0
	for i := 0; i < len(matches); i++ {
		match := matches[i]
		if match[1] == "don't()" {
			isEnabled = false
			continue
		} else if match[1] == "do()" {
			isEnabled = true
			continue
		} else {
			if !isEnabled {
				continue
			}

			number1, _ := strconv.Atoi(match[2])
			number2, _ := strconv.Atoi(match[3])
			// fmt.Println("First number:", number1)
			// fmt.Println("Second number:", number2)
			totalNumber += number1 * number2
		}
	}

	return totalNumber
}
