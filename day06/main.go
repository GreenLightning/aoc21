package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	school := arrayToInt(strings.Split(readFile("input.txt"), ","))

	{
		fmt.Println("--- Part One ---")
		var total int64
		for _, fish := range school {
			total += simulate(fish, 80)
		}
		fmt.Println(total)
	}

	{
		fmt.Println("--- Part Two ---")
		var total int64
		for _, fish := range school {
			total += simulate(fish, 256)
		}
		fmt.Println(total)
	}
}

var memory = make(map[int]int64)

func simulate(fish, days int) int64 {
	if fish >= days {
		return 1
	}
	split := days - fish
	if result, ok := memory[split]; ok {
		return result
	}
	result := simulate(6, split-1) + simulate(8, split-1)
	memory[split] = result
	return result
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
