package main

import (
	"fmt"
	"io/ioutil"
	//"math"
	"sort"
	"strconv"
	"strings"

	"github.com/yourbasic/bit"
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
	var max int64 = 0
	for i := 0; i < len(inputs); i++ {
		if len(inputs[i]) == 0 {
			continue
		}
		next, _ := strconv.ParseInt(inputs[i], 10, 64)
		if next > max {
			max = next
		}
		nums = append(nums, int(next))
	}
	nums = append(nums, 0)
	nums = append(nums, int(max)+3)
	return nums
}

func pt1(inputs []int) []int {
	sort.Ints(inputs)
	fmt.Printf("sorted list: %v\n", inputs)
	d1 := 0
	d3 := 0

	for i := 0; i < len(inputs)-1; i++ {
		switch inputs[i+1] - inputs[i] {
		case 1:
			d1++
		case 3:
			d3++
		}
	}
	fmt.Printf("d1: %v, d3: %v", d1, d3)
	return inputs
}

// check if inputs is still a valid series when dropping indices in 'drop'
func isValid(inputs []int, drop *bit.Set) (valid bool) {
	for i := 0; i < len(inputs)-1; i++ {
		if drop.Contains(i) {
			continue
		}
		toCheck := i + 1
		for drop.Contains(toCheck) {
			toCheck++
		}
		// if we go out of bounds, return false
		if toCheck >= len(inputs) {
			return false
		}
		if inputs[toCheck]-inputs[i] > 3 {
			return false
		}
	}
	return true
}

func drops3(drop *bit.Set, inputsLen int) bool {
	for i := 0; i < inputsLen-5; i++ {
		if drop.Contains(i) && drop.Contains(i+1) && drop.Contains(i+2) {
			return true
		}
	}
	return false
}

func topRecursiveScan(inputs []int, drop *bit.Set, index int, ch chan int) {
	ret := 0
	ret += recursiveScan(inputs, drop.Add(index), index+1)
	ret += recursiveScan(inputs, drop.Delete(index), index+1)
	ch <- ret
}

func recursiveScan(inputs []int, drop *bit.Set, index int) int {
	ret := 0
	if index == len(inputs) {
		if isValid(inputs, drop.Add(index)) {
			ret++
		}
		if isValid(inputs, drop.Delete(index)) {
			ret++
		}
		return ret
	}
	if drops3(drop, len(inputs)) {
		return ret
	}
	ret += recursiveScan(inputs, drop.Add(index), index+1)
	ret += recursiveScan(inputs, drop.Delete(index), index+1)
	return ret
}

func pt2(inputs []int) {
	ch := make(chan int)
	for i := 0; i < 16; i++ {
		drop := bit.New()
		tmp := i
		for j := 1; j < 5; j++ {
			if tmp%2 == 0 {
				drop.Add(j)
			}
			tmp /= 2
		}
		fmt.Println("starting with bitset: ", drop)
		go topRecursiveScan(inputs, drop, 5, ch)
	}
	total := 0
	for i := 0; i < 16; i++ {
		total += <-ch
		fmt.Println("running total ", total)
	}

	fmt.Printf("outcome: %v", total)
}

func main() {
	input := parseString()
	fmt.Println(len(input))
	input = pt1(input)
	pt2(input)
}
