package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

type Seat struct {
	row    float64
	column float64
	id     float64
}

func seat(row float64, column float64) Seat {
	return Seat{row, column, row*8 + column}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readInput(filename string) string {
	dat, err := ioutil.ReadFile(filename)
	check(err)
	return string(dat)
}

func parseSeat(seatStr string) Seat {
	row := 0.0
	column := 0.0
	for i, letter := range seatStr[:7] {
		switch letter {
		case 'F':
			continue
		case 'B':
			row += 64 / math.Pow(2, float64(i))
		}
	}
	for i, letter := range seatStr[7:] {
		switch letter {
		case 'L':
		case 'R':
			column += 4 / math.Pow(2, float64(i))
		}
	}
	return seat(row, column)
}

func main() {
	input := strings.Split(readInput("input"), "\n")
	fmt.Print(input)
	seats := [1000]Seat{}
	for _, str := range input {
		if len(str) > 0 {
			seat := parseSeat(str)
			seats[int(seat.id)] = seat
		}
	}
	fmt.Print(seats)
	found := false
	for i, seat := range seats {
		if found {
			if seat.id == 0.0 {
				fmt.Print("\n")
				fmt.Print(i)
				return
			}
		} else {
			if seat.id != 0.0 {
				found = true
			}
		}
	}
}
