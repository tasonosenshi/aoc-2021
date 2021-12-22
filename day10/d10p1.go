//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var syntaxScore int
var points = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

var pairOpen = map[rune]rune{
	')': '(',
	']': '[',
	'}': '{',
	'>': '<',
}

func main() {
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

		openStack := make([]rune, 0)
	row:
		for _, r := range row {
			switch r {
			case '(', '[', '{', '<':
				openStack = append(openStack, r)
			case ')', ']', '}', '>':
				if pairOpen[r] != openStack[len(openStack)-1] {
					syntaxScore += points[r]
					break row
				}

				openStack = openStack[:len(openStack)-1]
			}
		}
	}

	fmt.Println(syntaxScore)
}
