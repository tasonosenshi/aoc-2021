//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var heightmap [][]int

func main() {
	heightmap = make([][]int, 0)

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

		rowStrs := strings.Split(row, "")
		rowInts := make([]int, 0)
		for _, v := range rowStrs {
			i, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal(err)
			}
			rowInts = append(rowInts, i)
		}

		heightmap = append(heightmap, rowInts)
	}

	riskSum := 0
	for i, row := range heightmap {
		for j, v := range row {
			if i == 0 {
				if j == 0 {
					if v < heightmap[i+1][j] &&
						v < heightmap[i][j+1] {
						riskSum += v + 1
					}
				} else if j == len(heightmap[i])-1 {
					if v < heightmap[i+1][j] &&
						v < heightmap[i][j-1] {
						riskSum += v + 1
					}
				} else {
					if v < heightmap[i+1][j] &&
						v < heightmap[i][j-1] &&
						v < heightmap[i][j+1] {
						riskSum += v + 1
					}
				}
			} else if i == len(heightmap)-1 {
				if j == 0 {
					if v < heightmap[i-1][j] &&
						v < heightmap[i][j+1] {
						riskSum += v + 1
					}
				} else if j == len(heightmap[i])-1 {
					if v < heightmap[i-1][j] &&
						v < heightmap[i][j-1] {
						riskSum += v + 1
					}
				} else {
					if v < heightmap[i-1][j] &&
						v < heightmap[i][j-1] &&
						v < heightmap[i][j+1] {
						riskSum += v + 1
					}
				}
			} else {
				if j == 0 {
					if v < heightmap[i-1][j] &&
						v < heightmap[i+1][j] &&
						v < heightmap[i][j+1] {
						riskSum += v + 1
					}
				} else if j == len(heightmap[i])-1 {
					if v < heightmap[i-1][j] &&
						v < heightmap[i+1][j] &&
						v < heightmap[i][j-1] {
						riskSum += v + 1
					}
				} else {
					if v < heightmap[i-1][j] &&
						v < heightmap[i+1][j] &&
						v < heightmap[i][j-1] &&
						v < heightmap[i][j+1] {
						riskSum += v + 1
					}
				}
			}
		}
	}

	fmt.Println(riskSum)
}
