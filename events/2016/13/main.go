package main

import (
	"fmt"
	"github.com/pimvanhespen/advent-of-code/pkg/algorithms/astar"
	"github.com/pimvanhespen/advent-of-code/pkg/aoc"
	"github.com/pimvanhespen/advent-of-code/pkg/datastructures/heap"
	"io"
	"strconv"
)

const (
	Wall  = '#'
	Floor = '.'
	Path  = 'X'
)

type Input struct {
	MagicNumber int
	Target      Vec2
}

func main() {
	event := aoc.New(2016, 13, parse)
	fmt.Println("1:", aoc.Must(event.Run(part1)))
	fmt.Println("2:", aoc.Must(event.Run(part2)))
}

func parse(reader io.Reader) (Input, error) {
	b, err := io.ReadAll(reader)
	if err != nil {
		return Input{}, fmt.Errorf("read: %w", err)
	}
	b = b[:len(b)-1] // remove newline
	n, err := strconv.Atoi(string(b))
	if err != nil {
		return Input{}, fmt.Errorf("parse: %w", err)
	}

	return Input{
		MagicNumber: n,
		Target:      Vec2{31, 39},
	}, nil
}

func part1(input Input) string {

	m := NewMap(50, 50, isWall(input.MagicNumber))

	path := m.ShortestPath(Vec2{1, 1}, input.Target)

	return aoc.Result(len(path) - 1)
}

func part2(input Input) string {

	m := NewMap(50, 50, isWall(input.MagicNumber))

	options := m.Endpoints(Vec2{1, 1}, 50)

	return aoc.Result(len(options))
}

type Vec2 struct {
	X int
	Y int
}

func (c Vec2) String() string {
	return fmt.Sprintf("(%d,%d)", c.X, c.Y)
}

func ManhattanDistance(a, b Vec2) int {
	return abs(a.X-b.X) + abs(a.Y-b.Y)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type Map struct {
	Width  int
	Height int
	Data   [][]bool
}

func NewMap(width, height int, isWall func(int, int) bool) *Map {
	m := &Map{
		Width:  width,
		Height: height,
		Data:   make([][]bool, height),
	}
	for y := 0; y < height; y++ {
		m.Data[y] = make([]bool, width)
		for x := 0; x < width; x++ {
			m.Data[y][x] = isWall(x, y)
		}
	}
	return m
}

func (m *Map) Write(w io.Writer) (int, error) {
	return m.WritePath(w, nil)
}

func (m *Map) Neighbors(c Vec2) []Vec2 {
	var neighbors []Vec2
	for _, d := range []Vec2{{0, -1}, {0, 1}, {-1, 0}, {1, 0}} {
		x := c.X + d.X
		y := c.Y + d.Y
		if x < 0 || x >= m.Width || y < 0 || y >= m.Height {
			continue
		}
		if m.Data[y][x] {
			continue
		}
		neighbors = append(neighbors, Vec2{x, y})
	}
	return neighbors
}

func (m *Map) ShortestPath(from, to Vec2) []Vec2 {

	nodes := make(map[Vec2]astar.Node)

	// init nodes
	for y := 0; y < m.Height; y++ {
		for x := 0; x < m.Width; x++ {
			if m.Data[y][x] {
				continue
			}
			c := Vec2{x, y}
			nodes[c] = astar.Node(&node{
				coord:     c,
				neighbors: make([]astar.Neighbor, 0),
			})
		}
	}

	// link nodes
	for _, n := range nodes {
		for _, neighbor := range m.Neighbors(n.(*node).coord) {
			n.(*node).neighbors = append(n.(*node).neighbors, astar.Neighbor{
				Node: nodes[neighbor],
				Cost: 1,
			})
		}
	}

	// find path
	path := astar.AStar(
		heap.NewMin[float64, astar.Node](),
		nodes[from],
		nodes[to],
		func(a, b astar.Node) float64 {
			return float64(ManhattanDistance(a.(*node).coord, b.(*node).coord))
		},
	)

	if path == nil {
		return nil
	}

	var coords []Vec2
	for _, n := range path {
		coords = append(coords, n.(*node).coord)
	}
	return coords
}

func (m *Map) WritePath(stdout io.Writer, path []Vec2) (int, error) {
	lookup := make(map[Vec2]bool, len(path))
	for _, c := range path {
		lookup[c] = true
	}

	var wrote int
	for y := 0; y < m.Height; y++ {
		for x := 0; x < m.Width; x++ {
			var c byte
			if m.Data[y][x] {
				c = Wall
			} else if lookup[Vec2{x, y}] {
				c = Path
			} else {
				c = Floor
			}

			n, err := stdout.Write([]byte{c})
			wrote += n
			if err != nil {
				return wrote, err
			}
		}
		n, err := stdout.Write([]byte{'\n'})
		wrote += n
		if err != nil {
			return wrote, err
		}
	}

	return wrote, nil
}

func (m *Map) Endpoints(from Vec2, maxSteps int) []Vec2 {

	type State struct {
		Vec2
		Steps int
	}

	seen := make(map[Vec2]bool)

	var endpoints []Vec2

	queue := []State{{from, 0}}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if seen[current.Vec2] {
			continue
		}

		seen[current.Vec2] = true

		if current.Steps > maxSteps {
			continue
		}

		endpoints = append(endpoints, current.Vec2)

		if current.Steps == maxSteps {
			continue
		}

		for _, neighbor := range m.Neighbors(current.Vec2) {
			queue = append(queue, State{neighbor, current.Steps + 1})
		}
	}

	return endpoints
}

var _ astar.Node = &node{}

type node struct {
	coord     Vec2
	neighbors []astar.Neighbor
}

func (n *node) Neighbors() []astar.Neighbor {
	return n.neighbors
}

func (n *node) Equals(other astar.Node) bool {
	if other == nil {
		return false
	}
	o, ok := other.(*node)
	if !ok {
		return false
	}
	return n.coord == o.coord
}

func isWall(base int) func(int, int) bool {
	return func(x, y int) bool {
		b := uint(x*x + 3*x + 2*x*y + y + y*y + base)
		return aoc.CountBits(b)%2 == 1
	}
}
