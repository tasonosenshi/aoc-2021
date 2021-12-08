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

	var prevWindowValue int
	var count int

	window := make([]int, 3)

	scanner := bufio.NewScanner(input)
	for i := 0; scanner.Scan(); i++ {
		cur, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		window[i%3] = cur

		if i < 2 {
			continue
		}

		curWindowValue := window[0] + window[1] + window[2]
		fmt.Println(curWindowValue)

		if curWindowValue > prevWindowValue {
			count++
		}

		prevWindowValue = curWindowValue
	}

	// First comparision will always increment count, but first comparison shouldn't count
	count--

	fmt.Println(count)
}
