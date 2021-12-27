package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	lines := readLines("input.txt")

	grid := make([][]rune, len(lines))
	temp := make([][]rune, len(lines))
	for y, line := range lines {
		grid[y] = []rune(line)
		temp[y] = []rune(line)
	}

	steps := 0
	for {
		changed := false

		for y := 0; y < len(grid); y++ {
			copy(temp[y], grid[y])
		}
		for y := 0; y < len(grid); y++ {
			for x := 0; x < len(grid[y]); x++ {
				nx := x + 1
				if nx == len(grid[y]) {
					nx = 0
				}
				if grid[y][x] == '>' && grid[y][nx] == '.' {
					temp[y][x] = '.'
					temp[y][nx] = '>'
					changed = true
				}
			}
		}
		grid, temp = temp, grid

		for y := 0; y < len(grid); y++ {
			copy(temp[y], grid[y])
		}
		for y := 0; y < len(grid); y++ {
			for x := 0; x < len(grid[y]); x++ {
				ny := y + 1
				if ny == len(grid) {
					ny = 0
				}
				if grid[y][x] == 'v' && grid[ny][x] == '.' {
					temp[y][x] = '.'
					temp[ny][x] = 'v'
					changed = true
				}
			}
		}
		grid, temp = temp, grid

		steps++

		if !changed {
			fmt.Println("--- Part One ---")
			fmt.Println(steps)
			break
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

func check(err error) {
	if err != nil {
		panic(err)
	}
}
