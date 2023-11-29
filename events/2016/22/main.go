package main

import (
	"fmt"
	"github.com/pimvanhespen/advent-of-code/pkg/algorithms/astar"
	"github.com/pimvanhespen/advent-of-code/pkg/aoc"
	"github.com/pimvanhespen/advent-of-code/pkg/datastructures/heap"
	"io"
	"regexp"
	"strconv"
)

type Node struct {
	X, Y             int
	Size, Used, Free int
	Use              int
}

func (n Node) String() string {
	name := fmt.Sprintf("node-x%d-y%d", n.X, n.Y)
	return fmt.Sprintf("/dev/grid/%-15s %3dT %3dT %3dT %3d%%", name, n.Size, n.Used, n.Free, n.Use)
}

func (n Node) IsEmpty() bool {
	return n.Used == 0
}

func (n Node) IsFull() bool {
	return n.Used == n.Size
}

func (n Node) IsSame(other Node) bool {
	return n.X == other.X && n.Y == other.Y
}

func (n Node) CanMoveTo(other Node) bool {
	if n.IsEmpty() {
		return false
	}

	if n.IsSame(other) {
		return false
	}

	return n.Used <= other.Free
}

type Input struct {
	Nodes []Node
}

func main() {
	event := aoc.New(2016, 22, parse)
	fmt.Println("1:", aoc.Must(event.Run(part1)))
	fmt.Println("2:", aoc.Must(event.Run(part2)))
}

var re = regexp.MustCompile(`^/dev/grid/node-x(\d+)-y(\d+) +(\d+)T +(\d+)T +(\d+)T +(\d+)%`)

func parse(r io.Reader) (Input, error) {
	nodes, err := aoc.ParseLines(r, func(line string) (Node, error) {
		if line[0] != '/' {
			return Node{}, aoc.IgnoreLine
		}

		var n Node
		matches := re.FindStringSubmatch(line)
		if len(matches) != 7 {
			return n, fmt.Errorf("invalid line: %s", line)
		}
		n.X = aoc.Must(strconv.Atoi(matches[1]))
		n.Y = aoc.Must(strconv.Atoi(matches[2]))
		n.Size = aoc.Must(strconv.Atoi(matches[3]))
		n.Used = aoc.Must(strconv.Atoi(matches[4]))
		n.Free = aoc.Must(strconv.Atoi(matches[5]))
		n.Use = aoc.Must(strconv.Atoi(matches[6]))

		return n, nil
	})
	if err != nil {
		return Input{}, err
	}
	return Input{Nodes: nodes}, nil
}

func part1(input Input) string {

	var count int

	for i := 0; i < len(input.Nodes); i++ {
		for j := i + 1; j < len(input.Nodes); j++ {
			if input.Nodes[i].CanMoveTo(input.Nodes[j]) || input.Nodes[j].CanMoveTo(input.Nodes[i]) {
				count++
			}
		}
	}

	return aoc.Result(count)
}

type State uint8

const (
	Empty State = '_'
	Used  State = '.'
	Full  State = '#'
)

type Grid [][]byte

var _ astar.Node = (*asNode)(nil)

type asNode struct {
	neighbours []astar.Neighbor
	coord      Coord
	value      State
}

func (a *asNode) Neighbors() []astar.Neighbor {
	return a.neighbours
}

func (a *asNode) Equals(node astar.Node) bool {
	other, ok := node.(*asNode)
	if !ok || other == nil {
		panic("invalid node type")
	}

	if a == nil {
		panic("self is nil")
	}

	return a.coord == other.coord
}

func part2(input Input) string {

	var width int
	{
		var x int
		for _, n := range input.Nodes {
			x = max(x, n.X)
		}
		width = x + 1
	}

	asNodes := make(map[Coord]*asNode)
	for _, n := range input.Nodes {
		var value State

		switch {
		case n.IsEmpty():
			value = Empty
		case n.Size >= 100:
			value = Full
		default:
			value = Used
		}

		asNodes[Coord{X: n.X, Y: n.Y}] = &asNode{
			coord: Coord{X: n.X, Y: n.Y},
			value: value,
		}
	}

	for _, n := range asNodes {
		for _, offset := range []Coord{{X: 0, Y: -1}, {X: 0, Y: 1}, {X: -1, Y: 0}, {X: 1, Y: 0}} {
			coord := Coord{X: n.coord.X + offset.X, Y: n.coord.Y + offset.Y}
			if node, ok := asNodes[coord]; ok {
				if node.value == Full {
					continue
				}

				n.neighbours = append(n.neighbours, astar.Neighbor{
					Node: node,
					Cost: 1,
				})
			}
		}
	}

	var begin, end astar.Node
	for _, n := range asNodes {
		if n.value == Empty {
			begin = n
			break
		}
	}

	{
		lastOnFirstRow := Coord{X: width - 2, Y: 0}
		f, ok := asNodes[lastOnFirstRow]
		if !ok {
			panic("not found")
		}
		end = f
	}

	// Assumes there is unobstructed path on x[0] from begin to end
	path := astar.AStar(heap.NewMin[float64, astar.Node](), begin, end, heuristic)

	return aoc.Result(len(path) + 5*(width-2))
}

func heuristic(a, b astar.Node) float64 {
	return float64(manhattan(a.(*asNode).coord, b.(*asNode).coord))
}

type Coord struct {
	X, Y int
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func manhattan(a, b Coord) int {
	return abs(a.X-b.X) + abs(a.Y-b.Y)
}
