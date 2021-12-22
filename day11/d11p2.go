package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Octopus struct {
	X       int
	Y       int
	Energy  int
	Flashed bool
}

func (o *Octopus) North() *Octopus {
	if o == nil || o.Y == 0 {
		return nil
	}

	return octopuses[o.Y-1][o.X]
}

func (o *Octopus) South() *Octopus {
	if o == nil || o.Y == len(octopuses)-1 {
		return nil
	}

	return octopuses[o.Y+1][o.X]
}

func (o *Octopus) West() *Octopus {
	if o == nil || o.X == 0 {
		return nil
	}

	return octopuses[o.Y][o.X-1]
}

func (o *Octopus) East() *Octopus {
	if o == nil || o.X == len(octopuses[0])-1 {
		return nil
	}

	return octopuses[o.Y][o.X+1]
}

func (o *Octopus) NorthWest() *Octopus {
	return o.North().West()
}

func (o *Octopus) NorthEast() *Octopus {
	return o.North().East()
}

func (o *Octopus) SouthWest() *Octopus {
	return o.South().West()
}

func (o *Octopus) SouthEast() *Octopus {
	return o.South().East()
}

func (o *Octopus) EnergyInc() {
	if o == nil {
		return
	}

	o.Energy++
}

func (o *Octopus) AttemptFlash() {
	if o == nil || o.Flashed || o.Energy <= 9 {
		return
	}

	o.Flashed = true
	flashCount++

	o.NorthWest().EnergyInc()
	o.NorthWest().AttemptFlash()

	o.North().EnergyInc()
	o.North().AttemptFlash()

	o.NorthEast().EnergyInc()
	o.NorthEast().AttemptFlash()

	o.West().EnergyInc()
	o.West().AttemptFlash()

	o.East().EnergyInc()
	o.East().AttemptFlash()

	o.SouthWest().EnergyInc()
	o.SouthWest().AttemptFlash()

	o.South().EnergyInc()
	o.South().AttemptFlash()

	o.SouthEast().EnergyInc()
	o.SouthEast().AttemptFlash()
}

func (o *Octopus) ResetEnergy() {
	if o == nil || !o.Flashed {
		return
	}

	o.Flashed = false
	o.Energy = 0
}

var octopuses [][]*Octopus

var flashCount int

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

	rowc := 0
	for scanner.Scan() {
		row := scanner.Text()

		rowStrs := strings.Split(row, "")
		octopusesRow := make([]*Octopus, 0)
		for x, vstr := range rowStrs {
			v, err := strconv.Atoi(vstr)
			if err != nil {
				log.Fatal(err)
			}

			octopusesRow = append(octopusesRow, &Octopus{
				X:      x,
				Y:      rowc,
				Energy: v,
			})
		}

		octopuses = append(octopuses, octopusesRow)

		rowc++
	}

	stepCount := 1
	for {
		flashCount = 0
		for _, row := range octopuses {
			for _, octopus := range row {
				octopus.EnergyInc()
			}
		}

		for _, row := range octopuses {
			for _, octopus := range row {
				octopus.AttemptFlash()
			}
		}

		for _, row := range octopuses {
			for _, octopus := range row {
				octopus.ResetEnergy()
			}
		}

		if flashCount == 100 {
			break
		}

		stepCount++
	}

	fmt.Println(stepCount)
}
