package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	lines := readLines("input.txt")

	grid := make([][]int, len(lines))
	for y, line := range lines {
		grid[y] = make([]int, len(line))
		for x, rune := range line {
			grid[y][x] = int(rune - '0')
		}
	}

	flashed := make([][]bool, len(grid))
	for y, row := range grid {
		flashed[y] = make([]bool, len(row))
	}

	var flashes int
	for step := 1; ; step++ {
		for y := range grid {
			for x := range grid[y] {
				grid[y][x]++
			}
		}

		for {
			changed := false
			for y := range grid {
				for x := range grid[y] {
					if !flashed[y][x] && grid[y][x] > 9 {
						changed = true
						flashed[y][x] = true
						for dy := -1; dy <= 1; dy++ {
							for dx := -1; dx <= 1; dx++ {
								if y+dy >= 0 && y+dy < len(grid) && x+dx >= 0 && x+dx < len(grid[y+dy]) {
									grid[y+dy][x+dx]++
								}
							}
						}
					}
				}
			}
			if !changed {
				break
			}
		}

		off := false
		for y := range grid {
			for x := range grid[y] {
				if flashed[y][x] {
					flashes++
					flashed[y][x] = false
					grid[y][x] = 0
				} else {
					off = true
				}
			}
		}

		if step == 100 {
			fmt.Println("--- Part One ---")
			fmt.Println(flashes)
		}

		if !off {
			fmt.Println("--- Part Two ---")
			fmt.Println(step)
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
