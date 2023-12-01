package main

import (
	"fmt"
	"log"
	"strconv"
)

type ComputerV1 struct {
	program Input
	regs    map[string]int
	ptr     int
}

func NewComputerV1(input Input) *ComputerV1 {
	return &ComputerV1{
		program: input,
		regs: map[string]int{
			"a": 0,
			"b": 0,
			"c": 0,
			"d": 0,
		},
		ptr: 0,
	}
}

func (c *ComputerV1) Move(offset int) {
	c.ptr += offset
}

func (c *ComputerV1) Tgl(arg string) {
	defer c.Move(1)

	ptr := c.ptr + c.regs[arg]
	if ptr < 0 || ptr >= len(c.program) {
		return
	}

	var newInstr string
	switch c.program[ptr][0] {
	case "inc":
		newInstr = "dec"
	case "dec", "tgl":
		newInstr = "inc"
	case "jnz":
		newInstr = "cpy"
	case "cpy":
		newInstr = "jnz"
	default:
		panic(fmt.Sprintf("unknown instruction %s", c.program[ptr][0]))
	}

	log.Println("toggling", ptr, c.program[ptr], "to", newInstr)
	c.program[ptr][0] = newInstr
}

func (c *ComputerV1) Copy(arg1, arg2 string) {
	defer c.Move(1)
	if isRegister(arg2) {
		c.regs[arg2] = c.Value(arg1)
	}
}

func (c *ComputerV1) Inc(arg1 string) {
	defer c.Move(1)
	if isRegister(arg1) {
		c.regs[arg1]++
	}
}

func (c *ComputerV1) Dec(arg1 string) {
	defer c.Move(1)
	if isRegister(arg1) {
		c.regs[arg1]--
	}
}

func (c *ComputerV1) Jnz(arg1, arg2 string) {
	if c.Value(arg1) != 0 {
		c.Move(c.Value(arg2))
	} else {
		c.Move(1)
	}
}

func (c *ComputerV1) Value(arg1 string) int {
	if isRegister(arg1) {
		return c.regs[arg1]
	}

	v, err := strconv.Atoi(arg1)
	if err != nil {
		panic(err)
	}
	return v
}

func (c *ComputerV1) Run(regs map[string]int) {
	if regs != nil {
		c.regs = regs
	}

	for c.ptr >= 0 && c.ptr < len(c.program) {
		instr := c.program[c.ptr]

		//time.Sleep(1 * time.Millisecond)
		//log.Println(c.ptr, instr, c.regs)

		switch instr[0] {
		case "tgl":
			c.Tgl(instr[1])
		case "cpy":
			c.Copy(instr[1], instr[2])
		case "inc":
			c.Inc(instr[1])
		case "dec":
			c.Dec(instr[1])
		case "jnz":
			c.Jnz(instr[1], instr[2])
		default:
			panic(fmt.Sprintf("unknown instruction %s", instr[0]))
		}
	}
}
