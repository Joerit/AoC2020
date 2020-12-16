package main

import (
	"fmt"
	"io/ioutil"
	//"math"
	"strconv"
	"strings"
)

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

func pt1() {
	inputs := strings.Split(readInput("input"), "\n")
	nums := [25]int{}
	// initialise first 25 numbers
	for i := 0; i < 25; i++ {
		next, _ := strconv.ParseInt(inputs[i], 10, 64)
		nums[i] = int(next)
	}

	first := 0
	for i := 25; i < len(inputs); i++ {
		parsed, _ := strconv.ParseInt(inputs[i], 10, 64)
		next := int(parsed)
		// find valid sum
		isValid := false
		for _, x := range nums {
			for _, y := range nums {
				if x+y == next {
					isValid = true
					break
				}
				if isValid {
					break
				}
			}
		}
		if isValid {
			nums[first] = next
			first++
			if first >= len(nums) {
				first = 0
			}
		} else {
			fmt.Printf("%v is not valid\n", next)
			return
		}
	}
}

func main() {
	pt1()
}
