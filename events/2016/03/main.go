package main

import (
	"fmt"
	"io"

	"github.com/pimvanhespen/advent-of-code/pkg/aoc"
)

type Input []Triangle

type Triangle struct {
	A, B, C int
}

func (t Triangle) Valid() bool {
	return t.A+t.B > t.C &&
		t.A+t.C > t.B &&
		t.B+t.C > t.A
}

func main() {
	event := aoc.New(2016, 3, parse)
	fmt.Println("1:", aoc.Must(event.Run(part1)))
	fmt.Println("2:", aoc.Must(event.Run(part2)))
}

func parse(r io.Reader) (Input, error) {
	return aoc.ParseLines(r, func(line string) (Triangle, error) {
		var t Triangle
		_, err := fmt.Sscanf(line, "%d %d %d", &t.A, &t.B, &t.C)
		return t, err
	})
}

func part1(input Input) string {

	var count int
	for _, t := range input {
		if t.Valid() {
			count++
		}
	}

	return aoc.Result(count)
}

func part2(input Input) string {

	converted := make([]Triangle, 0, len(input))
	for i := 0; i < len(input); i += 3 {
		converted = append(converted, Triangle{
			input[i].A, input[i+1].A, input[i+2].A,
		})
		converted = append(converted, Triangle{
			input[i].B, input[i+1].B, input[i+2].B,
		})
		converted = append(converted, Triangle{
			input[i].C, input[i+1].C, input[i+2].C,
		})
	}

	var count int
	for _, t := range converted {
		if t.Valid() {
			count++
		}
	}

	return aoc.Result(count)
}
