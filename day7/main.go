package main

import (
	"fmt"
	"io/ioutil"
	//"math"
	"strconv"
	"strings"
)

type ContBag struct {
	amount int
	name   string
}

type Bag struct {
	name     string
	contains []ContBag
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

func strToBag(input string) Bag {
	fmt.Printf("splitting '%s'\n", input)
	split := strings.Split(input, " bags contain ")
	name := split[0]
	var contains []ContBag

	if split[1] == "no other bags." {
		contains = []ContBag{}
	} else {
		bags := strings.Split(split[1], ", ")
		contains = make([]ContBag, 0, len(bags))
		for _, bag := range bags {
			words := strings.Split(bag, " ")
			colour := words[1] + " " + words[2]
			count, _ := strconv.ParseInt(words[0], 10, 64)
			contains = append(contains, ContBag{int(count), colour})
		}
	}

	return Bag{name, contains}
}

func pt1() {
	inputs := strings.Split(readInput("input"), "\n")
	// parse lines to list of bags first
	bags := make([]Bag, 0, 1000)
	for _, input := range inputs {
		if len(input) != 0 {
			bags = append(bags, strToBag(input))
		}
	}

	for _, bag := range bags {
		fmt.Print(bag)
		fmt.Print("\n")
	}

	// scan bags, adding bags that contain shinygold
	goodBags := make(map[string]bool)
	goodBags["shiny gold"] = true
	goodBagsLen := 0
	for goodBagsLen != len(goodBags) {
		goodBagsLen = len(goodBags)
		fmt.Printf("there's %v goodBags\n", len(goodBags))
		for _, bag := range bags {
			for _, contbag := range bag.contains {
				if goodBags[contbag.name] {
					goodBags[bag.name] = true
				}
			}
		}
	}
	fmt.Printf("\n%v\n there's %v goodBags", goodBags, goodBagsLen)
}

func pt2() {
	inputs := strings.Split(readInput("input"), "\n")
	// parse lines to list of bags first
	bags := make(map[string]Bag)
	for _, input := range inputs {
		if len(input) != 0 {
			bag := strToBag(input)
			bags[bag.name] = bag
		}
	}

	fmt.Print(bags)

	// scan bags, adding bags that contain shinygold
	nextBags := make([]string, 0, 32)
	nextBags = append(nextBags, "shiny gold")
	total := 0

	for len(nextBags) != 0 {
		fmt.Printf("opening bags: \n%v\n", nextBags)
		// add bags to new slice
		nNextBags := make([]string, 0, 32)
		for _, bag := range nextBags {
			for _, nestedBag := range bags[bag].contains {
				for i := 0; i < nestedBag.amount; i++ {
					total++
					nNextBags = append(nNextBags, nestedBag.name)
				}
			}
		}
		// assing nextBags to newly built slice
		nextBags = nNextBags
	}
	fmt.Printf("total bags: %v", total)
}

func main() {
	pt1()
	pt2()
}
