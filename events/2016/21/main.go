package main

import (
	"bytes"
	"fmt"
	"github.com/pimvanhespen/advent-of-code/pkg/aoc"
	"io"
	"slices"
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
	res := make([]byte, len(input.Seed))
	copy(res, input.Seed)

	for _, instruction := range input.Instructions {
		instruction.Apply(res)
	}
	return string(res)
}

func part2(input Input) string {
	res := []byte("fbgdceah")

	// Duo to the nature of the instructions, we can't just apply them in reverse.
	// the rotate based on letter instruction is not (easily) reversible.
	// Instead, we can try all permutations of the instructions and see which one
	// results in the correct output.
	permutations := permute(res, input.Instructions)

	for _, permutation := range permutations {
		before := string(permutation)
		for _, instruction := range input.Instructions {
			instruction.Apply(permutation)
		}
		if bytes.Equal(permutation, res) {
			return before
		}
	}
	return "n/a"
}

func permute(input []byte, todo []Instruction) [][]byte {
	if len(todo) == 0 {
		return [][]byte{input}
	}

	var results [][]byte

	ins := todo[len(todo)-1]
	for _, undo := range ins.Undo(input) {
		results = append(results, permute(undo, todo[:len(todo)-1])...)
	}
	return results
}

type Instruction interface {
	Apply([]byte)
	Undo([]byte) [][]byte
}

type SwapPosition struct {
	From int
	To   int
}

func (s SwapPosition) Apply(data []byte) {
	data[s.From], data[s.To] = data[s.To], data[s.From]
}

func (s SwapPosition) Undo(data []byte) [][]byte {
	c := copyBytes(data)
	s.Apply(c)
	return [][]byte{c}
}

type SwapLetter struct {
	From byte
	To   byte
}

func (s SwapLetter) Apply(data []byte) {
	from := bytes.IndexByte(data, s.From)
	to := bytes.IndexByte(data, s.To)
	data[from], data[to] = data[to], data[from]
}

func (s SwapLetter) Undo(data []byte) [][]byte {
	c := copyBytes(data)
	s.Apply(c)
	return [][]byte{c}
}

type RotateLeft struct {
	Steps int
}

func (r RotateLeft) Apply(data []byte) {
	rotateLeft(data, r.Steps)
}

func (r RotateLeft) Undo(data []byte) [][]byte {
	c := copyBytes(data)
	rotateLeft(c, -r.Steps)
	return [][]byte{c}
}

type RotateRight struct {
	Steps int
}

func (r RotateRight) Apply(data []byte) {
	rotateLeft(data, -r.Steps)
}

func (r RotateRight) Undo(data []byte) [][]byte {
	c := copyBytes(data)
	rotateLeft(c, r.Steps)
	return [][]byte{c}
}

type RotateBasedOnPosition struct {
	Letter byte
}

func (r RotateBasedOnPosition) Apply(data []byte) {
	steps := bytes.IndexByte(data, r.Letter)
	if steps >= 4 {
		steps++
	}
	steps++

	rotateLeft(data, -steps)
}

func (r RotateBasedOnPosition) Undo(data []byte) [][]byte {
	var opts [][]byte
	for i := 0; i < len(data); i++ {
		prev := make([]byte, len(data))
		copy(prev, data)

		rot := i
		if i >= 4 {
			rot++
		}
		rot++

		rotateLeft(prev, rot)

		before := copyBytes(prev)

		r.Apply(prev)

		if bytes.Equal(prev, data) {
			opts = append(opts, before)
		}
	}
	return opts
}

type ReversePositions struct {
	From int
	To   int
}

func (r ReversePositions) Apply(data []byte) {
	slices.Reverse(data[r.From : r.To+1])
}

func (r ReversePositions) Undo(data []byte) [][]byte {
	c := copyBytes(data)
	r.Apply(c)
	return [][]byte{c}
}

type MovePosition struct {
	From int
	To   int
}

func (m MovePosition) Apply(input []byte) {
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

	copy(input, res)
}

func (m MovePosition) Undo(input []byte) [][]byte {
	c := copyBytes(input)
	m.From, m.To = m.To, m.From
	m.Apply(c)
	return [][]byte{c}
}

func rotateLeft(input []byte, steps int) {
	steps = (steps + len(input)) % len(input)

	if steps == 0 {
		return
	}

	before := make([]byte, len(input))
	copy(before, input)

	for i := range input {
		input[i] = before[(i+len(input)+steps)%len(input)]
	}
}

func copyBytes(data []byte) []byte {
	res := make([]byte, len(data))
	copy(res, data)
	return res
}
