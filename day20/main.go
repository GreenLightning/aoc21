package main

import (
	"bufio"
	"fmt"
	"os"
)

type Vector2 struct {
	x, y int
}

func (v Vector2) Plus(other Vector2) Vector2 {
	return Vector2{
		x: v.x + other.x,
		y: v.y + other.y,
	}
}

func (v Vector2) Minus(other Vector2) Vector2 {
	return Vector2{
		x: v.x - other.x,
		y: v.y - other.y,
	}
}

type Image struct {
	Background bool
	Pixels     map[Vector2]bool
	Min        Vector2
	Max        Vector2
}

func NewImage() *Image {
	return &Image{
		Pixels: make(map[Vector2]bool),
	}
}

func (image *Image) Set(pos Vector2, value bool) {
	if value != image.Background {
		image.Pixels[pos] = value
	}
}

func (image *Image) Get(pos Vector2) bool {
	if value, ok := image.Pixels[pos]; ok {
		return value
	}
	return image.Background
}

func (image *Image) Clear() {
	image.Background = false
	for pos := range image.Pixels {
		delete(image.Pixels, pos)
	}
	image.Min = Vector2{}
	image.Max = Vector2{}
}

func (image *Image) Print() {
	for y := image.Min.y; y <= image.Max.y; y++ {
		for x := image.Min.x; x <= image.Max.x; x++ {
			if image.Get(Vector2{x, y}) {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func main() {
	lines := readLines("input.txt")

	theAlgorithm := lines[0]

	image := NewImage()
	image.Max = Vector2{len(lines[2]) - 1, len(lines[2:]) - 1}
	for y, line := range lines[2:] {
		for x, char := range line {
			image.Set(Vector2{x, y}, char == '#')
		}
	}

	newImage := NewImage()
	for i := 1; i <= 50; i++ {
		if image.Background {
			newImage.Background = (theAlgorithm[511] == '#')
		} else {
			newImage.Background = (theAlgorithm[0] == '#')
		}
		newImage.Min = image.Min.Minus(Vector2{1, 1})
		newImage.Max = image.Max.Plus(Vector2{1, 1})
		for y := newImage.Min.y; y <= newImage.Max.y; y++ {
			for x := newImage.Min.x; x <= newImage.Max.x; x++ {
				var index int
				p := 8
				for dy := -1; dy <= 1; dy++ {
					for dx := -1; dx <= 1; dx++ {
						if image.Get(Vector2{x + dx, y + dy}) {
							index |= (1 << p)
						}
						p--
					}
				}
				newImage.Set(Vector2{x, y}, theAlgorithm[index] == '#')
			}
		}
		image, newImage = newImage, image
		newImage.Clear()

		if i == 2 {
			fmt.Println("--- Part One ---")
			fmt.Println(len(image.Pixels))
		}
	}

	{
		fmt.Println("--- Part Two ---")
		fmt.Println(len(image.Pixels))
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

func check(err error) {
	if err != nil {
		panic(err)
	}
}
