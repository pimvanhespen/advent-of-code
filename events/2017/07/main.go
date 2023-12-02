package main

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/pimvanhespen/advent-of-code/pkg/aoc"
)

type Node struct {
	Name     string
	Weight   int
	Parent   *Node
	Children []*Node
}

func (n *Node) String() string {
	return fmt.Sprintf("%s (%d)", n.Name, n.Weight)
}

func (n *Node) TotalWeight() int {
	total := n.Weight
	for _, child := range n.Children {
		total += child.TotalWeight()
	}
	return total
}

type Input []*Node

func main() {
	event := aoc.New(2017, 7, parse)
	fmt.Println("1:", aoc.Must(event.Run(part1)))
	fmt.Println("2:", aoc.Must(event.Run(part2)))
}

func parse(r io.Reader) (Input, error) {
	nodes, err := aoc.ParseLines(r, func(line string) (*Node, error) {

		parts := strings.Split(line, " -> ")

		left := strings.Split(parts[0], " ")

		name := left[0]
		ws := left[1][1 : len(left[1])-1]

		weight, err := strconv.Atoi(ws)
		if err != nil {
			return nil, err
		}

		var children []*Node

		if len(parts) > 1 {
			names := strings.Split(parts[1], ", ")
			children = make([]*Node, len(names))
			for i, child := range names {
				children[i] = &Node{Name: child}
			}
		}

		return &Node{
			Name:     name,
			Weight:   weight,
			Parent:   nil,
			Children: children,
		}, nil
	})

	if err != nil {
		return nil, err
	}

	link(nodes)
	return nodes, nil
}

func link(input Input) {
	refs := make(map[string]*Node)
	for _, node := range input {
		refs[node.Name] = node
	}

	for _, node := range input {
		for i, child := range node.Children {
			cnode := refs[child.Name]
			node.Children[i] = cnode
			cnode.Parent = node
		}
	}
}

func root(input Input) *Node {
	for _, node := range input {
		if node.Parent == nil {
			return node
		}

		curr := node
		for curr.Parent != nil {
			curr = curr.Parent
		}
		return curr
	}
	panic("no root found")
}

func part1(input Input) string {
	return root(input).Name
}

func part2(input Input) string {
	base := root(input)
	unbalanced := findUnbalanced(base)

	if unbalanced == nil {
		panic("no unbalanced found")
	}

	var common int
	if unbalanced.Parent.Children[0] != unbalanced {
		common = unbalanced.Parent.Children[0].TotalWeight()
	} else {
		common = unbalanced.Parent.Children[1].TotalWeight()
	}

	result := common - (len(unbalanced.Children) * unbalanced.Children[0].TotalWeight())

	return strconv.Itoa(result)
}

func findUnbalanced(node *Node) *Node {

	// keep digging until we find the last unbalanced node in the tree (the root cause)
	for _, child := range node.Children {
		if unbalanced := findUnbalanced(child); unbalanced != nil {
			return unbalanced
		}
	}

	// if the node has less than 3 children, it can't be unbalanced (in this scenario)
	if len(node.Children) < 3 {
		return nil
	}

	// if we get here, we're at the node that has the unbalanced children
	var normal int
	if node.Children[0].TotalWeight() == node.Children[1].TotalWeight() {
		normal = node.Children[0].TotalWeight()
	} else {
		normal = node.Children[2].TotalWeight()
	}

	for _, child := range node.Children {
		if child.TotalWeight() != normal {
			return child
		}
	}

	return nil
}
