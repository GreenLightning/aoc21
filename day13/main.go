package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Vector2 struct {
	x, y int
}

func (v Vector2) Max(other Vector2) Vector2 {
	return Vector2{
		x: max(v.x, other.x),
		y: max(v.y, other.y),
	}
}

func main() {
	lines := readLines("input.txt")

	var points []Vector2
	var folds []Vector2

	i := 0
	for ; i < len(lines); i++ {
		line := lines[i]
		if line == "" {
			i++
			break
		}
		parts := arrayToInt(strings.Split(line, ","))
		points = append(points, Vector2{parts[0], parts[1]})
	}

	foldRegex := regexp.MustCompile(`^fold along (x|y)=(\d+)$`)
	for ; i < len(lines); i++ {
		matches := foldRegex.FindStringSubmatch(lines[i])
		if matches[1] == "x" {
			folds = append(folds, Vector2{x: toInt(matches[2])})
		} else {
			folds = append(folds, Vector2{y: toInt(matches[2])})
		}
	}

	for fi, fold := range folds {
		for pi, point := range points {
			if fold.x > 0 {
				if point.x > fold.x {
					points[pi].x = fold.x - (point.x - fold.x)
				}
			} else {
				if point.y > fold.y {
					points[pi].y = fold.y - (point.y - fold.y)
				}
			}
		}

		if fi == 0 {
			fmt.Println("--- Part One ---")
			visible := make(map[Vector2]bool)
			for _, point := range points {
				visible[point] = true
			}
			fmt.Println(len(visible))
		}
	}

	{
		fmt.Println("--- Part Two ---")
		visible := make(map[Vector2]bool)
		max := Vector2{0, 0}
		for _, point := range points {
			visible[point] = true
			max = max.Max(point)
		}
		for y := 0; y <= max.y; y++ {
			for x := 0; x <= max.x; x++ {
				if visible[Vector2{x, y}] {
					fmt.Print("#")
				} else {
					fmt.Print(".")
				}
			}
			fmt.Println()
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

func max(x, y int) int {
	if y > x {
		return y
	}
	return x
}
