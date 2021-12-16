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

var lanternfish []int

func main() {
	lanternfish = make([]int, 0)

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

		fishStrs := strings.Split(row, ",")

		for _, fish := range fishStrs {
			fishInt, err := strconv.Atoi(fish)
			if err != nil {
				log.Fatal(err)
			}

			lanternfish = append(lanternfish, fishInt)
		}
	}

	fmt.Println("Initial state: ", lanternfish)

	for i := 1; i <= 80; i++ {
		newFishCount := 0

		for i := range lanternfish {
			lanternfish[i]--

			if lanternfish[i] < 0 {
				lanternfish[i] = 6
				newFishCount++
			}
		}

		newFish := make([]int, newFishCount)
		for i := range newFish {
			newFish[i] = 8
		}

		lanternfish = append(lanternfish, newFish...)

		if i < 15 {
			fmt.Println("After ", i, " days: ", lanternfish)
		}
	}

	fmt.Println(len(lanternfish))
}
