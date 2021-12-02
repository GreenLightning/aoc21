package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	lines := readLines("input.txt")
	regex := regexp.MustCompile(`^(forward|down|up) (\d+)$`)

	{
		fmt.Println("--- Part One ---")
		var pos, depth int
		for _, line := range lines {
			matches := regex.FindStringSubmatch(line)
			if matches == nil {
				panic(line)
			}

			value := toInt(matches[2])
			switch matches[1] {
			case "forward":
				pos += value
			case "down":
				depth += value
			case "up":
				depth -= value
			}
		}
		fmt.Println(pos * depth)
	}

	{
		fmt.Println("--- Part Two ---")
		var pos, depth, aim int
		for _, line := range lines {
			matches := regex.FindStringSubmatch(line)
			if matches == nil {
				panic(line)
			}

			value := toInt(matches[2])
			switch matches[1] {
			case "forward":
				pos += value
				depth += aim * value
			case "down":
				aim += value
			case "up":
				aim -= value
			}
		}
		fmt.Println(pos * depth)
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
