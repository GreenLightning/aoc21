package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const lenBits = 12

func main() {
	numbers := readBinaryNumbers("input.txt")

	{
		fmt.Println("--- Part One ---")

		bitCounters := make([]int, lenBits)
		for _, number := range numbers {
			for i := range bitCounters {
				if (number>>i)&1 == 1 {
					bitCounters[i]++
				}
			}
		}

		var gamma, epsilon int
		for i, counter := range bitCounters {
			if counter >= len(numbers)/2 {
				gamma |= (1 << i)
			} else {
				epsilon |= (1 << i)
			}
		}

		fmt.Println(gamma * epsilon)
	}

	{
		fmt.Println("--- Part Two ---")
		generator := filter(numbers, lenBits, false)
		scrubber := filter(numbers, lenBits, true)
		fmt.Println(generator * scrubber)
	}
}

func filter(numbers []int, lenBits int, leastCommon bool) int {
	candidates := make([]int, len(numbers))
	filtered := make([]int, len(numbers))
	copy(candidates, numbers)

	for i := lenBits - 1; i >= 0 && len(candidates) > 1; i-- {
		var counter int
		for _, number := range candidates {
			if (number>>i)&1 == 1 {
				counter++
			}
		}

		var target int
		if counter >= len(candidates)-counter {
			target = 1
		}
		if leastCommon {
			target = 1 - target
		}

		filtered = filtered[:0]
		for _, number := range candidates {
			if (number>>i)&1 == target {
				filtered = append(filtered, number)
			}
		}

		candidates, filtered = filtered, candidates
	}

	return candidates[0]
}

func readBinaryNumbers(filename string) []int {
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	var numbers []int
	for scanner.Scan() {
		number, err := strconv.ParseInt(scanner.Text(), 2, 0)
		check(err)
		numbers = append(numbers, int(number))
	}
	return numbers
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
