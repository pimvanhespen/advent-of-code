package main

import (
	"fmt"
	"github.com/pimvanhespen/advent-of-code/pkg/aoc"
	"io"
	"log"
	"strconv"
	"strings"
)

type Input = [][]string

func main() {
	event := aoc.New(2016, 23, parse)
	fmt.Println("1:", aoc.Must(event.Run(part1)))
	fmt.Println("2:", aoc.Must(event.Run(part2)))
}

func parse(r io.Reader) (Input, error) {
	return aoc.ParseLines(r, func(s string) ([]string, error) {
		return strings.Fields(s), nil
	})
}

func part1(input Input) string {
	c := NewComputer(input)
	c.A = 7

	c.Run()

	return aoc.Result(c.A)
}

func part2(input Input) string {
	c := NewComputer(input)
	c.A = 12

	c.Run()

	return aoc.Result(c.A)
}

func isRegister(s string) bool {
	return len(s) == 1 && strings.Contains("abcd", s)
}

type Computer struct {
	A, B, C, D int
	Program    []interface{}
}

func NewComputer(input Input) *Computer {
	c := new(Computer)

	// compile instructions
	for _, instr := range input {
		var instruction interface{}
		switch instr[0] {
		case "tgl":
			instruction = &Toggle{Offset: c.getReg(instr[1])}
		case "cpy":
			instruction = &Copy{From: c.getReg(instr[1]), To: c.getReg(instr[2])}
		case "inc":
			instruction = &Increment{Reg: c.getReg(instr[1])}
		case "dec":
			instruction = &Decrement{Reg: c.getReg(instr[1])}
		case "jnz":
			instruction = &Jump{Condition: c.getReg(instr[1]), Offset: c.getReg(instr[2])}
		default:
			panic(fmt.Sprintf("unknown instruction %s", instr[0]))
		}
		c.Program = append(c.Program, instruction)
	}

	return c
}

func (c *Computer) getReg(s string) *int {
	if !isRegister(s) {
		n, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		return &n
	}
	switch s {
	case "a":
		return &c.A
	case "b":
		return &c.B
	case "c":
		return &c.C
	case "d":
		return &c.D
	default:
		panic(fmt.Sprintf("unknown register %s", s))
	}
}

func (c *Computer) toggle(at int) {
	if at < 0 || at >= len(c.Program) {
		return
	}

	switch instr := c.Program[at].(type) {
	case *Toggle:
		c.Program[at] = &Increment{
			Reg: instr.Offset,
		}
	case *Copy:
		c.Program[at] = &Jump{
			Condition: instr.From,
			Offset:    instr.To,
		}
	case *Increment:
		c.Program[at] = &Decrement{
			Reg: instr.Reg,
		}
	case *Decrement:
		c.Program[at] = &Increment{
			Reg: instr.Reg,
		}
	case *Jump:
		c.Program[at] = &Copy{
			From: instr.Condition,
			To:   instr.Offset,
		}
	default:
		panic(fmt.Sprintf("unknown instruction %T", instr))
	}
}

func (c *Computer) Run() {
	var stackPtr int
	var cycles int

	for stackPtr >= 0 && stackPtr < len(c.Program) {
		cycles++
		v := c.Program[stackPtr]

		switch instr := v.(type) {
		case *Toggle:
			target := stackPtr + *instr.Offset
			//log.Printf("toggle  %d: %T", target, c.Program[target])
			c.toggle(target)
		case *Copy:
			*instr.To = *instr.From
		case *Increment:
			*instr.Reg++
		case *Decrement:
			*instr.Reg--
		case *Jump:
			if *instr.Condition != 0 {
				stackPtr += *instr.Offset
				continue
			}
		default:
			panic(fmt.Sprintf("unknown instruction %T", instr))
		}
		stackPtr++
	}

	log.Println("cycles:", cycles)
}

type Increment struct {
	Reg *int
}

type Decrement struct {
	Reg *int
}

type Copy struct {
	From, To *int
}

type Jump struct {
	Condition *int
	Offset    *int
}

type Toggle struct {
	Offset *int
}
