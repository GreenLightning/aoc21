package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Input struct {
	Patterns []string
	Outputs  []string
}

func main() {
	lines := readLines("input.txt")

	var inputs []Input
	for _, line := range lines {
		parts := strings.Split(line, "|")
		patterns := strings.Split(strings.TrimSpace(parts[0]), " ")
		outputs := strings.Split(strings.TrimSpace(parts[1]), " ")
		inputs = append(inputs, Input{Patterns: patterns, Outputs: outputs})
	}

	{
		fmt.Println("--- Part One ---")
		var count int
		for _, input := range inputs {
			for _, output := range input.Outputs {
				if len(output) == 2 || len(output) == 3 || len(output) == 4 || len(output) == 7 {
					count++
				}
			}
		}
		fmt.Println(count)
	}

	{
		fmt.Println("--- Part Two ---")
		var sum int

		permutations := generatePermutations([]int{0, 1, 2, 3, 4, 5, 6})

		// We will encode a seven-segment digit into an integer, where each
		// bit indicates whether the corresponding segment is on or off.
		// The values map then maps valid codes to their numerical values.
		values := make(map[int]int)
		//       gfedcba
		values[0b1110111] = 0
		values[0b0100100] = 1
		values[0b1011101] = 2
		values[0b1101101] = 3
		values[0b0101110] = 4
		values[0b1101011] = 5
		values[0b1111011] = 6
		values[0b0100101] = 7
		values[0b1111111] = 8
		values[0b1101111] = 9

		for _, input := range inputs {
		perms:
			for _, permutation := range permutations {
				for _, pattern := range input.Patterns {
					code := encode(pattern, permutation)
					if _, ok := values[code]; !ok {
						continue perms
					}
				}
				place := 1
				for i := len(input.Outputs) - 1; i >= 0; i-- {
					code := encode(input.Outputs[i], permutation)
					sum += values[code] * place
					place *= 10
				}
				break
			}
		}

		fmt.Println(sum)
	}
}

func generatePermutations(input []int) (permutations [][]int) {
	// QuickPerm implementation.
	// https://www.quickperm.org/

	elements := make([]int, len(input))
	copy(elements, input)

	permutation := make([]int, len(elements))
	copy(permutation, elements)
	permutations = append(permutations, permutation)

	p := make([]int, len(elements)+1)
	for i := range p {
		p[i] = i
	}
	
	for i := 1; i < len(elements); {
		p[i]--

		j := 0
		if i%2 == 1 {
			j = p[i]
		}

		elements[j], elements[i] = elements[i], elements[j]

		permutation := make([]int, len(elements))
		copy(permutation, elements)
		permutations = append(permutations, permutation)

		for i = 1; p[i] == 0; i++ {
			p[i] = i
		}
	}

	return
}

// encode takes a textual representation of a seven-segment digit using the
// letters 'a' through 'g' to denote the segments and a permutation of the
// range [0, 6] for mapping each letter to a different segment. It returns a
// codeword, where each bit indicates whether the corresponding segment is on
// or off.
func encode(text string, permutation []int) (code int) {
	for _, rune := range text {
		index := int(rune - 'a')
		index = permutation[index]
		code |= 1 << index
	}
	return
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
