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
	set := bit.New().AddRange('a', 'z'+1)
	for _, line := range lines {
		if len(line) == 0 {
			break
		}

		lineSet := bit.New()
		for _, char := range line {
			lineSet.Add(int(char))
		}
		fmt.Print("\nlineset:\n")
		fmt.Print(lineSet)
		set = set.And(lineSet)
		fmt.Print("\nset:\n")
		fmt.Print(set)
	}
	fmt.Print("\n")
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
