package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	nums := readNumbers("input.txt")

	{
		fmt.Println("--- Part One ---")
		count := 0
		for i := 0; i+1 < len(nums); i++ {
			if nums[i+1] > nums[i] {
				count++
			}
		}
		fmt.Println(count)
	}

	{
		fmt.Println("--- Part Two ---")
		count := 0
		for i := 0; i+3 < len(nums); i++ {
			if nums[i+3] > nums[i] {
				count++
			}
		}
		fmt.Println(count)
	}
}

func readNumbers(filename string) []int {
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	var numbers []int
	for scanner.Scan() {
		numbers = append(numbers, toInt(scanner.Text()))
	}
	return numbers
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
