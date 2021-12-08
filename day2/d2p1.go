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

type Position struct {
	Horizontal int
	Depth      int
}

func main() {
	input, err := os.Open("input.txt")
	defer input.Close()

	if err != nil {
		log.Fatal(err)
	}

	pos := Position{}

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		i, err := strconv.Atoi(fields[1])
		if err != nil {
			log.Fatal(err)
		}

		switch fields[0] {
		case "forward":
			pos.Horizontal += i
		case "up":
			pos.Depth -= i
		case "down":
			pos.Depth += i
		}
	}

	fmt.Println("Horizontal is: ", pos.Horizontal)
	fmt.Println("Depth is: ", pos.Depth)
	fmt.Println("Multiplied together: ", pos.Horizontal*pos.Depth)

}
