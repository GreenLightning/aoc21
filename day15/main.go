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

	{
		fmt.Println("--- Part One ---")
		fmt.Println(findShortestPath(grid))
	}

	size := len(grid)

	newGrid := make([][]int, 5*size)
	for y := range newGrid {
		newGrid[y] = make([]int, 5*size)
	}

	for y := 0; y < 5*size; y++ {
		for x := 0; x < 5*size; x++ {
			value := grid[y%size][x%size] + y/size + x/size
			if value > 9 {
				value -= 9
			}
			newGrid[y][x] = value
		}
	}

	{
		fmt.Println("--- Part Two ---")
		fmt.Println(findShortestPath(newGrid))
	}
}

func findShortestPath(grid [][]int) int {
	costs := make([][]int, len(grid))
	for y, row := range grid {
		costs[y] = make([]int, len(row))
	}

	h, w := len(grid), len(grid[0])

	costs[h-1][w-1] = grid[h-1][w-1]

	for d := h - 2; d >= 0; d-- {
		for i := 0; d+i < w; i++ {
			x, y := d+i, h-1-i
			if x == w-1 {
				costs[y][x] = grid[y][x] + costs[y+1][x]
			} else if y == h-1 {
				costs[y][x] = grid[y][x] + costs[y][x+1]
			} else {
				costs[y][x] = grid[y][x] + min(costs[y+1][x], costs[y][x+1])
			}
		}
	}

	for d := h - 2; d >= 0; d-- {
		for i := 0; d-i >= 0; i++ {
			x, y := i, d-i
			if x == w-1 {
				costs[y][x] = grid[y][x] + costs[y+1][x]
			} else if y == h-1 {
				costs[y][x] = grid[y][x] + costs[y][x+1]
			} else {
				costs[y][x] = grid[y][x] + min(costs[y+1][x], costs[y][x+1])
			}
		}
	}

	for {
		changed := false
		for y := 0; y < h; y++ {
			for x := 0; x < w; x++ {
				if x-1 >= 0 && grid[y][x]+costs[y][x-1] < costs[y][x] {
					costs[y][x] = grid[y][x] + costs[y][x-1]
					changed = true
				}
				if x+1 < w && grid[y][x]+costs[y][x+1] < costs[y][x] {
					costs[y][x] = grid[y][x] + costs[y][x+1]
					changed = true
				}
				if y-1 >= 0 && grid[y][x]+costs[y-1][x] < costs[y][x] {
					costs[y][x] = grid[y][x] + costs[y-1][x]
					changed = true
				}
				if y+1 < h && grid[y][x]+costs[y+1][x] < costs[y][x] {
					costs[y][x] = grid[y][x] + costs[y+1][x]
					changed = true
				}
			}
		}
		if !changed {
			break
		}
	}

	return costs[0][0] - grid[0][0]
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

func min(x, y int) int {
	if y < x {
		return y
	}
	return x
}
