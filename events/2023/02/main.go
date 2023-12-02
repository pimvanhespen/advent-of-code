package main

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/pimvanhespen/advent-of-code/pkg/aoc"
)

type Game struct {
	ID       int
	Revealed []Cubes
}

func (g Game) Required() Cubes {
	var m Cubes
	for _, kubes := range g.Revealed {
		m.Red = max(m.Red, kubes.Red)
		m.Green = max(m.Green, kubes.Green)
		m.Blue = max(m.Blue, kubes.Blue)
	}
	return m
}

type Cubes struct {
	Red, Green, Blue int
}

func CubesFromString(s string) Cubes {
	var k Cubes
	parts := strings.Split(s, ", ")
	for _, part := range parts {
		switch {
		case strings.HasSuffix(part, "red"):
			n, err := strconv.Atoi(strings.TrimSuffix(part, " red"))
			if err != nil {
				panic(err)
			}
			k.Red = n
		case strings.HasSuffix(part, "green"):
			n, err := strconv.Atoi(strings.TrimSuffix(part, " green"))
			if err != nil {
				panic(err)
			}
			k.Green = n
		case strings.HasSuffix(part, "blue"):
			n, err := strconv.Atoi(strings.TrimSuffix(part, " blue"))
			if err != nil {
				panic(err)
			}
			k.Blue = n
		default:
			panic("unknown kube")
		}
	}
	return k
}

type Input []Game

func main() {
	event := aoc.New(2023, 2, parse)
	fmt.Println("1:", aoc.Must(event.Run(part1)))
	fmt.Println("2:", aoc.Must(event.Run(part2)))
}

func parse(r io.Reader) (Input, error) {
	return aoc.ParseLines(r, func(line string) (Game, error) {

		parts := strings.Split(line, ": ")
		id := aoc.Must(strconv.Atoi(strings.Split(parts[0], " ")[1])) // Game (1)
		sets := strings.Split(parts[1], "; ")

		revealed := make([]Cubes, len(sets))
		for i, set := range sets {
			revealed[i] = CubesFromString(set)
		}

		game := Game{
			ID:       id,
			Revealed: revealed,
		}
		return game, nil
	})
}

func part1(input Input) string {

	keep := limit(Cubes{12, 13, 14})

	var sum int
	for _, game := range input {
		if keep(game.Required()) {
			sum += game.ID
		}
	}
	return strconv.Itoa(sum)
}

func part2(input Input) string {
	var sum int
	for _, game := range input {
		req := game.Required()
		sum += req.Red * req.Green * req.Blue
	}
	return strconv.Itoa(sum)
}

func limit(lim Cubes) func(Cubes) bool {
	return func(k Cubes) bool {
		return k.Red <= lim.Red && k.Green <= lim.Green && k.Blue <= lim.Blue
	}
}
