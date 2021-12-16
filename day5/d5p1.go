//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Point struct {
	X int
	Y int
}

type Line struct {
	h      bool
	v      bool
	Point1 Point
	Point2 Point
}

func (l *Line) isStraight() bool {
	l.h = l.Point1.X == l.Point2.X
	l.v = l.Point1.Y == l.Point2.Y

	return l.h || l.v
}

func (l *Line) xDirection() int {
	if l.Point2.X-l.Point1.X > 0 {
		return 1
	}

	return -1
}

func (l *Line) yDirection() int {
	if l.Point2.Y-l.Point1.Y > 0 {
		return 1
	}

	return -1
}

var lines []Line

// grid is a two layer map, row first then column, then value
// this allows a sparse format
var grid map[int]map[int]int

var re regexp.Regexp = *regexp.MustCompile(`(\d+),(\d+) -> (\d+),(\d+)`)

func main() {
	grid = make(map[int]map[int]int)

	// Open puzzle input
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer input.Close()

	// Open new scanner to read line by line
	scanner := bufio.NewScanner(input)

	// read in lines
	for scanner.Scan() {
		row := scanner.Text()

		valstrs := re.FindStringSubmatch(row)
		// chop off whole string match
		valstrs = valstrs[1:]

		var valints [4]int

		for i, val := range valstrs {
			valints[i], err = strconv.Atoi(val)
			if err != nil {
				log.Fatal(err)
			}
		}

		line := Line{
			Point1: Point{
				X: valints[0],
				Y: valints[1],
			},
			Point2: Point{
				X: valints[2],
				Y: valints[3],
			},
		}

		if line.isStraight() {
			lines = append(lines, line)
		}
	}

	for _, line := range lines {
		xdirection := line.xDirection()
		ydirection := line.yDirection()

		if line.h {
			for y := line.Point1.Y; y != line.Point2.Y; y += ydirection {
				if _, ok := grid[line.Point1.X]; !ok {
					grid[line.Point1.X] = make(map[int]int)
				}
				grid[line.Point1.X][y]++
			}
		} else if line.v {
			for x := line.Point1.X; x != line.Point2.X; x += xdirection {
				if _, ok := grid[x]; !ok {
					grid[x] = make(map[int]int)
				}
				grid[x][line.Point1.Y]++
			}
		} else {
			for x := line.Point1.X; x != line.Point2.X; x += xdirection {
				for y := line.Point1.Y; y != line.Point2.Y; y += ydirection {
					if _, ok := grid[x]; !ok {
						grid[x] = make(map[int]int)
					}
					grid[x][y]++
				}
			}
		}

		// all three scenarios skip Point2
		if _, ok := grid[line.Point2.X]; !ok {
			grid[line.Point2.X] = make(map[int]int)
		}
		grid[line.Point2.X][line.Point2.Y]++
	}

	var count int

	for _, row := range grid {
		for _, cell := range row {
			if cell >= 2 {
				count++
			}
		}
	}

	fmt.Println(count)
}
