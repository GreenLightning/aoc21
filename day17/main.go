package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input := readFile("input.txt")
	inputRegex := regexp.MustCompile(`target area: x=(-?\d+)\.\.(-?\d+), y=(-?\d+)\.\.(-?\d+)`)
	matches := inputRegex.FindStringSubmatch(input)

	x1 := toInt(matches[1])
	x2 := toInt(matches[2])
	y1 := toInt(matches[3])
	y2 := toInt(matches[4])
	xmin, xmax := min(x1, x2), max(x1, x2)
	ymin, ymax := min(y1, y2), max(y1, y2)

	count := 0
	highestOfAll := 0
	for iy := -10_000; iy < 10_000; iy++ {
		for ix := 1; ; ix++ {

			px, py := 0, 0
			vx, vy := ix, iy
			hit := false
			highest := 0
			for {
				px += vx
				py += vy
				vx -= sign(vx)
				vy -= 1

				if py > highest {
					highest = py
				}

				if xmin <= px && px <= xmax && ymin <= py && py <= ymax {
					hit = true
					break
				}

				if px > xmax || py < ymin {
					break
				}
			}

			if hit {
				count++
				if highest > highestOfAll {
					highestOfAll = highest
				}
			}

			if px > xmax {
				break
			}

		}
	}

	{
		fmt.Println("--- Part One ---")
		fmt.Println(highestOfAll)
	}

	{
		fmt.Println("--- Part Two ---")
		fmt.Println(count)
	}
}

func readFile(filename string) string {
	bytes, err := ioutil.ReadFile(filename)
	check(err)
	return strings.TrimSpace(string(bytes))
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

func min(x, y int) int {
	if y < x {
		return y
	}
	return x
}

func max(x, y int) int {
	if y > x {
		return y
	}
	return x
}
