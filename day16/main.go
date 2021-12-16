package main

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Packet struct {
	Version  int
	Type     int
	Value    int64
	Children []Packet
}

func main() {
	// Convert input to a string of binary digits so that
	// we can easily address each bit individually.
	bytes, err := hex.DecodeString(readFile("input.txt"))
	check(err)

	var builder strings.Builder
	for _, byte := range bytes {
		fmt.Fprintf(&builder, "%08b", byte)
	}
	input := builder.String()

	root, _ := parsePacket(input)

	{
		fmt.Println("--- Part One ---")
		fmt.Println(sumVersions(root))
	}

	{
		fmt.Println("--- Part Two ---")
		fmt.Println(calculate(root))
	}
}

func sumVersions(packet Packet) (sum int) {
	sum += packet.Version
	for _, child := range packet.Children {
		sum += sumVersions(child)
	}
	return
}

func calculate(packet Packet) (result int64) {
	switch packet.Type {
	case 0:
		for _, child := range packet.Children {
			result += calculate(child)
		}

	case 1:
		result = calculate(packet.Children[0])
		for i := 1; i < len(packet.Children); i++ {
			result *= calculate(packet.Children[i])
		}

	case 2:
		result = calculate(packet.Children[0])
		for i := 1; i < len(packet.Children); i++ {
			result = min(result, calculate(packet.Children[i]))
		}

	case 3:
		result = calculate(packet.Children[0])
		for i := 1; i < len(packet.Children); i++ {
			result = max(result, calculate(packet.Children[i]))
		}

	case 4:
		result = packet.Value

	case 5:
		if calculate(packet.Children[0]) > calculate(packet.Children[1]) {
			result = 1
		}

	case 6:
		if calculate(packet.Children[0]) < calculate(packet.Children[1]) {
			result = 1
		}

	case 7:
		if calculate(packet.Children[0]) == calculate(packet.Children[1]) {
			result = 1
		}
	}

	return
}

func parsePacket(input string) (Packet, string) {
	var packet Packet
	packet.Version, input = parseBinaryChunk(input, 3)
	packet.Type, input = parseBinaryChunk(input, 3)

	if packet.Type == 4 {
		var builder strings.Builder
		for {
			bit := input[:1]
			builder.WriteString(input[1:5])
			input = input[5:]
			if bit != "1" {
				break
			}
		}

		var err error
		packet.Value, err = strconv.ParseInt(builder.String(), 2, 0)
		check(err)

		return packet, input
	}

	bit := input[:1]
	input = input[1:]

	if bit == "0" {
		var total int
		total, input = parseBinaryChunk(input, 15)
		for target := len(input) - total; len(input) > target; {
			var child Packet
			child, input = parsePacket(input)
			packet.Children = append(packet.Children, child)
		}
	} else {
		var count int
		count, input = parseBinaryChunk(input, 11)
		for i := 0; i < count; i++ {
			var child Packet
			child, input = parsePacket(input)
			packet.Children = append(packet.Children, child)
		}
	}

	return packet, input
}

func parseBinaryChunk(input string, length int) (int, string) {
	result, err := strconv.ParseInt(input[:length], 2, 0)
	check(err)
	return int(result), input[length:]
}

func readFile(filename string) string {
	bytes, err := ioutil.ReadFile(filename)
	check(err)
	return strings.TrimSpace(string(bytes))
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func min(x, y int64) int64 {
	if y < x {
		return y
	}
	return x
}

func max(x, y int64) int64 {
	if y > x {
		return y
	}
	return x
}
