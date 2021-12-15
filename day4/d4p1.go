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

var bingoCage []int
var bingoRoll int

var bingoBoards []bingoBoard

type bingoBoard struct {
	board [5][5]int
}

func (b bingoBoard) hasBingo(rolls []int) bool {
	markerCount := 0

	for row := 0; row < 5; row++ {
		for _, roll := range rolls {
			if b.rowContains(row, roll) {
				markerCount++
			}

			if markerCount == 5 {
				return true
			}
		}

		markerCount = 0
	}

	for col := 0; col < 5; col++ {
		for _, roll := range rolls {
			if b.colContains(col, roll) {
				markerCount++
			}

			if markerCount == 5 {
				return true
			}
		}

		markerCount = 0
	}

	return false
}

func (b bingoBoard) rowContains(row int, val int) bool {
	for _, rv := range b.board[row] {
		if rv == val {
			return true
		}
	}

	return false
}

func (b bingoBoard) colContains(col int, val int) bool {
	colT := []int{
		b.board[0][col],
		b.board[1][col],
		b.board[2][col],
		b.board[3][col],
		b.board[4][col],
	}

	for _, rv := range colT {
		if rv == val {
			return true
		}
	}

	return false
}

func (b bingoBoard) boardScore(rolls []int) int {
	score := 0

	for row := 0; row < 5; row++ {
		for col := 0; col < 5; col++ {
			score += b.board[row][col]

			for _, roll := range rolls {
				if b.board[row][col] == roll {
					score -= b.board[row][col]
					break
				}
			}
		}
	}

	score *= rolls[len(rolls)-1]

	return score
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
	// Bingo Cage line
	scanner.Scan()
	bingoCageTxt := scanner.Text()
	bingoCageTxtArr := strings.Split(bingoCageTxt, ",")

	bingoCage = make([]int, len(bingoCageTxtArr))

	// Populate bingo cage
	for i, bingoCageVal := range bingoCageTxtArr {
		bingoCage[i], err = strconv.Atoi(bingoCageVal)
		if err != nil {
			log.Fatal(err)
		}
	}

	// newline after Bingo Cage
	scanner.Scan()
	isEndOfPrevBoard := true
	boardRow := 0

	// read in bingo boards
	for scanner.Scan() {
		row := scanner.Text()
		if isEndOfPrevBoard {
			isEndOfPrevBoard = false
			bingoBoards = append(bingoBoards, bingoBoard{})
			boardRow = 0
		} else if row == "" {
			isEndOfPrevBoard = true
			continue
		}

		rowStrs := strings.Fields(row)
		for i, str := range rowStrs {
			bingoBoards[len(bingoBoards)-1].board[boardRow][i], err = strconv.Atoi(str)
			if err != nil {
				log.Fatal(err)
			}
		}

		boardRow++
	}

	// need at least 5 rolls to have a bingo, so skip ahead to 5
	bingoRoll = 5

	var boardWithBingo *bingoBoard

	// Does any board have bingo?  If not, get another ball and try again.
	for boardWithBingo == nil {
		for _, b := range bingoBoards {
			if b.hasBingo(bingoCage[0:bingoRoll]) {
				boardWithBingo = &b
				break
			}
		}

		bingoRoll++
	}

	// bingoroll incremented after finding bingo board, so decrement to fix
	bingoRoll--

	fmt.Println(boardWithBingo.boardScore(bingoCage[0:bingoRoll]))
}
