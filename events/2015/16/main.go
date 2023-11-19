package main

import (
	"bytes"
	"fmt"
	"github.com/pimvanhespen/aoc/2015/pkg/aoc"
	"io"
	"strconv"
)

type TickerTape map[string]int

func (t TickerTape) ContainsAll(other TickerTape) bool {
	for key, value := range other {
		if t[key] != value {
			return false
		}
	}
	return true
}

type Aunt struct {
	Number int
	Values TickerTape
}

const ticker = `children: 3
cats: 7
samoyeds: 2
pomeranians: 3
akitas: 0
vizslas: 0
goldfish: 5
trees: 3
cars: 2
perfumes: 1`

func main() {
	reader, err := aoc.Get(16)
	if err != nil {
		panic(err)
	}

	got, err := parseTickerTape([]byte(ticker), []byte("\n"))
	if err != nil {
		panic(err)
	}

	aunts, err := parseAunts(reader)
	if err != nil {
		panic(err)
	}

	res1 := part1(got, aunts)
	fmt.Println("1:", res1)

	res2 := part2(got, aunts)
	fmt.Println("2:", res2)
}

func part1(got TickerTape, aunts []Aunt) int {
	keep := make([]int, 0, len(aunts))
	for _, aunt := range aunts {
		if !got.ContainsAll(aunt.Values) {
			continue
		}
		keep = append(keep, aunt.Number)
	}
	if len(keep) != 1 {
		panic("expected 1 aunt")
	}
	return keep[0]
}

func fewer(a, b int) bool   { return a < b }
func greater(a, b int) bool { return a > b }
func equal(a, b int) bool   { return a == b }

func correctAmount(key string, got, want int) bool {
	switch key {
	case "cats", "trees":
		return greater(got, want)
	case "pomeranians", "goldfish":
		return fewer(got, want)
	default:
		return equal(got, want)
	}
}

func part2(got TickerTape, aunts []Aunt) int {

	keep := make([]int, 0, len(aunts))
outer:
	for _, aunt := range aunts {
		for key, value := range aunt.Values {

			want := got[key]

			if !correctAmount(key, value, want) {
				continue outer
			}
		}

		keep = append(keep, aunt.Number)
	}
	if len(keep) != 1 {
		panic(fmt.Sprintf("expected 1 aunt, got %d", len(keep)))
	}
	return keep[0]
}

func parseAunts(reader io.Reader) ([]Aunt, error) {
	var aunts []Aunt

	b, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	b = bytes.TrimSpace(b)

	for _, line := range bytes.Split(b, []byte("\n")) {

		end := bytes.IndexByte(line, ':')
		num := line[4:end] // len("Sue ") == 4

		n, err := strconv.Atoi(string(num))
		if err != nil {
			return nil, fmt.Errorf("could not parse aunt number: %w", err)
		}

		tape, err := parseTickerTape(line[end+2:], []byte(", "))
		if err != nil {
			return nil, fmt.Errorf("could not parse ticker tape: %w", err)
		}

		aunts = append(aunts, Aunt{
			Number: n,
			Values: tape,
		})
	}

	return aunts, nil
}

func parseTickerTape(input []byte, sep []byte) (TickerTape, error) {
	tape := make(TickerTape)

	for _, field := range bytes.Split(input, sep) {
		pair := bytes.Split(field, []byte(": "))
		if len(pair) != 2 {
			return nil, fmt.Errorf("invalid field: %s", field)
		}

		value, err := strconv.Atoi(string(pair[1]))
		if err != nil {
			return nil, fmt.Errorf("invalid value: %s", pair[1])
		}

		tape[string(pair[0])] = value
	}

	return tape, nil
}
