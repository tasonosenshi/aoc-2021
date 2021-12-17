package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	set "github.com/deckarep/golang-set"
)

func main() {
	signalList := make([][]string, 0)
	outputList := make([][]string, 0)

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

		signalPatterns := strings.Fields(rowParts[0])
		outputs := strings.Fields(rowParts[1])

		signalList = append(signalList, signalPatterns)
		outputList = append(outputList, outputs)
	}

	var outputSum int
	for i, signalpatterns := range signalList {
		segmentMap := map[rune]set.Set{
			'a': set.NewSet(),
			'b': set.NewSet(),
			'c': set.NewSet(),
			'd': set.NewSet(),
			'e': set.NewSet(),
			'f': set.NewSet(),
			'g': set.NewSet(),
		}
		digitMap := map[rune]set.Set{
			'0': set.NewSet(),
			'1': set.NewSet(),
			'2': set.NewSet(),
			'3': set.NewSet(),
			'4': set.NewSet(),
			'5': set.NewSet(),
			'6': set.NewSet(),
			'7': set.NewSet(),
			'8': set.NewSet(),
			'9': set.NewSet(),
		}
		segs5 := make([]set.Set, 0)

		for _, signalpattern := range signalpatterns {
			switch len(signalpattern) {
			case 2:
				for _, r := range signalpattern {
					_ = digitMap['1'].Add(r)
				}
			case 3:
				for _, r := range signalpattern {
					_ = digitMap['7'].Add(r)
				}
			case 4:
				for _, r := range signalpattern {
					_ = digitMap['4'].Add(r)
				}
			case 7:
				for _, r := range signalpattern {
					_ = digitMap['8'].Add(r)
				}
			case 5:
				newSet := set.NewSet()
				for _, r := range signalpattern {
					_ = newSet.Add(r)
				}
				segs5 = append(segs5, newSet)
			}
		}

		// 1, 4, 7, and 8 are found
		// a is the difference between 7 and 1
		segmentMap['a'] = digitMap['7'].Difference(digitMap['1'])

		// look through the 5 segments and find 3
		// 1 is a subset of 3
		for i, s := range segs5 {
			if digitMap['1'].IsProperSubset(s) {
				digitMap['3'] = s
				segs5 = append(segs5[:i], segs5[i+1:]...)
			}
		}

		// use 3 and 4 to get segment b
		segmentMap['b'] = digitMap['4'].Difference(digitMap['3'])

		// the difference between 5 and 3 is also b
		for i, s := range segs5 {
			if s.Difference(digitMap['3']).Equal(segmentMap['b']) {
				digitMap['5'] = s
				segs5 = append(segs5[:i], segs5[i+1:]...)
			}
		}

		// 2 remains
		digitMap['2'] = segs5[0]

		// difference between 2 and 3 is e
		segmentMap['e'] = digitMap['2'].Difference(digitMap['3'])

		// difference between 5 and (the difference between 3 and 1) and
		// the differnce between that and B is F
		segmentMap['f'] = digitMap['5'].
			Difference(digitMap['3'].Difference(digitMap['1'])).
			Difference(segmentMap['b'])

		// difference between 1 and f is c
		segmentMap['c'] = digitMap['1'].Difference(segmentMap['f'])

		// 4 - b - 1 == d
		segmentMap['d'] = digitMap['4'].Difference(segmentMap['b']).
			Difference(digitMap['1'])

		// 8 - all segments found == g
		segmentMap['g'] = digitMap['8'].
			Difference(segmentMap['a']).
			Difference(segmentMap['b']).
			Difference(segmentMap['c']).
			Difference(segmentMap['d']).
			Difference(segmentMap['e']).
			Difference(segmentMap['f'])

		// 0 is all segments but d
		digitMap['0'] = segmentMap['a'].
			Union(segmentMap['b']).
			Union(segmentMap['c']).
			Union(segmentMap['e']).
			Union(segmentMap['f']).
			Union(segmentMap['g'])

		// 6 is all segments but c
		digitMap['6'] = segmentMap['a'].
			Union(segmentMap['b']).
			Union(segmentMap['d']).
			Union(segmentMap['e']).
			Union(segmentMap['f']).
			Union(segmentMap['g'])

		// 9 is all segments but e
		digitMap['9'] = segmentMap['a'].
			Union(segmentMap['b']).
			Union(segmentMap['c']).
			Union(segmentMap['d']).
			Union(segmentMap['f']).
			Union(segmentMap['g'])

		var outputb strings.Builder
		for _, output := range outputList[i] {
			newSet := set.NewSet()
			for _, r := range output {
				_ = newSet.Add(r)
			}

			for digit, digitSet := range digitMap {
				if digitSet.Equal(newSet) {
					outputb.WriteRune(digit)
				}
			}
		}

		outputVal, err := strconv.Atoi(outputb.String())
		if err != nil {
			log.Fatal(err)
		}
		outputSum += outputVal
	}

	fmt.Println(outputSum)
}
