package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Vector2 struct {
	x, y int
}

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
		for x := range costs[y] {
			costs[y][x] = math.MaxInt / 2
		}
	}

	h, w := len(grid), len(grid[0])
	costs[h-1][w-1] = grid[h-1][w-1]
	queue := []Vector2{Vector2{w - 2, h - 1}, Vector2{w - 1, h - 2}}

	for len(queue) != 0 {
		p := queue[0]
		queue = queue[1:]
		x, y := p.x, p.y

		if !(x >= 0 && x < w && y >= 0 && y < h) {
			continue
		}

		if x-1 >= 0 && grid[y][x]+costs[y][x-1] < costs[y][x] {
			costs[y][x] = grid[y][x] + costs[y][x-1]
			queue = append(queue, Vector2{x + 1, y}, Vector2{x, y - 1}, Vector2{x, y + 1})
		}
		if x+1 < w && grid[y][x]+costs[y][x+1] < costs[y][x] {
			costs[y][x] = grid[y][x] + costs[y][x+1]
			queue = append(queue, Vector2{x - 1, y}, Vector2{x, y - 1}, Vector2{x, y + 1})
		}
		if y-1 >= 0 && grid[y][x]+costs[y-1][x] < costs[y][x] {
			costs[y][x] = grid[y][x] + costs[y-1][x]
			queue = append(queue, Vector2{x - 1, y}, Vector2{x + 1, y}, Vector2{x, y + 1})
		}
		if y+1 < h && grid[y][x]+costs[y+1][x] < costs[y][x] {
			costs[y][x] = grid[y][x] + costs[y+1][x]
			queue = append(queue, Vector2{x - 1, y}, Vector2{x + 1, y}, Vector2{x, y - 1})
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
