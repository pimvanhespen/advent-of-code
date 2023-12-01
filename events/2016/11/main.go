package main

import (
	"fmt"
	"github.com/pimvanhespen/advent-of-code/pkg/aoc"
	"github.com/pimvanhespen/advent-of-code/pkg/datastructures/heap"
	"io"
	"math/bits"
	"os"
	"regexp"
	"runtime/pprof"
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

type Route struct {
	State State
	Steps uint8
}

func solve(initial State) int {

	f, err := os.Create("cpu.pprof")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if err = pprof.StartCPUProfile(f); err != nil {
		panic(err)
	}
	defer pprof.StopCPUProfile()

	// Goal is to move all components to the top floor
	// We can only move 2 components at a time

	queue := heap.NewMin[int, Route](heap.WithSize(1 << 16))
	queue.Push(Route{State: initial, Steps: 0}, 0)

	seen := make(map[State]struct{}, 1<<8)
	least := uint8(255)

	// BFS
	for !queue.Empty() {
		route := queue.Pop()

		if route.Steps > least {
			continue
		}

		state := route.State

		floors := state.NextFloors()
		for _, targetFloor := range floors {

			curr := state.Floors[state.Elevator]
			target := state.Floors[targetFloor]

			options := curr.Options(target)
			for _, move := range options {

				// build new state
				next := state.Next(targetFloor, move)
				next = normalize(next) // this is a huge optimization

				if _, ok := seen[next]; ok {
					continue
				}

				if done(next.Floors) {
					least = min(least, route.Steps+1)
					continue
				}

				seen[next] = struct{}{}

				newRoute := Route{State: next, Steps: route.Steps + 1}
				queue.Push(newRoute, int(newRoute.Steps))
			}
		}
	}

	return int(least)
}

func part1(input Input) string {
	count := solve(input.State)
	return aoc.Result(count)
}

func part2(input Input) string {
	s := input.State
	s.Floors[0].Add(Components{3 << 5, 3 << 5})
	count := solve(s)
	return aoc.Result(count)
}

func done(floors [4]Floor) bool {
	return floors[0].IsEmpty() && floors[1].IsEmpty() && floors[2].IsEmpty()
}

// --- Day 11: Radioisotope Thermoelectric Generators ---

type Components [2]uint8

func (c Components) Size() int {
	return bits.OnesCount8(c[0]) + bits.OnesCount8(c[1])
}

func (c Components) String() string {
	return fmt.Sprintf("C%08b R%08b", c[0], c[1])
}

type Floor struct {
	Chip uint8 // support up to 8 chips
	RTG  uint8 // support up to 8 RTGs
}

func (f *Floor) String() string {
	return fmt.Sprintf("C%08b R%08b", f.Chip, f.RTG)
}

func (f *Floor) IsEmpty() bool {
	return f.Chip|f.RTG == 0
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
	opts := make([]Components, 0, 4) // pre-allocate 4 options

	appendSafe := func(c Components) {
		if !f.IsSafeWithout(c) {
			return
		}
		if !target.IsSafeWith(c) {
			return
		}
		opts = append(opts, c)
	}

	// reduce loop iterations by only looping up to the highest bit set
	merged := f.Chip | f.RTG | target.Chip | target.RTG
	var maximum uint8 = 1 << bits.Len8(merged)
	var minimum uint8 = 1 << bits.TrailingZeros8(merged)

	for i := minimum; i < maximum; i <<= 1 {
		if f.Chip&i != 0 {
			appendSafe([2]uint8{i, 0})
		}

		if f.RTG&i != 0 {
			appendSafe([2]uint8{0, i})
		}

		if f.Chip&i != 0 && f.RTG&i != 0 {
			appendSafe([2]uint8{i, i})
		}

		for j := i << 1; j < maximum; j <<= 1 {
			if f.Chip&i != 0 && f.RTG&j != 0 {
				appendSafe([2]uint8{i, j})
			}

			if f.Chip&i != 0 && f.Chip&j != 0 {
				appendSafe([2]uint8{i | j, 0})
			}

			if f.RTG&i != 0 && f.RTG&j != 0 {
				appendSafe([2]uint8{0, i | j})
			}
		}
	}

	return opts
}

func (f *Floor) OptionsWrite(target Floor, opts []Components) int {
	opts = opts[:0] // reset slice
	var written int

	appendSafe := func(c Components) {
		if !f.IsSafeWithout(c) {
			return
		}
		if !target.IsSafeWith(c) {
			return
		}
		opts = append(opts, c)
		written++
	}

	// reduce loop iterations by only looping up to the highest bit set
	merged := f.Chip | f.RTG | target.Chip | target.RTG
	var maximum uint8 = 1 << bits.Len8(merged)
	var minimum uint8 = 1 << bits.TrailingZeros8(merged)

	for i := minimum; i < maximum; i <<= 1 {
		if f.Chip&i != 0 {
			appendSafe([2]uint8{i, 0})
		}

		if f.RTG&i != 0 {
			appendSafe([2]uint8{0, i})
		}

		if f.Chip&i != 0 && f.RTG&i != 0 {
			appendSafe([2]uint8{i, i})
		}

		for j := i << 1; j < maximum; j <<= 1 {
			if f.Chip&i != 0 && f.RTG&j != 0 {
				appendSafe([2]uint8{i, j})
			}

			if f.Chip&i != 0 && f.Chip&j != 0 {
				appendSafe([2]uint8{i | j, 0})
			}

			if f.RTG&i != 0 && f.RTG&j != 0 {
				appendSafe([2]uint8{0, i | j})
			}
		}
	}

	return written
}

func isSafe(chip, rtg uint8) bool {
	if chip == 0 || rtg == 0 {
		return true
	}

	return chip&^rtg == 0
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
	return s // copy by value
}

// Moves generates possible moves from the current state.
func (s State) Moves(from, to uint8) []Components {
	return s.Floors[from].Options(s.Floors[to])
}

func (s State) NextFloors() []uint8 {
	if s.Elevator == 0 {
		return []uint8{1}
	}
	if s.Elevator == 3 {
		return []uint8{2}
	}
	return []uint8{s.Elevator - 1, s.Elevator + 1}
}

func (s State) Next(floor uint8, move Components) State {
	next := s.Copy()
	next.Elevator = floor
	next.Floors[floor].Add(move)
	next.Floors[s.Elevator].Remove(move)
	return next
}

func normalize(s State) State {
	// we don't care which exact components are on which floor only the pairs matter
	// e.g. C1 R1 C2 R2 is the same as C2 R2 C1 R1

	// we can use a lookup table to translate the components to a new position
	// without using a map
	var (
		translations [8]uint8
		seen         uint8
		counter      uint8
	)

	normal := State{Elevator: s.Elevator}
	for i, f := range s.Floors {
		start := bits.Len8(f.Chip | f.RTG)
		end := bits.TrailingZeros8(f.Chip | f.RTG)

		for idx := end; idx < start; idx++ {
			pos := uint8(1 << idx)
			if f.Chip&pos == 0 && f.RTG&pos == 0 {
				continue
			}

			// check if we've seen this component before
			if seen&pos == 0 {
				seen |= pos
				translations[idx] = 1 << counter
				counter++
			}

			// get the normal position of the component
			v := translations[idx]

			if f.Chip&pos != 0 {
				normal.Floors[i].Chip |= v
			}

			if f.RTG&pos != 0 {
				normal.Floors[i].RTG |= v
			}
		}
	}

	return normal
}
