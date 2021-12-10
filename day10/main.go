package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	lines := readLines("input.txt")

	matching := map[rune]rune{'(': ')', '[': ']', '{': '}', '<': '>'}
	syntaxErrorPoints := map[rune]int{')': 3, ']': 57, '}': 1197, '>': 25137}
	autoCompletePoints := map[rune]int{')': 1, ']': 2, '}': 3, '>': 4}

	var syntaxErrorScore int
	var autoCompleteScores []int

lineLoop:
	for _, line := range lines {

		var stack []rune
		for _, rune := range line {
			switch rune {
			case '(', '[', '{', '<':
				stack = append(stack, matching[rune])
			case ')', ']', '}', '>':
				// Does not handle extra closing tokens.
				top := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if rune != top {
					syntaxErrorScore += syntaxErrorPoints[rune]
					continue lineLoop
				}
			}
		}

		// Line is incomplete.
		var score int
		for i := len(stack) - 1; i >= 0; i-- {
			score = 5*score + autoCompletePoints[stack[i]]
		}
		autoCompleteScores = append(autoCompleteScores, score)
	}

	{
		fmt.Println("--- Part One ---")
		fmt.Println(syntaxErrorScore)
	}

	{
		fmt.Println("--- Part Two ---")
		sort.Ints(autoCompleteScores)
		fmt.Println(autoCompleteScores[len(autoCompleteScores)/2])
	}
}

func readLines(filename string) []string {
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
