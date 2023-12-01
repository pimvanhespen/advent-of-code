package main

import (
	"fmt"
	"github.com/pimvanhespen/advent-of-code/pkg/aoc"
	"io"
	"log"
	"math"
	"regexp"
	"strings"
)

type Input struct {
	Elevator int
	Floors   []Floor
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

	elements := make(map[string]Element)

	floors, err := aoc.ParseLines(r, func(line string) (Floor, error) {

		f := Floor{}
		f.cs = NewSet[Component]()

		for _, match := range chipEx.FindAllStringSubmatch(line, -1) {

			el, ok := elements[match[1]]
			if !ok {
				elements[match[1]] = Element(len(elements))
				el = elements[match[1]]
			}

			f.cs.Add(Chip(el))
		}

		for _, match := range rtgEx.FindAllStringSubmatch(line, -1) {
			el, ok := elements[match[1]]
			if !ok {
				elements[match[1]] = Element(len(elements))
				el = elements[match[1]]
			}

			f.cs.Add(RTG(el))
		}

		return f, nil

	})
	if err != nil {
		return Input{}, err
	}
	return Input{Floors: floors}, nil
}

func part1(input Input) string {
	// Goal is to move all components to the top floor
	// We can only move 2 components at a time

	queue := NewQueue[State]()
	queue.Push(State{
		Elevator: 0,
		Floors:   input.Floors,
		Steps:    0,
	})

	least := math.MaxInt64

	seen := NewSet[string]()

	for queue.Len() > 0 {

		log.Println(queue.Len())

		state := queue.Pop()

		if state.Steps > least {
			continue
		}

		for _, targetFloor := range state.NextFloors() {

			curr := state.Floors[state.Elevator]
			target := state.Floors[targetFloor]

			// select all possible combinations of 2 components that mayb be moved to the next floor
			// select all possible combinations of 1 component that may be moved to the next floor

			for _, move := range curr.Moves(target) {

				// build new state

				next := state.Next(state.Elevator, targetFloor, move)

				switch {
				case next.Done():
					least = min(least, next.Steps)
				case seen.Contains(next.Key()):
					continue
				default:
					queue.Push(next)
				}
			}
		}
	}

	// select available floors

	return aoc.Result(least)
}

func part2(input Input) string {
	return "n/a"
}

// TODO: make this not a map? need an ordered set (for printing and caching)
type Floor struct {
	cs *Set[Component]
}

func (f Floor) String() string {
	var sb strings.Builder
	for c := range f.cs.m {
		sb.WriteString(c.String())
		sb.WriteString(", ")
	}
	return sb.String()
}

func (f Floor) Copy() Floor {
	return Floor{cs: f.cs.Copy()}
}

func (f Floor) CanLeave(components []Component) bool {
	next := f.cs.Copy()
	for _, c := range components {
		next.Remove(c)
	}
	return safeGroup(next.Items())
}

func (f Floor) CanAccept(component []Component) bool {
	next := f.cs.Copy()
	for _, c := range component {
		next.Add(c)
	}
	return safeGroup(next.Items())
}

// Moves returns all possible moves from this floor to the target floor
func (f Floor) Moves(target Floor) [][]Component {

	// generate permutations of 1 and 2 components
	var perms [][]Component

	for _, k := range f.cs.Items() {
		perms = append(perms, []Component{k})

		for _, k2 := range f.cs.Items() {
			if k == k2 {
				continue
			}
			perms = append(perms, []Component{k, k2})
		}
	}

	for i := len(perms) - 1; i >= 0; i-- {
		p := perms[i]
		if !safeGroup(p) {
			perms = append(perms[:i], perms[i+1:]...)
		} else if !f.CanLeave(p) {
			perms = append(perms[:i], perms[i+1:]...)
		} else if !target.CanAccept(p) {
			perms = append(perms[:i], perms[i+1:]...)
		}
	}

	return perms
}

func safeGroup(component []Component) bool {
	chips := NewSet[Element]()
	rtgs := NewSet[Element]()

	for _, c := range component {
		switch c.Type {
		case ChipType:
			chips.Add(c.Element)
		case RTGType:
			rtgs.Add(c.Element)
		}
	}

	if chips.Len() == 0 || rtgs.Len() == 0 {
		return true
	}

	// check whether all chips are protected by an RTG.
	return rtgs.Difference(chips).Len() == 0
}

type ComponentType uint8

const (
	RTGType ComponentType = iota
	ChipType
)

type Element uint8

type Component struct {
	Type    ComponentType
	Element Element
}

func (c Component) String() string {
	switch c.Type {
	case RTGType:
		return fmt.Sprintf("G%v", c.Element)
	case ChipType:
		return fmt.Sprintf("M%v", c.Element)
	default:
		panic("unknown component type")
	}
}

func Chip(el Element) Component {
	return Component{
		Type:    ChipType,
		Element: el,
	}
}

func RTG(el Element) Component {
	return Component{
		Type:    RTGType,
		Element: el,
	}
}

type State struct {
	Elevator int
	Floors   []Floor
	Steps    int
}

func (s State) Next(prev, floor int, move []Component) State {

	next := State{
		Elevator: floor,
		Floors:   make([]Floor, len(s.Floors)),
		Steps:    s.Steps + 1,
	}

	for i, f := range s.Floors {
		next.Floors[i] = f.Copy()
	}

	for _, c := range move {
		next.Floors[floor].cs.Add(c)
		next.Floors[prev].cs.Remove(c)
	}

	return next
}

func (s State) Done() bool {
	for i, f := range s.Floors {
		if i != len(s.Floors)-1 && f.cs.Len() > 0 {
			return false
		}
	}

	return true
}

func (s State) String() string {
	var sb strings.Builder
	for i, f := range s.Floors {
		sb.WriteString(fmt.Sprintf("F%d: %s\n", i+1, f))
	}
	return sb.String()
}

func (s State) NextFloors() []int {
	if s.Elevator == 0 {
		return []int{1}
	}

	if s.Elevator == len(s.Floors)-1 {
		return []int{len(s.Floors) - 2}
	}

	return []int{s.Elevator - 1, s.Elevator + 1}
}

func (s State) Key() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("E%d", s.Elevator))
	for _, f := range s.Floors {
		sb.WriteString(fmt.Sprintf("F%d", f))
		for _, c := range f.cs.Items() {
			sb.WriteString(c.String())
		}
	}
	return sb.String()
}

// ---- SET ----

type Set[T comparable] struct {
	m map[T]struct{}
}

func NewSet[T comparable](items ...T) *Set[T] {
	m := make(map[T]struct{}, len(items))
	for _, item := range items {
		m[item] = struct{}{}
	}
	return &Set[T]{m: m}
}

func (s *Set[T]) Items() []T {
	items := make([]T, 0, s.Len())
	for item := range s.m {
		items = append(items, item)
	}
	return items
}

func (s *Set[T]) Add(v T) {
	if _, ok := s.m[v]; ok {
		return
	}
	s.m[v] = struct{}{}
}

func (s *Set[T]) Remove(v T) {
	delete(s.m, v)
}

func (s *Set[T]) Contains(v T) bool {
	_, ok := s.m[v]
	return ok
}

func (s *Set[T]) Intersect(other *Set[T]) *Set[T] {
	result := NewSet[T]()
	for v := range s.m {
		if other.Contains(v) {
			result.Add(v)
		}
	}
	return result
}

func (s *Set[T]) Union(other *Set[T]) *Set[T] {
	result := NewSet[T]()
	for v := range s.m {
		result.Add(v)
	}
	for v := range other.m {
		result.Add(v)
	}
	return result
}

func (s *Set[T]) Difference(other *Set[T]) *Set[T] {
	result := NewSet[T]()
	for v := range s.m {
		if !other.Contains(v) {
			result.Add(v)
		}
	}
	return result
}

func (s *Set[T]) Len() int {
	return len(s.m)
}

func (s *Set[T]) Copy() *Set[T] {
	m := make(map[T]struct{}, len(s.m))
	for k, v := range s.m {
		m[k] = v
	}
	return &Set[T]{m: m}
}

type Queue[T any] struct {
	items []T
}

// Make Prio Queue

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

func (q *Queue[T]) Push(item T) {
	q.items = append(q.items, item)
}

func (q *Queue[T]) Pop() T {
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

func (q *Queue[T]) Len() int {
	return len(q.items)
}
