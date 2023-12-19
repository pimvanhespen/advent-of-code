package main

import (
	"bytes"
	"fmt"
	"image"
	"io"
	"log"
	"slices"
	"strconv"
	"strings"

	"github.com/pimvanhespen/advent-of-code/pkg/aoc"
)

type Step struct {
	Dir   image.Point
	Len   int
	Color string
}

type Input []Step

func main() {
	event := aoc.New(2023, 18, parse)
	fmt.Println("1:", aoc.Must(event.Run(part1)))
	fmt.Println("2:", aoc.Must(event.Run(part2)))
}

func parse(r io.Reader) (Input, error) {
	return aoc.ParseLines(r, func(s string) (Step, error) {
		if s == "" {
			return Step{}, aoc.IgnoreLine
		}
		parts := strings.Split(s, " ")
		var dir image.Point
		switch parts[0] {
		case "R":
			dir = image.Pt(1, 0)
		case "L":
			dir = image.Pt(-1, 0)
		case "U":
			dir = image.Pt(0, 1)
		case "D":
			dir = image.Pt(0, -1)
		}
		n, err := strconv.Atoi(parts[1])
		if err != nil {
			return Step{}, err
		}
		return Step{
			Dir:   dir,
			Len:   n,
			Color: parts[2][1:8],
		}, nil
	})
}

func flashReplace(grid [][]byte, starts image.Point, from, to byte) {

	if grid[starts.Y][starts.X] != from {
		return
	}

	grid[starts.Y][starts.X] = to

	// floodfill
	for _, dir := range []image.Point{
		image.Pt(1, 0),
		image.Pt(-1, 0),
		image.Pt(0, 1),
		image.Pt(0, -1),
	} {
		next := starts.Add(dir)
		if next.X < 0 || next.X >= len(grid[0]) || next.Y < 0 || next.Y >= len(grid) {
			continue
		}
		flashReplace(grid, next, from, to)
	}
}

func solve(input Input) int {
	pts := make([]image.Point, 0, len(input))
	curr := image.Pt(0, 0)
	var perimeter int
	//pts = append(pts, curr)
	for _, step := range input {
		perimeter += step.Len
		addition := step.Dir.Mul(step.Len)
		next := curr.Add(addition)
		log.Printf("%v + %v = %v", curr, addition, next)
		pts = append(pts, next)
		curr = next
	}

	area := Area(pts) / 2

	return area + perimeter/2 + 1
}

func part1(input Input) string {
	return fmt.Sprintf("%d", solve(input))
}

func solveFill(input Input) int {
	minX, minY, maxX, maxY := 0, 0, 0, 0

	curr := image.Pt(0, 0)

	m := make(map[image.Point]bool)
	m[curr] = true

	for _, step := range input {
		next := curr.Add(step.Dir.Mul(step.Len))

		for i := 1; i <= step.Len; i++ {
			m[curr.Add(step.Dir.Mul(i))] = true
		}

		minX = min(minX, next.X)
		minY = min(minY, next.Y)
		maxX = max(maxX, next.X)
		maxY = max(maxY, next.Y)

		curr = next
	}

	// draw the map

	grid := make([][]byte, maxY-minY+1)
	for i := range grid {
		grid[i] = make([]byte, maxX-minX+1)
	}

	for y := range grid {
		for x := range grid[y] {
			if m[image.Pt(x+minX, y+minY)] {
				grid[y][x] = '#'
			} else {
				grid[y][x] = '.'
			}
		}
	}

	for _, point := range outerring(len(grid[0]), len(grid)) {
		flashReplace(grid, point, '.', 'O')
	}

	floors := 0
	inside := 0
	walls := 0

	for i := range grid {
		floors += bytes.Count(grid[i], []byte("O"))
		inside += bytes.Count(grid[i], []byte("."))
		walls += bytes.Count(grid[i], []byte("#"))
		fmt.Println(string(grid[i]))
	}

	fmt.Println("floors", floors)
	fmt.Println("inside", inside)
	fmt.Println("walls", walls)

	return len(grid)*len(grid[0]) - floors
}

func outerring(w, h int) []image.Point {
	var points []image.Point
	for x := 0; x < w; x++ {
		points = append(points, image.Pt(x, 0))
		points = append(points, image.Pt(x, h-1))
	}
	for y := 1; y < h-1; y++ {
		points = append(points, image.Pt(0, y))
		points = append(points, image.Pt(w-1, y))
	}
	return points
}

func Picks(points []image.Point) int {

	var perimeter int
	var curr, next image.Point
	for i := range points {
		next = points[(i+1)%len(points)]
		perimeter += diff(next, curr)
		curr = next
	}

	A := Area(points) / 2

	return A + perimeter/2 + 1
}

func diff(a, b image.Point) int {
	dx, dy := a.X-b.X, a.Y-b.Y
	return abs(dx) + abs(dy)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// Area calculates the area inside the polygon https://en.wikipedia.org/wiki/Shoelace_formula
func Area(points []image.Point) int {

	var sum int
	for i := range points {
		a, b := points[i], points[(i+1)%len(points)]
		diff := a.X*b.Y - a.Y*b.X
		sum += diff
	}

	if sum < 0 {
		slices.Reverse(points)
		return Area(points)
	}

	return sum
}

func parseStep(s string) Step {
	s = strings.TrimPrefix(s, "#")

	var dir image.Point
	switch s[len(s)-1] {
	case '0':
		dir = image.Pt(1, 0)
	case '1':
		dir = image.Pt(0, -1)
	case '2':
		dir = image.Pt(-1, 0)
	case '3':
		dir = image.Pt(0, 1)
	default:
		panic("invalid dir")
	}

	s = s[:len(s)-1]

	n, err := strconv.ParseInt(s, 16, 64)
	if err != nil {
		panic(err)
	}

	return Step{
		Dir: dir,
		Len: int(n),
	}
}

func part2(input Input) string {

	steps := make([]Step, len(input))

	for i, step := range input {
		steps[i] = parseStep(step.Color)
	}

	return fmt.Sprintf("%d", solve(steps))
}
