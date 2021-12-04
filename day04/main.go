package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Board struct {
	numbers [][]int
	markers [][]bool
}

func (board *Board) HasWon() bool {
rows:
	for i := 0; i < len(board.markers); i++ {
		for j := 0; j < len(board.markers[i]); j++ {
			if !board.markers[i][j] {
				continue rows
			}
		}
		return true
	}

columns:
	for j := 0; j < len(board.markers[0]); j++ {
		for i := 0; i < len(board.markers); i++ {
			if !board.markers[i][j] {
				continue columns
			}
		}
		return true
	}

	return false
}

func (board *Board) Unmarked() (sum int) {
	for i, row := range board.numbers {
		for j, number := range row {
			if !board.markers[i][j] {
				sum += number
			}
		}
	}
	return
}

func main() {
	lines := readLines("input.txt")

	numbers := arrayToInt(strings.Split(lines[0], ","))

	var boards []*Board
	var current *Board
	for i := 1; i < len(lines); i++ {
		line := lines[i]
		line = strings.TrimSpace(line)

		if line == "" {
			current = nil
			continue
		}

		if current == nil {
			current = new(Board)
			boards = append(boards, current)
		}

		numbers := arrayToInt(regexp.MustCompile(` +`).Split(line, -1))
		current.numbers = append(current.numbers, numbers)
		markers := make([]bool, len(numbers))
		current.markers = append(current.markers, markers)
	}

	first := true
	for _, draw := range numbers {
		for boardIndex := 0; boardIndex < len(boards); boardIndex++ {
			board := boards[boardIndex]

		search:
			for i, row := range board.numbers {
				for j, number := range row {
					if number == draw {
						board.markers[i][j] = true
						break search
					}
				}
			}

			if board.HasWon() {
				if first {
					first = false
					fmt.Println("--- Part One ---")
					fmt.Println(board.Unmarked() * draw)
				}
				if len(boards) == 1 {
					fmt.Println("--- Part Two ---")
					fmt.Println(board.Unmarked() * draw)
				}

				copy(boards[boardIndex:], boards[boardIndex+1:])
				boards = boards[:len(boards)-1]
				boardIndex--
			}
		}
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

func arrayToInt(input []string) (output []int) {
	output = make([]int, len(input))
	for i, text := range input {
		output[i] = toInt(text)
	}
	return output
}

func toInt(s string) int {
	result, err := strconv.Atoi(s)
	check(err)
	return result
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
