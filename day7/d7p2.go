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

	max := 0
	for _, v := range crabs {
		if max < v {
			max = v
		}
	}

	// for each step from zero to the max crab distance
	for i := 0; i <= max; i++ {
		sum := 0

		for _, v := range crabs {
			// measure the distance between the step and each crab
			val := i - v
			if val < 0 {
				val = -val
			}

			// for every step in difference between a crab and the spot
			// add 1 more fuel cost than the last step
			for i := 1; i <= val; i++ {
				sum += i
			}
		}

		crabMap[i] = sum
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
