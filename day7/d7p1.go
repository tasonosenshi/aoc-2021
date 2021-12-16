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

var crabs []int
var crabMap map[int]int

func main() {
	crabs = make([]int, 0)
	crabMap = make(map[int]int)

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

		crabStrs := strings.Split(row, ",")

		for _, crab := range crabStrs {
			crabInt, err := strconv.Atoi(crab)
			if err != nil {
				log.Fatal(err)
			}

			crabs = append(crabs, crabInt)
		}
	}

	for _, v := range crabs {
		sum := 0

		for _, s := range crabs {
			val := s - v
			if val < 0 {
				val = -val
			}
			sum += val
		}

		crabMap[v] = sum
	}

	min := 0

	for _, v := range crabMap {
		if min == 0 {
			min = v
		} else if min > v {
			min = v
		}
	}

	fmt.Println(min)
}
