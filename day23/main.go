package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

type Board [28]rune

type State struct {
	board  Board
	energy int
}

var energyPerStep = [4]int{1, 10, 100, 1000}

func main() {
	lines := readLines("input.txt")

	{
		fmt.Println("--- Part One ---")
		fmt.Println(solve(lines, 8))
	}

	{
		fmt.Println("--- Part Two ---")
		newLines := make([]string, len(lines)+2)
		copy(newLines, lines[:len(lines)-2])
		newLines[len(lines)-2] = "  #D#C#B#A#"
		newLines[len(lines)-1] = "  #D#B#A#C#"
		copy(newLines[len(lines):], lines[len(lines)-2:])
		fmt.Println(solve(newLines, 16))
	}
}

func solve(lines []string, target int) int {
	grid := make([][]rune, len(lines))
	for y, line := range lines {
		grid[y] = []rune(line)
	}

	var queue PriorityQueue
	queue.Push(State{board: serialize(grid)})

	var visited = make(map[Board]bool)

	for !queue.Empty() {
		state := queue.Pop()
		deserialize(grid, state.board)

		good := 0
		for ay, row := range grid {
			for ax, actor := range row {
				if !(actor >= 'A' && actor <= 'D') {
					continue
				}

				// In room?
				if ay >= 2 {

					// Already in target position?
					if ax == 3+2*int(actor-'A') {
						ok := true
						for y := ay + 1; grid[y][ax] != '#'; y++ {
							if grid[y][ax] != actor {
								ok = false
								break
							}
						}
						if ok {
							good++
							continue
						}
					}

					// Blocked?
					if ay >= 3 && grid[ay-1][ax] != '.' {
						continue
					}

					for tx := ax - 1; ; tx-- {
						if tx == 3 || tx == 5 || tx == 7 || tx == 9 {
							continue
						}

						if grid[1][tx] != '.' {
							break
						}

						grid[1][tx] = actor
						grid[ay][ax] = '.'

						steps := (ay - 1) + abs(ax-tx)

						board := serialize(grid)
						if !visited[board] {
							visited[board] = true
							queue.Push(State{
								board:  board,
								energy: state.energy + steps*energyPerStep[actor-'A'],
							})
						}

						grid[1][tx] = '.'
						grid[ay][ax] = actor
					}

					for tx := ax + 1; ; tx++ {
						if tx == 3 || tx == 5 || tx == 7 || tx == 9 {
							continue
						}

						if grid[1][tx] != '.' {
							break
						}

						grid[1][tx] = actor
						grid[ay][ax] = '.'

						steps := (ay - 1) + abs(ax-tx)

						board := serialize(grid)
						if !visited[board] {
							visited[board] = true
							queue.Push(State{
								board:  board,
								energy: state.energy + steps*energyPerStep[actor-'A'],
							})
						}

						grid[1][tx] = '.'
						grid[ay][ax] = actor
					}

				} else {

					tx := 3 + 2*int(actor-'A')

					ty := len(lines) - 2
					for grid[ty][tx] == actor {
						ty--
					}

					// Room blocked?
					if grid[ty][tx] != '.' {
						continue
					}

					// Hallway blocked?
					dx := sign(tx - ax)
					blocked := false
					for x := ax + dx; x != tx; x += dx {
						if grid[ay][x] != '.' {
							blocked = true
						}
					}
					if blocked {
						continue
					}

					grid[ty][tx] = actor
					grid[ay][ax] = '.'

					steps := abs(ay-ty) + abs(ax-tx)

					board := serialize(grid)
					if !visited[board] {
						visited[board] = true
						queue.Push(State{
							board:  board,
							energy: state.energy + steps*energyPerStep[actor-'A'],
						})
					}

					grid[ty][tx] = '.'
					grid[ay][ax] = actor
				}
			}
		}

		if good == target {
			return state.energy
		}
	}

	return -1
}

func serialize(grid [][]rune) (board Board) {
	index := 0
	for _, row := range grid {
		for _, char := range row {
			if char == '.' || (char >= 'A' && char <= 'D') {
				board[index] = char
				index++
			}
		}
	}
	return
}

func deserialize(grid [][]rune, board Board) {
	index := 0
	for _, row := range grid {
		for x, char := range row {
			if char == '.' || (char >= 'A' && char <= 'D') {
				row[x] = board[index]
				index++
			}
		}
	}
	return
}

type PriorityStorage []State

func (s PriorityStorage) Len() int {
	return len(s)
}

func (s PriorityStorage) Less(i, j int) bool {
	return s[i].energy < s[j].energy
}

func (s PriorityStorage) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s *PriorityStorage) Push(x interface{}) {
	item := x.(State)
	*s = append(*s, item)
}

func (s *PriorityStorage) Pop() interface{} {
	len := len(*s)
	item := (*s)[len-1]
	*s = (*s)[:len-1]
	return item
}

type PriorityQueue struct {
	storage PriorityStorage
}

func (q *PriorityQueue) Len() int {
	return len(q.storage)
}

func (q *PriorityQueue) Empty() bool {
	return len(q.storage) == 0
}

func (q *PriorityQueue) Push(item State) {
	heap.Push(&q.storage, item)
}

func (q *PriorityQueue) Pop() State {
	return heap.Pop(&q.storage).(State)
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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
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
