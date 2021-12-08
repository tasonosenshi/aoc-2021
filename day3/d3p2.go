package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type bitTree struct {
	root *bitTreeNode
}

type bitTreeNode struct {
	count int
	zero  *bitTreeNode
	one   *bitTreeNode
}

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer input.Close()

	t := bitTree{
		root: &bitTreeNode{
			zero: &bitTreeNode{},
			one:  &bitTreeNode{},
		},
	}

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		bitfield := scanner.Text()

		bits := strings.Split(bitfield, "")

		// build the tree
		// start at root
		curNode := t.root

		for _, bit := range bits {
			switch bit {
			case "0":
				curNode = curNode.zero
			case "1":
				curNode = curNode.one
			}

			// if curNode. zero is nil, this is an uninitialized node
			if curNode.zero == nil {
				curNode.zero = &bitTreeNode{}
				curNode.one = &bitTreeNode{}
			}

			curNode.count++
		}
	}

	var o strings.Builder
	curNode := t.root
	for {
		if curNode.zero.count == 0 {
			if curNode.one.count == 0 {
				break
			} else {
				o.WriteRune('1')
				curNode = curNode.one
				continue
			}
		} else if curNode.one.count == 0 {
			o.WriteRune('0')
			curNode = curNode.zero
			continue
		}

		if curNode.zero.count <= curNode.one.count {
			o.WriteRune('1')
			curNode = curNode.one
		} else {
			o.WriteRune('0')
			curNode = curNode.zero
		}
	}

	oxygen, err := strconv.ParseUint(o.String(), 2, 12)
	if err != nil {
		log.Fatal(err)
	}

	var c strings.Builder
	curNode = t.root
	for {
		fmt.Println(curNode)
		if curNode.zero.count == 0 {
			if curNode.one.count == 0 {
				break
			} else {
				c.WriteRune('1')
				curNode = curNode.one
				continue
			}
		} else if curNode.one.count == 0 {
			c.WriteRune('0')
			curNode = curNode.zero
			continue
		}

		if curNode.zero.count > curNode.one.count {
			c.WriteRune('1')
			curNode = curNode.one
		} else {
			c.WriteRune('0')
			curNode = curNode.zero
		}
	}

	co2, err := strconv.ParseUint(c.String(), 2, 12)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(oxygen, co2, oxygen*co2)
}
