package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Vector2 struct {
	x, y int
}

func main() {
	lines := readLines("input.txt")
	lineRegex := regexp.MustCompile(`^(\d+),(\d+) -> (\d+),(\d+)$`)

	coverageStraight := make(map[Vector2]int)
	coverageDiagonal := make(map[Vector2]int)

	for _, line := range lines {
		if line == "" {
			continue
		}

		matches := lineRegex.FindStringSubmatch(line)
		if matches == nil {
			panic(line)
		}

		x1 := toInt(matches[1])
		y1 := toInt(matches[2])
		x2 := toInt(matches[3])
		y2 := toInt(matches[4])

		if x1 == x2 {
			dir := sign(y2 - y1)
			for y := y1; y != y2+dir; y += dir {
				coverageStraight[Vector2{x1, y}]++
				coverageDiagonal[Vector2{x1, y}]++
			}
		} else if y1 == y2 {
			dir := sign(x2 - x1)
			for x := x1; x != x2+dir; x += dir {
				coverageStraight[Vector2{x, y1}]++
				coverageDiagonal[Vector2{x, y1}]++
			}
		} else {
			xdir := sign(x2 - x1)
			ydir := sign(y2 - y1)
			for x, y := x1, y1; x != x2+xdir && y != y2+ydir; x, y = x+xdir, y+ydir {
				coverageDiagonal[Vector2{x, y}]++
			}
		}
	}

	{
		fmt.Println("--- Part One ---")
		count := 0
		for _, c := range coverageStraight {
			if c > 1 {
				count++
			}
		}
		fmt.Println(count)
	}

	{
		fmt.Println("--- Part Two ---")
		count := 0
		for _, c := range coverageDiagonal {
			if c > 1 {
				count++
			}
		}
		fmt.Println(count)
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

func sign(x int) int {
	if x > 0 {
		return 1
	}
	if x < 0 {
		return -1
	}
	return 0
}
