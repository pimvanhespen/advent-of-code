package astar

import (
	"github.com/pimvanhespen/advent-of-code/pkg/aoc"
	"math"
)

type Neighbor struct {
	Node Node
	Cost float64
}

type Node interface {
	Neighbors() []Neighbor
	Equals(Node) bool
}

type Queue interface {
	Push(Node, float64)
	Pop() Node
	Len() int
}

type Heuristic[N aoc.Numeric] func(Node, Node) N

func AStar(frontier Queue, begin, end Node, heuristic Heuristic[float64]) []Node {
	frontier.Push(begin, 0)

	cameFrom := make(map[Node]Node)
	gScore := make(map[Node]float64)

	gScore[begin] = 0

	for frontier.Len() > 0 {
		current := frontier.Pop()
		if current.Equals(end) {
			return reconstructPath(cameFrom, current)
		}

		for _, neighbor := range current.Neighbors() {
			tentativeG := gScore[current] + neighbor.Cost

			neighborG, ok := gScore[neighbor.Node]
			if !ok {
				neighborG = math.MaxFloat64
			}

			if tentativeG < neighborG {
				gScore[neighbor.Node] = tentativeG
				cameFrom[neighbor.Node] = current

				fScore := tentativeG + heuristic(neighbor.Node, end)
				frontier.Push(neighbor.Node, fScore)
			}
		}
	}

	return nil
}

func reconstructPath(cameFrom map[Node]Node, current Node) []Node {
	var path []Node
	for current != nil {
		path = append(path, current)
		current = cameFrom[current]
	}
	return path
}
