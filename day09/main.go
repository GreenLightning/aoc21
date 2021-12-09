package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	basins := make([][]int, len(grid))
	for y, row := range grid {
		basins[y] = make([]int, len(row))
	}

	var count, sum int
	for y := range grid {
		for x := range grid[y] {
			if x-1 >= 0 && grid[y][x-1] <= grid[y][x] {
				continue
			}
			if x+1 < len(grid[y]) && grid[y][x+1] <= grid[y][x] {
				continue
			}
			if y-1 >= 0 && grid[y-1][x] <= grid[y][x] {
				continue
			}
			if y+1 < len(grid) && grid[y+1][x] <= grid[y][x] {
				continue
			}
			count++
			basins[y][x] = count
			sum += 1 + grid[y][x]
		}
	}

	fmt.Println("--- Part One ---")
	fmt.Println(sum)

	for {
		changed := false
		for y := range basins {
			for x := range basins[y] {
				if basins[y][x] == 0 {
					continue
				}
				if x-1 >= 0 && basins[y][x-1] == 0 && grid[y][x-1] != 9 {
					basins[y][x-1] = basins[y][x]
					changed = true
				}
				if x+1 < len(basins[y]) && basins[y][x+1] == 0 && grid[y][x+1] != 9 {
					basins[y][x+1] = basins[y][x]
					changed = true
				}
				if y-1 >= 0 && basins[y-1][x] == 0 && grid[y-1][x] != 9 {
					basins[y-1][x] = basins[y][x]
					changed = true
				}
				if y+1 < len(basins) && basins[y+1][x] == 0 && grid[y+1][x] != 9 {
					basins[y+1][x] = basins[y][x]
					changed = true
				}
			}
		}
		if !changed {
			break
		}
	}

	sizes := make([]int, count+1)
	for y := range basins {
		for x := range basins[y] {
			sizes[basins[y][x]]++
		}
	}

	// Remove the first entry, which counts the positions
	// not assigned to a basin (i.e. 9s).
	sizes = sizes[1:]

	sort.Ints(sizes)

	fmt.Println("--- Part Two ---")
	fmt.Println(sizes[len(sizes)-1] * sizes[len(sizes)-2] * sizes[len(sizes)-3])
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
