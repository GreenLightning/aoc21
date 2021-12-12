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
		paths := buildPaths([]Path{}, Path{"start"}, connections, false, "")
		fmt.Println(len(paths))
	}

	{
		fmt.Println("--- Part Two ---")
		paths := buildPaths([]Path{}, Path{"start"}, connections, true, "")
		fmt.Println(len(paths))
	}
}

// buildPaths constructs all paths starting with prefix using the available connections.
// This function appends its results to paths and then returns paths (like append()).
// If twiceAllowed is true, it is allowed to visit a small cave twice after the prefix.
// If mustVisit is not empty, this small cave must be visited again after the prefix.
func buildPaths(paths []Path, prefix Path, connections []Connection, twiceAllowed bool, mustVisit string) []Path {

	// First we have to find a connection that starts with the last element of the prefix.
	connector := prefix[len(prefix)-1]
	for _, conn := range connections {
		if connector != conn[0] {
			continue
		}

		end := conn[1]
		isSmall := (end == strings.ToLower(end))

		// Make a new path from prefix and the given connection.
		path := make(Path, len(prefix)+1)
		copy(path, prefix)
		path[len(path)-1] = end

		// If we are at "end" this is a complete path.
		if end == "end" {
			// But we only add it to the output if it is valid,
			// i.e. we do not have a small cave left to visit.
			if mustVisit == "" {
				paths = append(paths, path)
			}
			continue
		}

		if isSmall && twiceAllowed {
			// Build all the paths that use end twice by using
			// the same connections and setting mustVisit to end.
			paths = buildPaths(paths, path, connections, false, end)
		}

		var filteredConnections = connections
		if isSmall {
			// If this is a small cave, we remove all the connections leading
			// to the cave to avoid visiting it again.
			filteredConnections = make([]Connection, 0, len(connections))
			for _, conn := range connections {
				if conn[1] != end {
					filteredConnections = append(filteredConnections, conn)
				}
			}
		}

		var filteredMustVisit = mustVisit
		if end == mustVisit {
			filteredMustVisit = ""
		}

		paths = buildPaths(paths, path, filteredConnections, twiceAllowed, filteredMustVisit)
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
