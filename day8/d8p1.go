//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var outputList []string

func main() {
	outputList = make([]string, 0)

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

		rowParts := strings.Split(row, " | ")
		outputs := strings.Fields(rowParts[1])

		outputList = append(outputList, outputs...)
	}

	var count int
	for _, output := range outputList {
		segs := len(output)
		if segs == 2 || segs == 4 || segs == 3 || segs == 7 {
			count++
		}
	}

	fmt.Println(count)
}
