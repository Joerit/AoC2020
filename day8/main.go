package main

import (
	"fmt"
	"io/ioutil"
	//"math"
	"strconv"
	"strings"

	"github.com/yourbasic/bit"
)

/////////////////////////////
// Xomputor stuff

type Xomputor struct {
	acc                 int
	pointer             int
	instructions        []Instruction
	visitedInstructions *bit.Set
}

// build fresh
func xomputor() Xomputor {
	return Xomputor{
		0, 0,
		make([]Instruction, 0, 16),
		new(bit.Set),
	}
}

// build copy
func (x *Xomputor) copyXomputor() Xomputor {
	inst := make([]Instruction, len(x.instructions))
	copy(inst, x.instructions)
	return Xomputor{0, 0, inst, new(bit.Set)}
}

func (x *Xomputor) run() *Xomputor {
	for {
		if x.visitedInstructions.Contains(x.pointer) || x.pointer >= len(x.instructions) {
			return x
		}
		x.visitedInstructions.Add(x.pointer)

		switch x.instructions[x.pointer].name {
		case nop:
		case jmp:
			x.pointer += x.instructions[x.pointer].arg - 1
		case acc:
			x.acc += x.instructions[x.pointer].arg
		default:
			panic("Unknown instruction")
		}
		x.pointer++
	}
	return x
}

// enum for instruction names
type instructionName int

const (
	nop instructionName = iota
	jmp
	acc
)

type Instruction struct {
	name instructionName
	arg  int
}

func parse(input string) Xomputor {
	x := xomputor()
	for _, line := range strings.Split(input, "\n") {
		if len(line) == 0 {
			continue
		}
		x.instructions = append(x.instructions, strToInstr(line))
	}
	return x
}

func strToInstr(input string) Instruction {
	fmt.Printf("splitting '%s'\n", input)
	split := strings.Split(input, " ")
	var name instructionName

	switch split[0] {
	case "nop":
		name = nop
	case "jmp":
		name = jmp
	case "acc":
		name = acc
	}

	arg, _ := strconv.ParseInt(split[1], 10, 64)

	return Instruction{name, int(arg)}
}

// Xomputor end
//////////////////////////////
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
	x := parse(readInput("input"))
	// parse lines to list of bags first
	fmt.Printf("%v\n", x)
	x.run()
	fmt.Printf("%v\n", x)
}

func pt2() {
	x := parse(readInput("input"))
	// parse lines to list of bags first
	ch := make(chan Xomputor)
	for i := 0; i < len(x.instructions); i++ {
		go checkMod(&x, i, ch)
	}
	fmt.Printf("done, now waiting\n")
	fixedX := <-ch
	fmt.Printf("%v\n", fixedX)
}

func checkMod(x *Xomputor, i int, ch chan Xomputor) {
	newX := x.copyXomputor()
	switch newX.instructions[i].name {
	case nop:
		newX.instructions[i].name = jmp
	case jmp:
		newX.instructions[i].name = nop
	default:
		return
	}

	newX.run()

	if newX.pointer == len(newX.instructions) {
		fmt.Printf("found working X\n")
		ch <- newX
		close(ch)
	}
	return
}

func main() {
	pt1()
	pt2()
}
