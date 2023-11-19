package main

import (
	"github.com/pimvanhespen/advent-of-code/pkg/aoc"
	"io"
	"log"
)

func main() {

	r, err := aoc.NewChallenge(2015, 3).Input()
	if err != nil {
		log.Fatal(err)
	}

	directions, err := parse(r)
	if err != nil {
		log.Fatal(err)
	}

	one := solve1(directions)
	log.Println("one:", one)

	two := solve2(directions)
	log.Println("two:", two)
}

func parse(reader io.Reader) ([]Direction, error) {
	b, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	s := make([]Direction, 0, len(b))

	for _, c := range b {
		s = append(s, Direction(c))
	}

	return s, nil
}

type Direction uint8

const (
	Up    Direction = '^'
	Down  Direction = 'v'
	Left  Direction = '<'
	Right Direction = '>'
)

type Point struct {
	X, Y int
}

func (p Point) Move(d Direction) Point {
	switch d {
	case Up:
		p.Y++
	case Down:
		p.Y--
	case Left:
		p.X--
	case Right:
		p.X++
	default:
		panic("invalid direction: " + string(d))
	}

	return p
}

func solve1(directions []Direction) int {

	m := make(map[Point]uint16)

	var p Point

	m[p] = 1

	for _, d := range directions {
		p = p.Move(d)
		m[p]++
	}

	var multiple int

	for _, v := range m {
		if v > 0 {
			multiple++
		}
	}

	return multiple
}

func solve2(directions []Direction) int {

	var santa, robo Point

	m := map[Point]uint{
		Point{}: 2,
	}

	move := func(pos Point, d Direction) Point {
		pos = pos.Move(d)

		m[pos]++

		return pos
	}

	for i, d := range directions {
		if i%2 == 0 {
			santa = move(santa, d)
		} else {
			robo = move(robo, d)
		}
	}

	var multiple int

	for _, v := range m {
		if v > 0 {
			multiple++
		}
	}

	return multiple
}
