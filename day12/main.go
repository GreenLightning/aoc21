package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Connection []string
type Path []string

func main() {
	lines := readLines("input.txt")

	var connections []Connection
	for _, line := range lines {
		if line != "" {
			parts := strings.Split(line, "-")
			connections = append(connections, Connection{parts[0], parts[1]})
			connections = append(connections, Connection{parts[1], parts[0]})
		}
	}

	// Remove connections that end at "start" or start at "end".
	for i := len(connections) - 1; i >= 0; i-- {
		conn := connections[i]
		if conn[1] == "start" || conn[0] == "end" {
			copy(connections[i:], connections[i+1:])
			connections = connections[:len(connections)-1]
		}
	}

	{
		fmt.Println("--- Part One ---")
		paths := buildPaths([]Path{}, Path{"start"}, connections, false)
		fmt.Println(len(paths))
	}

	{
		fmt.Println("--- Part Two ---")
		paths := buildPaths([]Path{}, Path{"start"}, connections, true)

		unique := make(map[string]bool)
		for _, path := range paths {
			unique[fmt.Sprint(path)] = true
		}

		fmt.Println(len(unique))
	}
}

// buildPaths constructs all paths starting with prefix using the available connections.
// This function appends its results to paths and then returns paths (like append()).
// If extra is true it is allowed to visit a small cave twice after the prefix.
func buildPaths(paths []Path, prefix Path, connections []Connection, extra bool) []Path {
	connector := prefix[len(prefix)-1]
	for _, conn := range connections {
		if connector != conn[0] {
			continue
		}
		end := conn[1]
		path := make(Path, len(prefix)+1)
		copy(path, prefix)
		path[len(path)-1] = end
		if end == "end" {
			paths = append(paths, path)
			continue
		}
		if extra {
			paths = buildPaths(paths, path, connections, false)
		}
		var filteredConnections = connections
		if end == strings.ToLower(end) {
			filteredConnections = make([]Connection, 0, len(connections))
			for _, conn := range connections {
				if conn[1] != end {
					filteredConnections = append(filteredConnections, conn)
				}
			}
		}
		paths = buildPaths(paths, path, filteredConnections, extra)
	}
	return paths
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
