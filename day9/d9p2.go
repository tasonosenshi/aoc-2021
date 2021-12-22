package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var heightmap [][]Point

var lowPoints PointSet

var basins []PointSet

// +gen set
type Point struct {
	X     int
	Y     int
	Value int
}

func (p *Point) north() *Point {
	if p.Y == 0 {
		return nil
	}

	return &heightmap[p.Y-1][p.X]
}

func (p *Point) east() *Point {
	if p.X == len(heightmap[p.Y])-1 {
		return nil
	}

	return &heightmap[p.Y][p.X+1]
}

func (p *Point) south() *Point {
	if p.Y == len(heightmap)-1 {
		return nil
	}

	return &heightmap[p.Y+1][p.X]
}

func (p *Point) west() *Point {
	if p.X == 0 {
		return nil
	}

	return &heightmap[p.Y][p.X-1]
}

func (p *Point) isLessThanNorth() bool {
	if p.north() == nil {
		return true
	}

	return p.Value < p.north().Value
}

func (p *Point) isLessThanEast() bool {
	if p.east() == nil {
		return true
	}

	return p.Value < p.east().Value
}

func (p *Point) isLessThanSouth() bool {
	if p.south() == nil {
		return true
	}

	return p.Value < p.south().Value
}

func (p *Point) isLessThanWest() bool {
	if p.west() == nil {
		return true
	}

	return p.Value < p.west().Value
}

func main() {
	heightmap = make([][]Point, 0)
	lowPoints = NewPointSet()

	// Open puzzle input
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer input.Close()

	// Open new scanner to read line by line
	scanner := bufio.NewScanner(input)

	// read in lines
	var y int
	for scanner.Scan() {
		row := scanner.Text()

		rowStrs := strings.Split(row, "")
		rowPoints := make([]Point, 0)
		for x, valstr := range rowStrs {
			v, err := strconv.Atoi(valstr)
			if err != nil {
				log.Fatal(err)
			}
			rowPoints = append(rowPoints, Point{
				X:     x,
				Y:     y,
				Value: v,
			})
		}

		heightmap = append(heightmap, rowPoints)
		y++
	}

	// find and add low points
	for _, row := range heightmap {
		for _, p := range row {
			if p.isLessThanNorth() &&
				p.isLessThanEast() &&
				p.isLessThanSouth() &&
				p.isLessThanWest() {
				lowPoints.Add(p)
			}
		}
	}

	// flood fill basins
	basins = make([]PointSet, 0)
	for p := range lowPoints {
		basins = append(basins, NewPointSet(p))
	}

	var queue []Point
	for _, basin := range basins {
		// basin only has low point to start with
		// start by adding surrounding points to queue
		p := basin.ToSlice()[0]
		if p.north() != nil {
			queue = append(queue, *p.north())
		}

		if p.east() != nil {
			queue = append(queue, *p.east())
		}

		if p.south() != nil {
			queue = append(queue, *p.south())
		}

		if p.west() != nil {
			queue = append(queue, *p.west())
		}

		for len(queue) != 0 {
			p = queue[0]
			queue = queue[1:]

			if p.Value != 9 && !basin.Contains(p) {
				basin.Add(p)

				if p.north() != nil {
					queue = append(queue, *p.north())
				}

				if p.east() != nil {
					queue = append(queue, *p.east())
				}

				if p.south() != nil {
					queue = append(queue, *p.south())
				}

				if p.west() != nil {
					queue = append(queue, *p.west())
				}
			}
		}
	}

	var largestThreeSizes []int
	//find three largest basins
	for i := 0; i < 3; i++ {
		largestIdx := 0
		largestValue := 0

		for j, basin := range basins {
			if basin.Cardinality() > largestValue {
				largestIdx = j
				largestValue = basin.Cardinality()
			}
		}

		largestThreeSizes = append(largestThreeSizes, largestValue)
		basins = append(basins[:largestIdx], basins[largestIdx+1:]...)
	}

	fmt.Println(largestThreeSizes[0] * largestThreeSizes[1] * largestThreeSizes[2])
}
