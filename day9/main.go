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

func parseString() []int {
	inputs := strings.Split(readInput("input"), "\n")
	nums := make([]int, 0, 64)
	for i := 0; i < len(inputs); i++ {
		next, _ := strconv.ParseInt(inputs[i], 10, 64)
		nums = append(nums, int(next))
	}
	return nums
}

func pt1(inputs []int) int {
	nums := [25]int{}
	// initialise first 25 numbers
	for i := 0; i < 25; i++ {
		nums[i] = inputs[i]
	}

	first := 0
	for i := 25; i < len(inputs); i++ {
		next := inputs[i]
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
			return next
		}
	}
	panic("i'm expecting to find a number")
}

func pt2(inputs []int, target int) {
	for i := 0; i < len(inputs); i++ {
		acc := 0
		j := i
		for acc < target {
			acc += inputs[j]
			j++
		}
		if acc == target {
			fmt.Printf("target %v found (%v), summing\n", target, acc)
			min := target
			max := 0
			for ; i <= j; i++ {
				fmt.Printf("%v ", inputs[i])
				if inputs[i] < min {
					min = inputs[i]
				}
				if inputs[i] > max {
					max = inputs[i]
				}

			}
			fmt.Printf("min: %v, max: %v, add: %v", min, max, min+max)
			return
		}
	}
}

func main() {
	input := parseString()
	target := pt1(input)
	pt2(input, target)
}
