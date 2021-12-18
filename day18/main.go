package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Number interface {
	ExplodeOnce(level int) (Number, int, int)
	AddLeft(value int) Number
	AddRight(value int) Number
	SplitOnce() Number
	Magnitude() int
}

type Pair struct {
	left, right Number
}

func (pair *Pair) ExplodeOnce(level int) (Number, int, int) {
	if level >= 4 {
		l := pair.left.(*Regular).value
		r := pair.right.(*Regular).value
		return &Regular{0}, l, r
	}

	if result, l, r := pair.left.ExplodeOnce(level + 1); result != nil {
		return &Pair{result, pair.right.AddLeft(r)}, l, 0
	}

	if result, l, r := pair.right.ExplodeOnce(level + 1); result != nil {
		return &Pair{pair.left.AddRight(l), result}, 0, r
	}

	return nil, 0, 0
}

func (pair *Pair) AddLeft(value int) Number {
	return &Pair{pair.left.AddLeft(value), pair.right}
}

func (pair *Pair) AddRight(value int) Number {
	return &Pair{pair.left, pair.right.AddRight(value)}
}

func (pair *Pair) SplitOnce() Number {
	if result := pair.left.SplitOnce(); result != nil {
		return &Pair{result, pair.right}
	}
	if result := pair.right.SplitOnce(); result != nil {
		return &Pair{pair.left, result}
	}
	return nil
}

func (pair *Pair) Magnitude() int {
	return 3*pair.left.Magnitude() + 2*pair.right.Magnitude()
}

type Regular struct {
	value int
}

func (r *Regular) ExplodeOnce(level int) (Number, int, int) {
	return nil, 0, 0
}

func (r *Regular) AddLeft(value int) Number {
	return &Regular{r.value + value}
}

func (r *Regular) AddRight(value int) Number {
	return &Regular{r.value + value}
}

func (r *Regular) SplitOnce() Number {
	if r.value >= 10 {
		return &Pair{&Regular{r.value / 2}, &Regular{(r.value + 1) / 2}}
	}
	return nil
}

func (r *Regular) Magnitude() int {
	return r.value
}

func reduce(number Number) Number {
	for {
		if result, _, _ := number.ExplodeOnce(0); result != nil {
			number = result
			continue
		}
		if result := number.SplitOnce(); result != nil {
			number = result
			continue
		}
		return number
	}
}

func add(left, right Number) Number {
	result := &Pair{left, right}
	return reduce(result)
}

func main() {
	lines := readLines("input.txt")

	var numbers []Number
	for _, line := range lines {
		numbers = append(numbers, parseNumber(strings.NewReader(line)))
	}

	{
		fmt.Println("--- Part One ---")
		result := numbers[0]
		for i := 1; i < len(numbers); i++ {
			result = add(result, numbers[i])
		}
		fmt.Println(result.Magnitude())
	}

	{
		fmt.Println("--- Part Two ---")
		best := 0
		for i := 0; i < len(numbers); i++ {
			for j := i + 1; j < len(numbers); j++ {
				if magnitude := add(numbers[i], numbers[j]).Magnitude(); magnitude > best {
					best = magnitude
				}
				if magnitude := add(numbers[j], numbers[i]).Magnitude(); magnitude > best {
					best = magnitude
				}
			}
		}
		fmt.Println(best)
	}
}

func parseNumber(reader io.RuneScanner) Number {
	r, _, err := reader.ReadRune()
	check(err)

	if r == '[' {
		left := parseNumber(reader)
		r, _, err = reader.ReadRune()
		check(err)
		if r != ',' {
			panic(r)
		}
		right := parseNumber(reader)
		r, _, err = reader.ReadRune()
		check(err)
		if r != ']' {
			panic(r)
		}
		return &Pair{left, right}
	}

	check(reader.UnreadRune())
	var builder strings.Builder
	for {
		r, _, err := reader.ReadRune()
		if err == io.EOF {
			break
		}
		check(err)
		if !(r >= '0' && r <= '9') {
			check(reader.UnreadRune())
			break
		}
		builder.WriteRune(r)
	}

	return &Regular{toInt(builder.String())}
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
