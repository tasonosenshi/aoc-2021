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
	Aim        int
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
			pos.Depth += pos.Aim * i
		case "up":
			pos.Aim -= i
		case "down":
			pos.Aim += i
		}
	}

	fmt.Println("Horizontal is: ", pos.Horizontal)
	fmt.Println("Depth is: ", pos.Depth)
	fmt.Println("Multiplied together: ", pos.Horizontal*pos.Depth)

}
