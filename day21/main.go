package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	var startingPositions [2]int

	lines := readLines("input.txt")
	lineRegex := regexp.MustCompile(`Player (\d+) starting position: (\d+)`)
	for player, line := range lines {
		matches := lineRegex.FindStringSubmatch(line)
		startingPositions[player] = toInt(matches[2]) - 1
	}

	{
		var positions = startingPositions
		var scores [2]int
		var dice, rolls int
		for player := 0; ; player = 1 - player {
			var moves int
			for i := 0; i < 3; i++ {
				moves += dice + 1
				dice = (dice + 1) % 100
				rolls++
			}
			positions[player] = (positions[player] + moves) % 10
			scores[player] += positions[player] + 1
			if scores[player] >= 1000 {
				fmt.Println("--- Part One ---")
				fmt.Println(scores[1-player] * rolls)
				break
			}
		}
	}

	{
		type State struct {
			positions [2]int
			scores    [2]int
		}

		// Map from a move (sum of three dice) to how frequently it occurs.
		var moves = map[int]int64{3: 1, 4: 3, 5: 6, 6: 7, 7: 6, 8: 3, 9: 1}

		var wins [2]int64
		states := make(map[State]int64)
		states[State{positions: startingPositions}] = 1
		for player := 0; len(states) != 0; player = 1 - player {
			newStates := make(map[State]int64)
			for state, stateCount := range states {
				for move, moveCount := range moves {
					newState := state
					newState.positions[player] = (newState.positions[player] + move) % 10
					newState.scores[player] += newState.positions[player] + 1
					if newState.scores[player] >= 21 {
						wins[player] += stateCount * moveCount
					} else {
						newStates[newState] += stateCount * moveCount
					}
				}
			}
			states = newStates
		}

		fmt.Println("--- Part Two ---")
		if wins[1] > wins[0] {
			wins[0] = wins[1]
		}
		fmt.Println(wins[0])
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
