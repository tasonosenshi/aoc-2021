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

var lanternfishMap map[int]int

func main() {
	lanternfish = make([]int, 0)
	lanternfishMap = map[int]int{
		0: 0,
		1: 0,
		2: 0,
		3: 0,
		4: 0,
		5: 0,
		6: 0,
		7: 0,
		8: 0,
	}

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

	for _, v := range lanternfish {
		lanternfishMap[v]++
	}

	for i := 1; i <= 256; i++ {
		lanternfishNextDayMap := map[int]int{
			0: lanternfishMap[1],
			1: lanternfishMap[2],
			2: lanternfishMap[3],
			3: lanternfishMap[4],
			4: lanternfishMap[5],
			5: lanternfishMap[6],
			6: lanternfishMap[7] + lanternfishMap[0],
			7: lanternfishMap[8],
			8: lanternfishMap[0],
		}

		lanternfishMap = lanternfishNextDayMap
	}

	totalFish := 0

	for _, v := range lanternfishMap {
		totalFish += v
	}

	fmt.Println(totalFish)
}
