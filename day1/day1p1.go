package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	input, err := os.Open("input.txt")
	defer input.Close()

	if err != nil {
		log.Fatal(err)
	}

	var prev int
	var count int

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		cur, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		if cur > prev {
			count++
		}

		prev = cur
	}

	// First comparision will always increment count, but first comparison shouldn't count
	count--

	fmt.Println(count)
}
