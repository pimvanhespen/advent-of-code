package main

import (
	"fmt"
	"io"
	"strings"

	"github.com/pimvanhespen/advent-of-code/pkg/aoc"
)

type Input struct {
	Seed []byte
	Size int
}

func main() {
	event := aoc.New(2016, 16, parse)
	fmt.Println("1:", aoc.Must(event.Run(part1)))
	fmt.Println("2:", aoc.Must(event.Run(part2)))
}

func parse(r io.Reader) (Input, error) {

	b, err := aoc.ReadAll(r)
	if err != nil {
		return Input{}, err
	}

	return Input{
		Seed: b,
		Size: 272,
	}, nil
}

func AND(a, b byte) byte {
	if a == b {
		return '1'
	}
	return '0'
}

func checksumBytes(data []byte) []byte {

	half := make([]byte, len(data)/2)

	for i := 0; i < len(data); i += 2 {
		half[i/2] = AND(data[i], data[i+1])
	}

	if len(half)%2 == 0 {
		return checksumBytes(half)
	} else {
		return half
	}
}

func invert(data []byte) []byte {
	result := make([]byte, len(data))

	for i := 0; i < len(data); i++ {
		switch data[i] {
		case '0':
			result[len(result)-1-i] = '1'
		case '1':
			result[len(result)-1-i] = '0'
		}
	}

	return result
}

func dragon(data []byte) []byte {

	result := make([]byte, 0, len(data)*2+1)
	result = append(result, data...)
	result = append(result, '0')
	result = append(result, invert(data)...)

	return result
}

func part1(input Input) string {

	data := input.Seed

	for len(data) < input.Size {
		data = dragon(data)
	}

	data = data[:input.Size]

	var sb strings.Builder
	for _, b := range checksumBytes(data) {
		sb.WriteString(fmt.Sprintf("%b", b))
	}

	return string(checksumBytes(data))
}

func part2(input Input) string {
	data := input.Seed

	size := 35_651_584 // ~35MB

	for len(data) < size {
		data = dragon(data)
	}

	data = data[:size]

	return string(checksumBytes(data))
}
