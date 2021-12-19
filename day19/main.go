package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Vector3 struct {
	x, y, z int
}

func (v Vector3) Plus(other Vector3) Vector3 {
	return Vector3{
		x: v.x + other.x,
		y: v.y + other.y,
		z: v.z + other.z,
	}
}

func (v Vector3) Minus(other Vector3) Vector3 {
	return Vector3{
		x: v.x - other.x,
		y: v.y - other.y,
		z: v.z - other.z,
	}
}

func (v Vector3) Cross(o Vector3) Vector3 {
	return Vector3{
		x: v.y*o.z - v.z*o.y,
		y: v.z*o.x - v.x*o.z,
		z: v.x*o.y - v.y*o.x,
	}
}

func (v Vector3) Negated() Vector3 {
	return Vector3{-v.x, -v.y, -v.z}
}

func (v Vector3) ManhattenLength() int {
	return abs(v.x) + abs(v.y) + abs(v.z)
}

func (v Vector3) ManhattenDistance(o Vector3) int {
	return v.Minus(o).ManhattenLength()
}

type Matrix3 [3][3]int

func (m Matrix3) Times(v Vector3) Vector3 {
	return Vector3{
		m[0][0]*v.x + m[0][1]*v.y + m[0][2]*v.z,
		m[1][0]*v.x + m[1][1]*v.y + m[1][2]*v.z,
		m[2][0]*v.x + m[2][1]*v.y + m[2][2]*v.z,
	}
}

var rotations []Matrix3

func init() {
	// Calculate the rotation matrices corresponding to the 24 possible orientations the scanners can be in.
	directions := []Vector3{Vector3{1, 0, 0}, Vector3{0, 1, 0}, Vector3{0, 0, 1}, Vector3{-1, 0, 0}, Vector3{0, -1, 0}, Vector3{0, 0, -1}}
	for _, first := range directions {
		for _, second := range directions {
			if second != first && second.Negated() != first {
				third := first.Cross(second)
				rotations = append(rotations, Matrix3{
					{first.x, second.x, third.x},
					{first.y, second.y, third.y},
					{first.z, second.z, third.z},
				})
			}
		}
	}
}

type Scanner struct {
	ID int

	Position    Vector3 // in world coordinates
	Orientation int     // so that BeaconsByOrientation[orientation] are in world coordinates

	Beacons              []Vector3
	BeaconsByOrientation [][]Vector3
}

// Attempts to align current with target.
// Sets current.Position and current.Orientation and returns true on success.
func align(current *Scanner, target *Scanner) bool {
	// Propose an orientation and two indices that form a correspondence.
	for orientation := 0; orientation < len(rotations); orientation++ {
		for cbi := 0; cbi < len(current.Beacons); cbi++ {
			for tbi := 0; tbi < len(target.Beacons); tbi++ {
				// From this we can calculate the assumed position of the scanner.
				aligningBeaconWorld := target.Position.Plus(target.BeaconsByOrientation[target.Orientation][tbi])
				position := aligningBeaconWorld.Minus(current.BeaconsByOrientation[orientation][cbi])

				// Check if the position is consistent.
				overlap := 0
				for cbi := 0; cbi < len(current.Beacons); cbi++ {
					beaconWorld := position.Plus(current.BeaconsByOrientation[orientation][cbi])
					beaconTarget := beaconWorld.Minus(target.Position)
					if abs(beaconTarget.x) > 1000 || abs(beaconTarget.y) > 1000 || abs(beaconTarget.z) > 1000 {
						continue
					}
					found := false
					for tbi := 0; tbi < len(target.Beacons); tbi++ {
						if target.BeaconsByOrientation[target.Orientation][tbi] == beaconTarget {
							found = true
							break
						}
					}
					if found {
						overlap++
					} else {
						overlap = -1
						break
					}
				}

				if overlap >= 12 {
					current.Position = position
					current.Orientation = orientation
					return true
				}
			}
		}
	}
	return false
}

func main() {
	lines := readLines("input.txt")
	scannerRegex := regexp.MustCompile(`--- scanner (\d+) ---`)
	var scanners []*Scanner
	var currentScanner *Scanner
	for _, line := range lines {
		if line == "" {
			continue
		}

		if matches := scannerRegex.FindStringSubmatch(line); matches != nil {
			currentScanner = &Scanner{ID: toInt(matches[1])}
			scanners = append(scanners, currentScanner)
			continue
		}

		parts := arrayToInt(strings.Split(line, ","))
		currentScanner.Beacons = append(currentScanner.Beacons, Vector3{parts[0], parts[1], parts[2]})
	}

	// Precompute the rotated positions of the beacons depending on the orientation of the scanner.
	for _, scanner := range scanners {
		scanner.BeaconsByOrientation = make([][]Vector3, len(rotations))
		for i, rotation := range rotations {
			beacons := make([]Vector3, len(scanner.Beacons))
			for bi, beacon := range scanner.Beacons {
				beacons[bi] = rotation.Times(beacon)
			}
			scanner.BeaconsByOrientation[i] = beacons
		}
	}

aligning:
	for aligned := 1; aligned < len(scanners); aligned++ {
		for ci := aligned; ci < len(scanners); ci++ {
			current := scanners[ci]
			for ti := 0; ti < aligned; ti++ {
				target := scanners[ti]
				ok := align(current, target)
				if ok {
					scanners[aligned], scanners[ci] = scanners[ci], scanners[aligned]
					continue aligning
				}
			}
		}
		panic("aligning failed")
	}

	{
		fmt.Println("--- Part One ---")
		beacons := make(map[Vector3]bool)
		for _, scanner := range scanners {
			for bi := 0; bi < len(scanner.Beacons); bi++ {
				beaconWorld := scanner.Position.Plus(scanner.BeaconsByOrientation[scanner.Orientation][bi])
				beacons[beaconWorld] = true
			}
		}
		fmt.Println(len(beacons))
	}

	{
		fmt.Println("--- Part Two ---")
		best := 0
		for i := 0; i < len(scanners); i++ {
			for j := i + 1; j < len(scanners); j++ {
				distance := scanners[i].Position.ManhattenDistance(scanners[j].Position)
				if distance > best {
					best = distance
				}
			}
		}
		fmt.Println(best)
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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
