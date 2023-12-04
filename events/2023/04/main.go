package main

import (
	"fmt"
	"io"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/pimvanhespen/advent-of-code/pkg/aoc"
)

type Input []Card

type Card struct {
	ID      int
	Winning []int
	Have    []int
}

func (c Card) String() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("%3d: ", c.ID))
	for _, v := range c.Winning {
		sb.WriteString(fmt.Sprintf("%2d ", v))
	}
	sb.WriteString(" | ")
	for _, v := range c.Have {
		sb.WriteString(fmt.Sprintf("%2d ", v))
	}
	return sb.String()
}

func (c Card) Matches() int {
	return contains(c.Winning, c.Have)
}

func (c Card) Score() int {
	n := c.Matches()
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	default:
		return int(math.Pow(2, float64(n-1)))
	}
}

func contains(winning []int, have []int) int {
	var count int
	for _, h := range have {
		if slices.Contains(winning, h) {
			count++
		}
	}
	return count
}

func main() {
	event := aoc.New(2023, 4, parse)
	fmt.Println("1:", aoc.Must(event.Run(part1)))
	fmt.Println("2:", aoc.Must(event.Run(part2)))
}

func parse(r io.Reader) (Input, error) {
	return aoc.ParseLines(r, func(line string) (Card, error) {
		parts := strings.Split(line, ": ")

		numString := strings.TrimSpace(parts[0][5:])

		cardID, err := strconv.Atoi(numString)
		if err != nil {
			return Card{}, fmt.Errorf("could not parse card id: %w", err)
		}

		lists := strings.Split(parts[1], " | ")
		if len(lists) != 2 {
			return Card{}, fmt.Errorf("expected 2 lists, got %d", len(lists))
		}
		winning, err := parseInts(lists[0], " ")
		if err != nil {
			return Card{}, fmt.Errorf("parse winning: %w", err)
		}

		have, err := parseInts(lists[1], " ")
		if err != nil {
			return Card{}, fmt.Errorf("parse have: %w", err)
		}

		return Card{
			ID:      cardID,
			Winning: winning,
			Have:    have,
		}, nil
	})
}

func parseInts(s string, sep string) ([]int, error) {
	parts := strings.Split(s, sep)
	ints := make([]int, 0, len(parts))
	for _, v := range parts {
		v = strings.TrimSpace(v)
		if v == "" {
			continue
		}
		n, err := strconv.Atoi(v)
		if err != nil {
			return nil, fmt.Errorf("couldn't parse %q: %w", v, err)
		}
		ints = append(ints, n)
	}
	return ints, nil
}

func part1(input Input) string {
	var sum int
	for _, card := range input {
		sum += card.Score()
	}
	return aoc.Result(sum)
}

func part2(input Input) string {
	cards := make([]int, len(input)+1)

	for _, card := range input {
		amount := card.Matches()
		for i := 1 + card.ID; i <= card.ID+amount; i++ {
			additional := 1 + cards[card.ID]
			cards[i] += additional
		}
	}

	sum := len(input)
	for _, v := range cards {
		sum += v
	}
	return aoc.Result(sum)
}
