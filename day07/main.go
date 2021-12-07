package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func main() {
	positions := arrayToInt(strings.Split(readFile("input.txt"), ","))

	min, max := positions[0], positions[0]
	for _, p := range positions {
		if p < min {
			min = p
		} else if p > max {
			max = p
		}
	}

	{
		fmt.Println("--- Part One ---")
		best := math.MaxInt
		for i := min; i <= max; i++ {
			fuel := 0
			for _, p := range positions {
				fuel += abs(p - i)
			}
			if fuel < best {
				best = fuel
			}
		}
		fmt.Println(best)
	}

	{
		fmt.Println("--- Part Two ---")
		best := math.MaxInt
		for i := min; i <= max; i++ {
			fuel := 0
			for _, p := range positions {
				distance := abs(p - i)
				fuel += distance * (distance + 1) / 2
			}
			if fuel < best {
				best = fuel
			}
		}
		fmt.Println(best)
	}
}

func readFile(filename string) string {
	bytes, err := ioutil.ReadFile(filename)
	check(err)
	return strings.TrimSpace(string(bytes))
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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
