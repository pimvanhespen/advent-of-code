package main

import (
	"fmt"
	"github.com/pimvanhespen/advent-of-code/pkg/aoc"
	"io"
	"strconv"
	"strings"
)

type Instruction struct {
	Op   string
	Args []any
}

func (i Instruction) String() string {
	args := make([]string, len(i.Args))
	for x, arg := range i.Args {
		args[x] = fmt.Sprintf("%v", arg)
	}
	return fmt.Sprintf("%s %s", i.Op, strings.Join(args, ", "))
}

type Input struct {
	Program []Instruction
}

func main() {

	reader, err := aoc.NewChallenge(2015, 23).Input()
	if err != nil {
		panic(err)
	}

	input, err := parse(reader)
	if err != nil {
		panic(err)
	}

	fmt.Println("Part 1:", part1(input))
	fmt.Println("Part 2:", part2(input))
}

func parse(reader io.Reader) (Input, error) {

	program, err := aoc.ParseLines(reader, parseLine)
	if err != nil {
		return Input{}, err
	}

	input := Input{
		Program: program,
	}

	return input, nil
}

func parseLine(line string) (Instruction, error) {
	oper := line[:3]
	args := strings.Split(line[4:], ", ")

	switch oper {
	case "hlf":
		return Instruction{Op: oper, Args: []any{args[0]}}, nil
	case "tpl":
		return Instruction{Op: oper, Args: []any{args[0]}}, nil
	case "inc":
		return Instruction{Op: oper, Args: []any{args[0]}}, nil
	case "jmp":
		n, err := strconv.Atoi(args[0])
		if err != nil {
			return Instruction{}, err
		}
		return Instruction{Op: oper, Args: []any{n}}, nil
	case "jie":
		n, err := strconv.Atoi(args[1])
		if err != nil {
			return Instruction{}, err
		}
		return Instruction{Op: oper, Args: []any{args[0], n}}, nil
	case "jio":
		n, err := strconv.Atoi(args[1])
		if err != nil {
			return Instruction{}, err
		}
		return Instruction{Op: oper, Args: []any{args[0], n}}, nil
	}

	return Instruction{}, fmt.Errorf("unknown operation: %s", oper)
}

type Register struct {
	value int
}

type Computer struct {
	counter   int
	registers map[string]*Register
	stack     []Instruction
}

func (c *Computer) reg(id any) *Register {
	s, ok := id.(string)
	if !ok {
		panic(fmt.Errorf("unknown register: %v (%T)", id, id))
	}
	reg, ok := c.registers[id.(string)]
	if !ok {
		panic(fmt.Errorf("unknown register: %s", s))
	}
	return reg
}

func (c *Computer) Cycle() bool {
	if c.counter < 0 || c.counter >= len(c.stack) {
		return false
	}

	counter := c.counter
	defer func() {
		if counter == c.counter {
			panic(fmt.Errorf("instruction pointer did not change: %d", c.counter))
		}
	}()

	instr := c.stack[c.counter]

	fmt.Println(instr.String())

	switch instr.Op {
	case "hlf":
		r := c.reg(instr.Args[0])
		r.value /= 2

	case "tpl":
		r := c.reg(instr.Args[0])
		r.value *= 3

	case "inc":
		r := c.reg(instr.Args[0])
		r.value++

	case "jmp":
		c.counter += instr.Args[0].(int)
		return true

	case "jie":
		r := c.reg(instr.Args[0])
		if r.value%2 == 0 {
			c.counter += instr.Args[1].(int)
			return true
		}

	case "jio":
		r := c.reg(instr.Args[0])
		if r.value == 1 {
			c.counter += instr.Args[1].(int)
			return true
		}

	default:
		panic(fmt.Errorf("unknown operation: %s", instr.Op))
	}

	c.counter++
	return true
}

func (c *Computer) String() string {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("[%3d] ", c.counter))
	for name, reg := range c.registers {
		sb.WriteString(fmt.Sprintf("%s=%d, ", name, reg.value))
	}
	return sb.String()
}

func part1(input Input) int {

	computer := Computer{
		registers: map[string]*Register{
			"a": {value: 0},
			"b": {value: 0},
		},
		stack: input.Program,
	}

	fmt.Println(computer.String())
	for computer.Cycle() {
		fmt.Println(computer.String())
	}

	return computer.reg("b").value
}

func part2(input Input) int {

	computer := Computer{
		registers: map[string]*Register{
			"a": {value: 1},
			"b": {value: 0},
		},
		stack: input.Program,
	}

	fmt.Println(computer.String())
	for computer.Cycle() {
		fmt.Println(computer.String())
	}

	return computer.reg("b").value
}
