package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Vector3 struct {
	x, y, z int
}

func (v Vector3) GetAxis(axis int) int {
	switch axis {
	case 0:
		return v.x
	case 1:
		return v.y
	case 2:
		return v.z
	default:
		panic(fmt.Sprintf("invalid axis: %d", axis))
	}
}

func (v *Vector3) SetAxis(axis int, value int) {
	switch axis {
	case 0:
		v.x = value
	case 1:
		v.y = value
	case 2:
		v.z = value
	default:
		panic(fmt.Sprintf("invalid axis: %d", axis))
	}
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

func (v Vector3) Times(factor int) Vector3 {
	return Vector3{
		x: factor * v.x,
		y: factor * v.y,
		z: factor * v.z,
	}
}

func (v Vector3) Min(other Vector3) Vector3 {
	return Vector3{
		x: min(v.x, other.x),
		y: min(v.y, other.y),
		z: min(v.z, other.z),
	}
}

func (v Vector3) Max(other Vector3) Vector3 {
	return Vector3{
		x: max(v.x, other.x),
		y: max(v.y, other.y),
		z: max(v.z, other.z),
	}
}

type Cuboid struct {
	min Vector3
	max Vector3
}

func (c *Cuboid) Valid() bool {
	return c.min.x <= c.max.x && c.min.y <= c.max.y && c.min.z <= c.max.z
}

func (c *Cuboid) Size() Vector3 {
	return c.max.Minus(c.min).Plus(Vector3{1, 1, 1})
}

func (c *Cuboid) Contains(other Cuboid) bool {
	return c.min.x <= other.min.x && other.max.x <= c.max.x &&
		c.min.y <= other.min.y && other.max.y <= c.max.y &&
		c.min.z <= other.min.z && other.max.z <= c.max.z
}

func (c *Cuboid) Intersects(other Cuboid) bool {
	return !(c.max.x < other.min.x || other.max.x < c.min.x ||
		c.max.y < other.min.y || other.max.y < c.min.y ||
		c.max.z < other.min.z || other.max.z < c.min.z)
}

func (c *Cuboid) Union(other Cuboid) Cuboid {
	return Cuboid{
		c.min.Min(other.min),
		c.max.Max(other.max),
	}
}

func (c *Cuboid) TrimmedTo(region Cuboid) Cuboid {
	return Cuboid{
		c.min.Max(region.min),
		c.max.Min(region.max),
	}
}

func (c Cuboid) String() string {
	return fmt.Sprintf("x=%d..%d, y=%d..%d, z=%d..%d", c.min.x, c.max.x, c.min.y, c.max.y, c.min.z, c.max.z)
}

// Node is part of a kd-tree of "on" cuboids.
type Node struct {
	leaf bool

	// for leaf nodes
	cuboid Cuboid

	// for inner nodes
	axis  int
	value int
	left  *Node
	right *Node
}

// node must be a leaf node.
func add(node *Node, cuboid Cuboid) *Node {
	if node.cuboid.Contains(cuboid) {
		return node
	}

	if cuboid.Contains(node.cuboid) {
		return &Node{leaf: true, cuboid: cuboid}
	}

	// If we can find a separating axis, we can trivially construct an inner node.
	for a := 0; a < 3; a++ {
		if cuboid.max.GetAxis(a) < node.cuboid.min.GetAxis(a) {
			leaf := &Node{leaf: true, cuboid: cuboid}
			return &Node{axis: a, value: cuboid.max.GetAxis(a), left: leaf, right: node}
		}
		if cuboid.min.GetAxis(a) > node.cuboid.max.GetAxis(a) {
			leaf := &Node{leaf: true, cuboid: cuboid}
			return &Node{axis: a, value: cuboid.min.GetAxis(a) - 1, left: node, right: leaf}
		}
	}

	// Find an axis with partial overlap and split the cuboid,
	// then recursively handle the part that still overlaps the node.
	for a := 0; a < 3; a++ {
		if v := node.cuboid.min.GetAxis(a); cuboid.min.GetAxis(a) < v && cuboid.max.GetAxis(a) >= v {
			outside := cuboid
			outside.max.SetAxis(a, v-1)
			leaf := &Node{leaf: true, cuboid: outside}
			inside := cuboid
			inside.min.SetAxis(a, v)
			return &Node{axis: a, value: v - 1, left: leaf, right: add(node, inside)}
		}
		if v := node.cuboid.max.GetAxis(a); cuboid.max.GetAxis(a) > v && cuboid.min.GetAxis(a) <= v {
			outside := cuboid
			outside.min.SetAxis(a, v+1)
			leaf := &Node{leaf: true, cuboid: outside}
			inside := cuboid
			inside.max.SetAxis(a, v)
			return &Node{axis: a, value: v, left: add(node, inside), right: leaf}
		}
	}

	panic("invalid state in add")
}

// node must be a leaf node.
func remove(node *Node, cuboid Cuboid) *Node {
	if cuboid.Contains(node.cuboid) {
		return nil
	}

	if !cuboid.Intersects(node.cuboid) {
		return node
	}

	// Find an axis with partial overlap and split the node,
	// then recursively handle the part that still overlaps the cuboid.
	for a := 0; a < 3; a++ {
		if v := cuboid.min.GetAxis(a); node.cuboid.min.GetAxis(a) < v && node.cuboid.max.GetAxis(a) >= v {
			outside := node.cuboid
			outside.max.SetAxis(a, v-1)
			left := &Node{leaf: true, cuboid: outside}

			inside := node.cuboid
			inside.min.SetAxis(a, v)
			right := &Node{leaf: true, cuboid: inside}
			right = remove(right, cuboid)
			if right == nil {
				return left
			}

			return &Node{axis: a, value: v - 1, left: left, right: right}
		}
		if v := cuboid.max.GetAxis(a); node.cuboid.max.GetAxis(a) > v && node.cuboid.min.GetAxis(a) <= v {
			outside := node.cuboid
			outside.min.SetAxis(a, v+1)
			right := &Node{leaf: true, cuboid: outside}

			inside := node.cuboid
			inside.max.SetAxis(a, v)
			left := &Node{leaf: true, cuboid: inside}
			left = remove(left, cuboid)
			if left == nil {
				return right
			}

			return &Node{axis: a, value: v, left: left, right: right}
		}
	}

	panic("invalid state in remove")
}

func execute(node *Node, on bool, cuboid Cuboid) *Node {
	if node == nil {
		if on {
			return &Node{leaf: true, cuboid: cuboid}
		} else {
			return nil
		}
	}

	if !node.leaf {
		// Split cuboid along our split plane and execute on our children.
		leftCuboid, rightCuboid := cuboid, cuboid
		leftCuboid.max.SetAxis(node.axis, min(leftCuboid.max.GetAxis(node.axis), node.value))
		rightCuboid.min.SetAxis(node.axis, max(rightCuboid.min.GetAxis(node.axis), node.value+1))
		if leftCuboid.Valid() {
			node.left = execute(node.left, on, leftCuboid)
		}
		if rightCuboid.Valid() {
			node.right = execute(node.right, on, rightCuboid)
		}
		if node.left == nil {
			return node.right
		}
		if node.right == nil {
			return node.left
		}
		return node
	}

	if on {
		return add(node, cuboid)
	} else {
		return remove(node, cuboid)
	}
}

func count(node *Node) int64 {
	if node == nil {
		return 0
	}
	if node.leaf {
		size := node.cuboid.Size()
		return int64(size.x) * int64(size.y) * int64(size.z)
	}
	return count(node.left) + count(node.right)
}

func main() {
	lines := readLines("input.txt")
	lineRegex := regexp.MustCompile(`(on|off) x=(-?\d+)..(-?\d+),y=(-?\d+)..(-?\d+),z=(-?\d+)..(-?\d+)`)
	initializationRegion := Cuboid{Vector3{-50, -50, -50}, Vector3{50, 50, 50}}

	var part1 *Node
	var part2 *Node

	for _, line := range lines {
		matches := lineRegex.FindStringSubmatch(line)

		on := matches[1] == "on"
		cuboid := Cuboid{
			Vector3{toInt(matches[2]), toInt(matches[4]), toInt(matches[6])},
			Vector3{toInt(matches[3]), toInt(matches[5]), toInt(matches[7])},
		}

		if trimmed := cuboid.TrimmedTo(initializationRegion); trimmed.Valid() {
			part1 = execute(part1, on, trimmed)
		}

		part2 = execute(part2, on, cuboid)
	}

	{
		fmt.Println("--- Part One ---")
		fmt.Println(count(part1))
	}

	{
		fmt.Println("--- Part Two ---")
		fmt.Println(count(part2))
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

func min(x, y int) int {
	if y < x {
		return y
	}
	return x
}

func max(x, y int) int {
	if y > x {
		return y
	}
	return x
}
