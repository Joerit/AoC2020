package main

import (
	"fmt"
	"io/ioutil"
	//"math"
	//"math/big"
	"sort"
	"strconv"
	"strings"
	"time"

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
	fmt.Printf("d1: %v, d3: %v\n", d1, d3)
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

func topRecursiveScan(inputs []int, drop *bit.Set, index int, ch chan int) {
	if !isValid(inputs, drop) {
		ch <- 0
		return
	}
	if index >= len(inputs)-1 {
		ch <- 1
		return
	}
	ret := recursiveScan(inputs, drop.Add(index+1), index+1)
	ret += recursiveScan(inputs, drop.Delete(index+1), index+1)
	ch <- ret
}

func recursiveScan(inputs []int, drop *bit.Set, index int) int {
	if !isValid(inputs, drop) {
		return 0
	}
	if index >= len(inputs)-1 {
		return 1
	}
	ret := recursiveScan(inputs, drop.Add(index+1), index+1)
	ret += recursiveScan(inputs, drop.Delete(index+1), index+1)
	return ret
}

func pt2v2(inputs []int) {
	ch := make(chan int)

	// split on steps of 3
	start := 0
	splits := 0
	for end := 1; end <= len(inputs)-2; end++ {
		if inputs[end+1]-inputs[end] == 3 {
			go topRecursiveScan(inputs[start:end+1], bit.New(), 0, ch)
			splits++
			start = end + 1
		}
	}

	//fmt.Println("split ", splits, " times")
	total := 1
	for ; splits > 0; splits-- {
		total *= <-ch
		//fmt.Println("running total: ", total)
	}
	fmt.Println("total: ", total)
}

func pt2v2serial(inputs []int) {
	// split on steps of 3
	start := 0
	total := 1

	for end := 1; end <= len(inputs)-2; end++ {
		if inputs[end+1]-inputs[end] == 3 {
			total *= recursiveScan(inputs[start:end+1], bit.New(), 0)
			start = end + 1
		}
	}
	fmt.Println("total: ", total)
}

func main() {
	input := parseString()
	fmt.Println(len(input))
	input = pt1(input)

	fmt.Println("parallel")
	start := time.Now()
	pt2v2(input)
	duration := time.Since(start)
	fmt.Println("start: ", start, " since: ", duration)

	fmt.Println("serial")
	start = time.Now()
	pt2v2serial(input)
	duration = time.Since(start)
	fmt.Println("start: ", start, " since: ", duration)
}
