package main

import (
	"fmt"
	"io/ioutil"
	//"math"
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

func parse(str string) *bit.Set {
	lines := strings.Split(str, "\n")
	set := bit.New()
	for _, line := range lines {
		for _, char := range line {
			set.Add(int(char))
		}
	}
	return set
}

func main() {
	inputs := strings.Split(readInput("input"), "\n\n")
	total := 0
	for _, input := range inputs {
		total += parse(input).Size()
	}
	fmt.Printf("total %d", total)

}
