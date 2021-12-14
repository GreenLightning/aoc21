package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
)

func main() {
	lines := readLines("input.txt")

	template := lines[0]

	rules := make(map[string][2]string)
	ruleRegex := regexp.MustCompile(`^(..) -> (.)$`)
	for _, line := range lines {
		matches := ruleRegex.FindStringSubmatch(line)
		if matches != nil {
			pattern, target := matches[1], matches[2]
			rules[pattern] = [2]string{pattern[0:1] + target, target + pattern[1:2]}
		}
	}

	{
		fmt.Println("--- Part One ---")
		fmt.Println(polymerize(template, rules, 10))
	}

	{
		fmt.Println("--- Part Two ---")
		fmt.Println(polymerize(template, rules, 40))
	}
}

func polymerize(template string, rules map[string][2]string, steps int) int64 {
	polymer := make(map[string]int64)
	for i := 0; i+2 <= len(template); i++ {
		polymer[template[i:i+2]]++
	}

	for step := 0; step < steps; step++ {
		newPolymer := make(map[string]int64)
		for pair, count := range polymer {
			if replacements, ok := rules[pair]; ok {
				newPolymer[replacements[0]] += count
				newPolymer[replacements[1]] += count
			} else {
				newPolymer[pair] += count
			}
		}
		polymer = newPolymer
	}

	frequencies := make(map[string]int64)
	for pair, count := range polymer {
		frequencies[pair[0:1]] += count
		frequencies[pair[1:2]] += count
	}
	frequencies[template[0:1]]++
	frequencies[template[len(template)-1:]]++

	var counts []int64
	for _, count := range frequencies {
		counts = append(counts, count/2)
	}

	sort.Slice(counts, func(i, j int) bool { return counts[i] < counts[j] })
	return counts[len(counts)-1] - counts[0]
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
