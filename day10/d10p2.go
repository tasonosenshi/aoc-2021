package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

var points = map[rune]int{
	')': 1,
	']': 2,
	'}': 3,
	'>': 4,
}

var pairOpen = map[rune]rune{
	')': '(',
	']': '[',
	'}': '{',
	'>': '<',
}

var pairClose = map[rune]rune{
	'(': ')',
	'[': ']',
	'{': '}',
	'<': '>',
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

	incompleteStacks := make([][]rune, 0)

	// read in lines
	for scanner.Scan() {
		row := scanner.Text()

		openStack := make([]rune, 0)
		isCorrupted := false
	row:
		for _, r := range row {
			switch r {
			case '(', '[', '{', '<':
				openStack = append(openStack, r)
			case ')', ']', '}', '>':
				if pairOpen[r] != openStack[len(openStack)-1] {
					isCorrupted = true
					break row
				}

				openStack = openStack[:len(openStack)-1]
			}
		}

		if !isCorrupted {
			incompleteStacks = append(incompleteStacks, openStack)
		}
	}

	scores := make([]int, 0)
	for _, stack := range incompleteStacks {
		syntaxScore := 0
		for len(stack) != 0 {
			r := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			syntaxScore = syntaxScore*5 + points[pairClose[r]]
		}

		scores = append(scores, syntaxScore)
	}

	sort.Ints(scores)

	fmt.Println(scores[len(scores)/2])
}
