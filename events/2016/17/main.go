package main

import (
	"crypto/md5"
	"fmt"
	"github.com/pimvanhespen/advent-of-code/pkg/datastructures/heap"
	"io"

	"github.com/pimvanhespen/advent-of-code/pkg/aoc"
)

type Input struct {
	Passcode string
}

func main() {
	event := aoc.New(2016, 17, parse)
	fmt.Println("1:", aoc.Must(event.Run(part1)))
	fmt.Println("2:", aoc.Must(event.Run(part2)))
}

type Direction uint8

const (
	Up Direction = iota
	Down
	Left
	Right
)

func (d Direction) byte() byte {
	switch d {
	case Up:
		return 'U'
	case Down:
		return 'D'
	case Left:
		return 'L'
	case Right:
		return 'R'
	}
	panic("invalid direction")
}

func doors(passcode string, path string) [4]bool {
	sum := md5.Sum([]byte(passcode + path))
	hex := fmt.Sprintf("%x", sum)
	return [4]bool{
		hex[Up] > 'a' && hex[Up] < 'g',
		hex[Down] > 'a' && hex[Down] < 'g',
		hex[Left] > 'a' && hex[Left] < 'g',
		hex[Right] > 'a' && hex[Right] < 'g',
	}
}

func parse(r io.Reader) (Input, error) {
	b, err := aoc.ReadAll(r)
	if err != nil {
		return Input{}, err
	}

	return Input{
		Passcode: string(b),
	}, nil
}

func part1(input Input) string {

	q := heap.NewMin[int, string]() // priority queue
	q.Push("", 0)

	var leastSize int = 9999
	var leastPath string

	for !q.Empty() {
		elem := q.Pop()

		if len(elem) > leastSize {
			continue
		}

		if finished([]byte(elem)) {
			if len(elem) < leastSize {
				leastSize = len(elem)
				leastPath = elem
			}
			continue
		}

		for i, open := range doors(input.Passcode, elem) {
			if !open {
				continue
			}

			path := append([]byte(elem), Direction(i).byte())

			if accessible(location(path)) {
				q.Push(elem+string(Direction(i).byte()), len(elem)+1)
			}
		}
	}

	return leastPath
}

type Vec2 struct {
	X, Y int
}

func accessible(v Vec2) bool {
	return v.X >= 0 && v.X < 4 && v.Y >= 0 && v.Y < 4
}

func location(path []byte) Vec2 {
	var v Vec2
	for _, c := range path {
		switch c {
		case 'U':
			v.Y--
		case 'D':
			v.Y++
		case 'L':
			v.X--
		case 'R':
			v.X++
		}
	}
	return v
}

func finished(path []byte) bool {
	return location(path) == Vec2{3, 3}
}

func part2(input Input) string {
	q := heap.NewMin[int, string]() // priority queue
	q.Push("", 0)

	var longest string

	for !q.Empty() {
		elem := q.Pop()

		if finished([]byte(elem)) {
			if len(elem) > len(longest) {
				longest = elem
			}
			continue
		}

		for i, open := range doors(input.Passcode, elem) {
			if !open {
				continue
			}

			path := append([]byte(elem), Direction(i).byte())

			if accessible(location(path)) {
				q.Push(elem+string(Direction(i).byte()), len(elem)+1)
			}
		}
	}

	return aoc.Result(len(longest))
}
