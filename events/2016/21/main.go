package main

import (
	"bytes"
	"fmt"
	"github.com/pimvanhespen/advent-of-code/pkg/aoc"
	"io"
	"strconv"
	"strings"
)

type Input struct {
	Seed         []byte
	Instructions []Instruction
}

func main() {
	event := aoc.New(2016, 21, parse)
	fmt.Println("1:", aoc.Must(event.Run(part1)))
	fmt.Println("2:", aoc.Must(event.Run(part2)))
}

type Instruction interface {
	Execute([]byte) []byte
}

type SwapPosition struct {
	From int
	To   int
}

func (s SwapPosition) Execute(input []byte) []byte {
	res := make([]byte, len(input))
	copy(res, input)
	res[s.From], res[s.To] = res[s.To], res[s.From]
	return res
}

type SwapLetter struct {
	From byte
	To   byte
}

func (s SwapLetter) Execute(input []byte) []byte {
	res := make([]byte, len(input))
	copy(res, input)
	from := bytes.IndexByte(res, s.From)
	to := bytes.IndexByte(res, s.To)
	res[from], res[to] = res[to], res[from]
	return res
}

type RotateLeft struct {
	Steps int
}

func (r RotateLeft) Execute(input []byte) []byte {
	res := make([]byte, len(input))
	copy(res, input)

	return rotateLeft(res, r.Steps)
}

func rotateLeft(input []byte, steps int) []byte {
	steps = (steps + len(input)) % len(input)

	if steps == 0 {
		return input
	}

	return append(input[steps:], input[:steps]...)
}

type RotateRight struct {
	Steps int
}

func (r RotateRight) Execute(input []byte) []byte {
	res := make([]byte, len(input))
	copy(res, input)
	return rotateLeft(res, -r.Steps)
}

type RotateBasedOnPosition struct {
	Letter byte
}

func (r RotateBasedOnPosition) Execute(input []byte) []byte {
	steps := bytes.IndexByte(input, r.Letter)
	if steps >= 4 {
		steps++
	}
	steps++

	res := make([]byte, len(input))
	copy(res, input)
	return rotateLeft(res, -steps)
}

type ReversePositions struct {
	From int
	To   int
}

func (r ReversePositions) Execute(input []byte) []byte {

	res := make([]byte, len(input))
	copy(res, input)

	for i := 0; i <= (r.To - r.From); i++ {
		res[r.From+i] = input[r.To-i]
	}

	return res
}

type MovePosition struct {
	From int
	To   int
}

func (m MovePosition) Execute(input []byte) []byte {
	res := make([]byte, len(input))
	copy(res, input)

	// move by overwriting the destination
	if m.From < m.To {
		// move to the right
		copy(res[m.From:], input[m.From+1:m.To+1])
	} else {
		// move to the left
		copy(res[m.To+1:], input[m.To:m.From])
	}
	res[m.To] = input[m.From]

	return res
}

func parse(reader io.Reader) (Input, error) {
	instructions, err := aoc.ParseLines(reader, func(line string) (Instruction, error) {
		fields := strings.Fields(line)

		switch fields[0] {
		case "swap":
			switch fields[1] {
			case "position":
				from, err := strconv.Atoi(fields[2])
				if err != nil {
					return nil, err
				}
				to, err := strconv.Atoi(fields[5])
				if err != nil {
					return nil, err
				}
				return &SwapPosition{
					From: from,
					To:   to,
				}, nil
			case "letter":
				return &SwapLetter{
					From: fields[2][0],
					To:   fields[5][0],
				}, nil
			}
		case "rotate":
			switch fields[1] {
			case "left":
				steps, err := strconv.Atoi(fields[2])
				if err != nil {
					return nil, err
				}
				return &RotateLeft{
					Steps: steps,
				}, nil
			case "right":
				steps, err := strconv.Atoi(fields[2])
				if err != nil {
					return nil, err
				}
				return &RotateRight{
					Steps: steps,
				}, nil
			case "based":
				return &RotateBasedOnPosition{
					Letter: fields[6][0],
				}, nil
			}
		case "reverse":
			from, err := strconv.Atoi(fields[2])
			if err != nil {
				return nil, err
			}
			to, err := strconv.Atoi(fields[4])
			if err != nil {
				return nil, err
			}
			return &ReversePositions{
				From: from,
				To:   to,
			}, nil
		case "move":
			from, err := strconv.Atoi(fields[2])
			if err != nil {
				return nil, err
			}
			to, err := strconv.Atoi(fields[5])
			if err != nil {
				return nil, err
			}
			return &MovePosition{
				From: from,
				To:   to,
			}, nil
		}

		return nil, nil
	})
	if err != nil {
		return Input{}, err
	}

	return Input{
		Seed:         []byte("abcdefgh"),
		Instructions: instructions,
	}, nil
}

func part1(input Input) string {
	res := input.Seed
	for _, instruction := range input.Instructions {
		res = instruction.Execute(res)
	}
	return string(res)
}

func part2(input Input) string {
	return "n/a"
}
