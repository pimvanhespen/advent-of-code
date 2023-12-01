package main

import (
	"fmt"
	"github.com/pimvanhespen/advent-of-code/pkg/aoc"
	"io"
	"log"
	"regexp"
	"strings"
)

type Input struct {
	Elevator int
	State    State
}

func main() {
	event := aoc.New(2016, 11, parse)
	fmt.Println("1:", aoc.Must(event.Run(part1)))
	fmt.Println("2:", aoc.Must(event.Run(part2)))
}

var (
	chipEx = regexp.MustCompile(`(\w+)-compatible microchip`)
	rtgEx  = regexp.MustCompile(`(\w+) generator`)
)

func parse(r io.Reader) (Input, error) {

	elements := make(map[string]uint8)

	floors, err := aoc.ParseLines(r, func(line string) (Floor, error) {
		var f Floor

		for _, match := range chipEx.FindAllStringSubmatch(line, -1) {

			el, ok := elements[match[1]]
			if !ok {
				elements[match[1]] = uint8(1 << len(elements))
				el = elements[match[1]]
			}

			f.Chip |= el
		}

		for _, match := range rtgEx.FindAllStringSubmatch(line, -1) {
			el, ok := elements[match[1]]
			if !ok {
				elements[match[1]] = uint8(1 << len(elements))
				el = elements[match[1]]
			}

			f.RTG |= el
		}

		return f, nil

	})
	if err != nil {
		return Input{}, err
	}

	var s State
	for i, f := range floors {
		s.Floors[i] = f // panics if len(floors) > 4
	}

	return Input{State: s}, nil
}

func part1(input Input) string {
	// Goal is to move all components to the top floor
	// We can only move 2 components at a time

	type Route struct {
		State State
		Steps uint8
	}

	queue := make([]Route, 0, 1000)
	queue = append(queue, Route{State: input.State, Steps: 0})
	least := uint8(255)

	seen := map[State]struct{}{}

	var count uint64

	for len(queue) > 0 {
		count++
		if count%100000 == 0 {
			log.Println(count, len(queue), least)
		}

		route := queue[0]
		queue = queue[1:]

		if route.Steps > least {
			continue
		}

		state := route.State

		floors := state.NextFloors()
		for _, targetFloor := range floors {

			curr := state.Floors[state.Elevator]
			target := state.Floors[targetFloor]

			// select all possible combinations of 2 components that mayb be moved to the next floor
			// select all possible combinations of 1 component that may be moved to the next floor

			opts := curr.Options(target)
			for _, move := range opts {

				// build new state

				next := state.Next(targetFloor, move)

				fmt.Println(next.String())

				if done(next.Floors) {
					least = min(least, route.Steps+1)
					continue
				}

				if _, ok := seen[next]; ok {
					continue
				}

				seen[next] = struct{}{}
				queue = append(queue, Route{State: next, Steps: route.Steps + 1})
			}
		}
	}

	// select available floors

	return aoc.Result(least)
}

func done(floors [4]Floor) bool {
	return floors[0].IsEmpty() && floors[1].IsEmpty() && floors[2].IsEmpty()
}

func part2(input Input) string {
	return "n/a"
}

// --- Day 11: Radioisotope Thermoelectric Generators ---

type Components [2]uint8

func (c Components) String() string {
	return fmt.Sprintf("C%08b R%08b", c[0], c[1])
}

type Floor struct {
	Chip uint8 // support up to 8 chips
	RTG  uint8 // support up to 8 RTGs
}

func (f *Floor) Bytes() [2]byte {
	return [2]byte{f.Chip, f.RTG}
}

func (f *Floor) String() string {
	return fmt.Sprintf("C%08b R%08b", f.Chip, f.RTG)
}

func (f *Floor) IsEmpty() bool {
	return f.Chip == 0 && f.RTG == 0
}

func (f *Floor) IsSafe() bool {
	return isSafe(f.Chip, f.RTG)
}

func (f *Floor) IsSafeWithout(c Components) bool {
	return isSafe(f.Chip&^c[0], f.RTG&^c[1])
}

func (f *Floor) IsSafeWith(c Components) bool {
	return isSafe(f.Chip|c[0], f.RTG|c[1])
}

func (f *Floor) Add(c Components) {
	f.Chip |= c[0]
	f.RTG |= c[1]
}

func (f *Floor) Remove(c Components) {
	f.Chip &^= c[0]
	f.RTG &^= c[1]
}

func (f *Floor) Options(target Floor) []Components {
	var opts []Components

	appendSafe := func(c Components) {
		if f.IsSafeWithout(c) && target.IsSafeWith(c) {
			opts = append(opts, c)
		}
	}

	for i := uint8(0); i < 8; i++ {
		// Single layer: chip or RTG
		if f.Chip&(1<<i) != 0 {
			appendSafe([2]uint8{1 << i, 0})
		}

		if f.RTG&(1<<i) != 0 {
			appendSafe([2]uint8{0, 1 << i})
		}

		if f.Chip&(1<<i) != 0 && f.RTG&(1<<i) != 0 {
			appendSafe([2]uint8{1 << i, 1 << i})
		}

		// Double layer: chip and RTG
		for j := uint8(i + 1); j < 8; j++ {
			if f.Chip&(1<<i) != 0 && f.RTG&(1<<j) != 0 {
				appendSafe([2]uint8{1 << i, 1 << j})
			}

			// Double layer: 2 chips
			if f.Chip&(1<<i) != 0 && f.Chip&(1<<j) != 0 {
				appendSafe([2]uint8{1 << i, 1 << j})
			}

			// Double layer: 2 RTGs
			if f.RTG&(1<<i) != 0 && f.RTG&(1<<j) != 0 {
				appendSafe([2]uint8{1 << i, 1 << j})
			}
		}
	}

	return opts
}

func isSafe(chip, rtg uint8) bool {
	return chip == 0 || rtg&^chip == 0
}

type State struct {
	Elevator uint8
	Floors   [4]Floor
}

func (s State) Bytes() []byte {
	var b [9]byte
	b[0] = s.Elevator
	for i, f := range s.Floors {
		b[i+1] = f.Chip
		b[i+5] = f.RTG
	}
	return b[:]
}

func (s State) String() string {
	var sb strings.Builder

	for i := len(s.Floors) - 1; i >= 0; i-- {
		f := s.Floors[i]
		_, _ = fmt.Fprintf(&sb, "F%d ", i+1)
		if i == int(s.Elevator) {
			sb.WriteString("E ")
		} else {
			sb.WriteString(". ")
		}

		for idx := uint8(0); idx < 8; idx++ {
			if f.Chip&(1<<idx) != 0 {
				sb.WriteString(fmt.Sprintf("C%d ", idx+1))
			} else {
				sb.WriteString(" . ")
			}
			if f.RTG&(1<<idx) != 0 {
				sb.WriteString(fmt.Sprintf("R%d ", idx+1))
			} else {
				sb.WriteString(" . ")
			}
		}
		sb.WriteString(f.String())
		sb.WriteString("\n")
	}

	return sb.String()
}

func (s State) Copy() State {
	return State{
		Elevator: s.Elevator,
		Floors:   [4]Floor{s.Floors[0], s.Floors[1], s.Floors[2], s.Floors[3]},
	}
}

// Moves generates possible moves from the current state.
func (s State) Moves(from, to uint8) []Components {

	moves := s.Floors[from].Options(s.Floors[to])

	// todo: extra pruning

	return moves
}

func (s State) NextFloors() []uint8 {
	var floors []uint8
	if s.Elevator > 0 {
		floors = append(floors, s.Elevator-1)
	}
	if s.Elevator < 3 {
		floors = append(floors, s.Elevator+1)
	}
	return floors
}

func (s State) Next(floor uint8, move Components) State {
	next := s.Copy()
	next.Elevator = floor
	next.Floors[floor].Add(move)
	next.Floors[s.Elevator].Remove(move)
	return next
}

func (s State) Valid() bool {

	var c [2]uint16

	for _, f := range s.Floors {
		c[0] += uint16(f.Chip)
		c[1] += uint16(f.RTG)
		if !f.IsSafe() {
			return false
		}
	}

	if c[0] != c[1] {
		return false
	}

	return true
}
