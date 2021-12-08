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

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer input.Close()

	bitCounts := make([]int, 12)

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		bitfield := scanner.Text()

		bits := strings.Split(bitfield, "")

		for i, bit := range bits {
			switch bit {
			case "1":
				bitCounts[i]++
			case "0":
				bitCounts[i]--
			}
		}
	}

	var b strings.Builder
	for i, bit := range bitCounts {
		fmt.Println(i, " in bitCounts")

		if bit < 0 {
			b.WriteRune('0')
		} else {
			b.WriteRune('1')
		}
	}

	gamma, err := strconv.ParseUint(b.String(), 2, 12)
	if err != nil {
		log.Fatal(err)
	}

	//bitmask for only the lower 12 bits
	epsilon := ^gamma & 0b000111111111111

	fmt.Println(gamma, epsilon, gamma*epsilon)
}
