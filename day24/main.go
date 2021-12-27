package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	var as, bs, cs []int

	lines := readLines("input.txt")
	aRegex := regexp.MustCompile(`^add x (-?\d+)$`)
	bRegex := regexp.MustCompile(`^div z (\d+)$`)
	cRegex := regexp.MustCompile(`^add y (-?\d+)$`)
	for i := 0; i < 14; i++ {
		aMatches := aRegex.FindStringSubmatch(lines[18*i+5])
		as = append(as, toInt(aMatches[1]))
		bMatches := bRegex.FindStringSubmatch(lines[18*i+4])
		bs = append(bs, toInt(bMatches[1]))
		cMatches := cRegex.FindStringSubmatch(lines[18*i+15])
		cs = append(cs, toInt(cMatches[1]))
	}

	digits := make([]int, 14)

	fmt.Println("--- Part One ---")
	for i := range digits {
		digits[i] = 9
	}
loop1:
	for digits[0] >= 1 {
		z := 0
		for i, d := range digits {
			t := z%26 + as[i]
			z /= bs[i]
			if t != d {
				if bs[i] != 1 {
					decrease(digits, i)
					continue loop1
				}
				z = 26*z + (d + cs[i])
			}
		}
		if z == 0 {
			for _, d := range digits {
				fmt.Print(d)
			}
			fmt.Println()
			break
		}
	}

	fmt.Println("--- Part Two ---")
	for i := range digits {
		digits[i] = 1
	}
loop2:
	for digits[0] <= 9 {
		z := 0
		for i, d := range digits {
			t := z%26 + as[i]
			z /= bs[i]
			if t != d {
				if bs[i] != 1 {
					increase(digits, i)
					continue loop2
				}
				z = 26*z + (d + cs[i])
			}
		}
		if z == 0 {
			for _, d := range digits {
				fmt.Print(d)
			}
			fmt.Println()
			break
		}
	}
}

func increase(digits []int, i int) {
	for i > 0 && digits[i] == 9 {
		i--
	}
	digits[i]++
	for j := i + 1; j < len(digits); j++ {
		digits[j] = 1
	}
}

func decrease(digits []int, i int) {
	for i > 0 && digits[i] == 1 {
		i--
	}
	digits[i]--
	for j := i + 1; j < len(digits); j++ {
		digits[j] = 9
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
