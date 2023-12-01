package main

import (
	"bytes"
	"fmt"
	"github.com/pimvanhespen/advent-of-code/pkg/algorithms/astar"
	"github.com/pimvanhespen/advent-of-code/pkg/datastructures/heap"
	"io"
	"log"
	"math"
	"math/bits"

	"github.com/pimvanhespen/advent-of-code/pkg/aoc"
)

type Input [][]byte

func main() {
	event := aoc.New(2016, 24, parse)
	fmt.Println("1:", aoc.Must(event.Run(part1)))
	fmt.Println("2:", aoc.Must(event.Run(part2)))
}

func parse(r io.Reader) (Input, error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	b = bytes.TrimSpace(b)
	lines := bytes.Split(b, []byte("\n"))
	return lines, nil
}

func part1(input Input) string {
	distances := findDistances(input)

	all := uint64(1<<len(distances) - 1)

	res := shortestPath(distances, func(path Path) bool {
		return all == path.seen
	})
	return fmt.Sprint(res)
}

func part2(input Input) string {
	distances := findDistances(input)

	all := uint64(1<<len(distances) - 1)
	res := shortestPath(distances, func(path Path) bool {
		return all == path.seen && path.route[len(path.route)-1] == '0'
	})
	return fmt.Sprint(res)
}

type Vec2 struct {
	X, Y int
}

var _ astar.Node = &Node{}

type Node struct {
	Neighbours []astar.Neighbor
	Coord      Vec2
}

func (n *Node) Neighbors() []astar.Neighbor {
	return n.Neighbours
}

func (n *Node) Equals(node astar.Node) bool {
	other := node.(*Node)
	if other == nil {
		return false
	}

	return n.Coord == other.Coord
}

func Heuristic(a, b astar.Node) float64 {

	ca := a.(*Node).Coord
	cb := b.(*Node).Coord

	return float64(abs(ca.X-cb.X) + abs(ca.Y-cb.Y))
}

type Distances map[byte]map[byte]int

func findDistances(input Input) Distances {
	m := make(map[byte]Vec2)

	for y, line := range input {
		for x, c := range line {
			if c >= '0' && c <= '9' {
				m[c] = Vec2{X: x, Y: y}
			}
		}
	}

	// Use A* to find the shortest path
	// 1. Convert to aStar nodes
	// 2. Run A*
	// 3. Profit

	nodes := make(map[Vec2]*Node)

	for y, line := range input {
		for x, c := range line {
			if c == '#' {
				continue
			}

			coord := Vec2{X: x, Y: y}
			node := &Node{Coord: coord}
			nodes[coord] = node
		}
	}

	for _, node := range nodes {
		// Add neighbours
		for _, delta := range []Vec2{{X: 1, Y: 0}, {X: -1, Y: 0}, {X: 0, Y: 1}, {X: 0, Y: -1}} {
			coord := Vec2{X: node.Coord.X + delta.X, Y: node.Coord.Y + delta.Y}
			if neighbour, ok := nodes[coord]; ok {
				node.Neighbours = append(node.Neighbours, astar.Neighbor{
					Node: neighbour,
					Cost: 1,
				})
			}
		}
	}

	distances := make(Distances)

	for k := range m {
		distances[k] = make(map[byte]int)
	}

	for a, ca := range m {
		for b, cb := range m {
			if a == b {
				continue
			}

			if _, ok := distances[b][a]; ok {
				continue
			}

			queue := heap.NewMin[float64, astar.Node]()

			path := astar.AStar(queue, nodes[ca], nodes[cb], Heuristic)
			distances[a][b] = len(path) - 1
			distances[b][a] = len(path) - 1
		}
	}

	return distances
}

type CompleteFunc func(Path) bool

func isComplete(path string, distances Distances) bool {
	for k := range distances {
		if bytes.IndexByte([]byte(path), k) == -1 {
			return false
		}
	}
	return true
}

type Path struct {
	route string
	cost  int
	seen  uint64
}

func shortestPath(distances Distances, isComplete CompleteFunc) int {
	begin := Path{
		route: "0",
		cost:  0,
		seen:  1,
	}

	var longestEdge int
	for _, v := range distances {
		for _, d := range v {
			longestEdge = max(longestEdge, d)
		}
	}

	var all uint64
	{ //
		var n int
		for range distances {
			all |= 1 << n
			n++
		}
	}

	queue := heap.NewMin[int, Path]()
	queue.Push(begin, 0)

	least := Path{cost: math.MaxInt}

	for queue.Len() > 0 {
		path := queue.Pop()

		if path.cost >= least.cost {
			continue
		}

		// Prune paths that are too long while not visiting new nodes
		if bits.OnesCount64(path.seen&all) < len(path.route)/2 {
			continue
		}

		if isComplete(path) {
			// we have a path
			log.Println("found", path)
			if path.cost < least.cost {
				least = path
			}
			continue
		}

		for next := range distances {

			pos := path.route[len(path.route)-1]
			if next == pos {
				continue // don't stay in the same place
			}

			option := Path{
				route: path.route + string(next),
				cost:  path.cost + distances[pos][next],
				seen:  path.seen | (1 << (next - '0')),
			}

			missing := bits.OnesCount64(all ^ option.seen)
			value := option.cost + missing*longestEdge
			queue.Push(option, value)
		}
	}
	return least.cost
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}
