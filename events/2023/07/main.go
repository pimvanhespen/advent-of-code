package main

import (
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"

	"github.com/pimvanhespen/advent-of-code/pkg/aoc"
)

type Input = []Player

type HandType uint8

func (h HandType) String() string {
	switch h {
	case HighCard:
		return "High Card"
	case OnePair:
		return "One Pair"
	case TwoPairs:
		return "Two Pairs"
	case ThreeOfAKind:
		return "Three of a Kind"
	case FullHouse:
		return "Full House"
	case FourOfAKind:
		return "Four of a Kind"
	case FiveOfAKind:
		return "Five of a Kind"
	}
	return "unknown"
}

const (
	HighCard HandType = iota
	OnePair
	TwoPairs
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

type Player struct {
	Cards string
	Bid   int
}

func (p Player) String() string {
	return fmt.Sprintf("%5s %3d", p.Cards, p.Bid)
}

func main() {
	event := aoc.New(2023, 7, parse)
	fmt.Println("1:", aoc.Must(event.Run(part1)))
	fmt.Println("2:", aoc.Must(event.Run(part2)))
}

func parse(r io.Reader) (Input, error) {
	return aoc.ParseLines(r, func(s string) (Player, error) {

		parts := strings.Split(s, " ")

		bid, err := strconv.Atoi(parts[1])
		if err != nil {
			return Player{}, fmt.Errorf("parse bid: %w", err)
		}

		return Player{
			Cards: parts[0],
			Bid:   bid,
		}, nil
	})
}

func part1(input Input) string {

	sortInput(input, CardValue1, Type1)

	var sum int

	for i, p := range input {
		sum += p.Bid * (len(input) - i)
	}

	return aoc.Result(sum)
}

func part2(input Input) string {

	sortInput(input, CardValue2, Type2)

	var sum int

	for i, p := range input {
		sum += p.Bid * (len(input) - i)
	}

	return aoc.Result(sum)
}

func sortInput(input Input, val func(byte) int, typer func(string) HandType) {
	sort.Slice(input, func(i, j int) bool {
		ta, tb := typer(input[i].Cards), typer(input[j].Cards)

		if ta != tb {
			return ta > tb
		}

		for x := 0; x < 5; x++ {
			if val(input[i].Cards[x]) != val(input[j].Cards[x]) {
				return val(input[i].Cards[x]) > val(input[j].Cards[x])
			}
		}
		return false
	})
}

func Type1(s string) HandType {
	cards := []byte(s)
	if len(cards) != 5 {
		panic("invalid hand")
	}

	groups := groupCards(cards)

	switch len(groups) {
	case 1:
		return FiveOfAKind
	case 2:
		if len(groups[0]) == 1 || len(groups[0]) == 4 {
			return FourOfAKind
		}
		return FullHouse
	case 3:
		if len(groups[0]) == 3 || len(groups[1]) == 3 || len(groups[2]) == 3 {
			return ThreeOfAKind
		}
		return TwoPairs
	case 4:
		return OnePair
	case 5:
		return HighCard
	}

	panic("invalid hand")
}

func Type2(s string) HandType {
	if len(s) != 5 {
		panic("invalid hand")
	}

	jokers := strings.Count(s, "J")
	if jokers == 0 {
		return Type1(s)
	}

	s = strings.ReplaceAll(s, "J", "")

	groups := groupCards([]byte(s))
	switch jokers {
	case 1:
		switch len(groups) {
		case 1:
			return FiveOfAKind
		case 2:
			if len(groups[0]) == len(groups[1]) {
				return FullHouse
			}
			return FourOfAKind
		case 3:
			return ThreeOfAKind
		case 4:
			return OnePair
		}
	case 2:
		switch len(groups) {
		case 1:
			return FiveOfAKind
		case 2:
			return FourOfAKind
		case 3:
			return ThreeOfAKind
		}
	case 3:
		switch len(groups) {
		case 1:
			return FiveOfAKind
		case 2:
			return FourOfAKind
		}
	case 4, 5:
		return FiveOfAKind
	}
	panic("invalid hand")
}

func CardValue1(c byte) int {
	switch c {
	case 'A':
		return 14
	case 'K':
		return 13
	case 'Q':
		return 12
	case 'J':
		return 11
	case 'T':
		return 10
	case '9', '8', '7', '6', '5', '4', '3', '2':
		return int(c - '0')
	default:
		panic("invalid card")
	}
}

func CardValue2(c byte) int {
	switch c {
	case 'A':
		return 14
	case 'K':
		return 13
	case 'Q':
		return 12
	case 'J':
		return 1
	case 'T':
		return 10
	case '9', '8', '7', '6', '5', '4', '3', '2':
		return int(c - '0')
	default:
		panic("invalid card")
	}
}

func groupCards(b []byte) [][]byte {
	sort.Slice(b, func(i, j int) bool {
		return b[i] < b[j]
	})

	var result [][]byte
	var current []byte
	for _, c := range b {
		if len(current) == 0 || current[0] == c {
			current = append(current, c)
			continue
		}
		result = append(result, current)
		current = []byte{c}
	}
	if len(current) > 0 {
		result = append(result, current)
	}
	return result
}
