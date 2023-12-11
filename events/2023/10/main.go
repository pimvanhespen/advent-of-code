package main

import (
	"bytes"
	"fmt"
	"github.com/pimvanhespen/advent-of-code/pkg/aoc"
	"io"
	"log"
)

type Vec2 struct {
	X, Y int
}

type Input struct {
	Width, Height int
	Grid          []byte
}

func (i Input) Get(pos Vec2) (byte, bool) {
	if pos.X < 0 || pos.X >= i.Width || pos.Y < 0 || pos.Y >= i.Height {
		return 0, false
	}
	xy := pos.Y*i.Width + pos.X
	return i.Grid[xy], true
}

func (i Input) GetStart() Vec2 {
	start := bytes.IndexByte(i.Grid, 'S')
	if start == -1 {
		panic("no start found")
	}
	return Vec2{start % i.Width, start / i.Width}
}

func (i Input) Set(pos Vec2, b byte) {
	if pos.X < 0 || pos.X >= i.Width || pos.Y < 0 || pos.Y >= i.Height {
		return
	}
	xy := pos.Y*i.Width + pos.X
	i.Grid[xy] = b
}

func (i Input) String() string {
	var buf bytes.Buffer
	for y := 0; y < i.Height; y++ {
		buf.Write(i.Grid[y*i.Width : (y+1)*i.Width])
		buf.WriteByte('\n')
	}
	return buf.String()
}

func main() {
	event := aoc.New(2023, 10, parse)
	fmt.Println("1:", aoc.Must(event.Run(part1)))
	fmt.Println("2:", aoc.Must(event.Run(part2)))
}

func parse(r io.Reader) (Input, error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return Input{}, err
	}

	b = bytes.TrimSpace(b)
	width := bytes.IndexByte(b, '\n')
	height := bytes.Count(b, []byte{'\n'}) + 1

	return Input{
		Width:  width,
		Height: height,
		Grid:   bytes.ReplaceAll(b, []byte{'\n'}, nil),
	}, nil
}

func part1(input Input) string {

	loop := getLoop(input)
	return fmt.Sprint(len(loop) / 2)
}

type Grid struct {
	Width, Height int
	Grid          []byte
}

func NewGrid(width, height int) Grid {
	return Grid{
		Width:  width,
		Height: height,
		Grid:   make([]byte, width*height),
	}
}

func (g *Grid) Draw(pos Vec2, data [][]byte) {
	for y, row := range data {
		for x, b := range row {
			g.Set(Vec2{pos.X + x, pos.Y + y}, b)
		}
	}
}

func (g *Grid) Set(pos Vec2, b byte) {
	if pos.X < 0 || pos.X >= g.Width || pos.Y < 0 || pos.Y >= g.Height {
		return
	}
	xy := pos.Y*g.Width + pos.X
	g.Grid[xy] = b
}

func (g *Grid) Fill(pos Vec2, before byte, b byte) {
	if pos.X < 0 || pos.X >= g.Width || pos.Y < 0 || pos.Y >= g.Height {
		return
	}

	xy := pos.Y*g.Width + pos.X
	if g.Grid[xy] != before {
		return
	}

	g.Grid[xy] = b
	for _, dir := range Directions {
		g.Fill(Vec2{pos.X + dir.X, pos.Y + dir.Y}, before, b)
	}
}

func (g *Grid) Get(vec2 Vec2) byte {
	if vec2.X < 0 || vec2.X >= g.Width || vec2.Y < 0 || vec2.Y >= g.Height {
		return 0
	}
	return g.Grid[vec2.Y*g.Width+vec2.X]
}

func (g *Grid) String() string {
	var buf bytes.Buffer
	for y := 0; y < g.Height; y++ {
		buf.Write(g.Grid[y*g.Width : (y+1)*g.Width])
		buf.WriteByte('\n')
	}
	return buf.String()
}

func (g *Grid) InBounds(pos Vec2) bool {
	return pos.X >= 0 && pos.X < g.Width && pos.Y >= 0 && pos.Y < g.Height
}

func (g *Grid) Neighbours(vec2 Vec2) []byte {
	var nbrs []byte
	for y := -1 + vec2.Y; y <= 1+vec2.Y; y++ {
		for x := -1 + vec2.X; x <= 1+vec2.X; x++ {
			if x == vec2.X && y == vec2.Y {
				continue
			}
			if !g.InBounds(Vec2{x, y}) {
				continue
			}
			nbrs = append(nbrs, g.Get(Vec2{x, y}))
		}
	}
	return nbrs

}

func part2(input Input) string {

	// take the input, find the loop, then use a flash fill algorithm to define whatever is outside the loop
	// then we also know whatever is inside the loop

	loop := getLoop(input)
	m := make(map[Vec2]bool)
	for _, l := range loop {
		m[l] = true
	}

	var startShape Shape

	start := loop[0]

	log.Printf("start: %v", start)

	startShape = [][]byte{
		{'.', '.', '.'},
		{'.', 'S', '.'},
		{'.', '.', '.'},
	}

	// Define the start shape based on the connectors
	if b, ok := input.Get(Vec2{start.X, start.Y - 1}); ok && bytes.ContainsAny([]byte{b}, "|F7") {
		startShape[0][1] = 'S'
	}
	if b, ok := input.Get(Vec2{start.X, start.Y + 1}); ok && bytes.ContainsAny([]byte{b}, "|JL") {
		startShape[2][1] = 'S'
	}
	if b, ok := input.Get(Vec2{start.X - 1, start.Y}); ok && bytes.ContainsAny([]byte{b}, string("-LF")) {
		startShape[1][0] = 'S'
	}
	if b, ok := input.Get(Vec2{start.X + 1, start.Y}); ok && bytes.ContainsAny([]byte{b}, string("-J7")) {
		startShape[1][2] = 'S'
	}

	// upscale the image by 3x, redraw the loop, then use flashflood to fill the rest
	// upscale

	scaled := NewGrid(input.Width*3, input.Height*3)

	// redraw grid
	for y := 0; y < input.Height; y++ {
		for x := 0; x < input.Width; x++ {
			pos := Vec2{x * 3, y * 3}
			switch input.Grid[y*input.Width+x] {
			case 'F':
				scaled.Draw(pos, ShapeF)
			case '7':
				scaled.Draw(pos, Shape7)
			case 'J':
				scaled.Draw(pos, ShapeJ)
			case 'L':
				scaled.Draw(pos, ShapeL)
			case '|':
				scaled.Draw(pos, ShapePipe)
			case '-':
				scaled.Draw(pos, ShapeDash)
			case 'S':
				scaled.Draw(pos, startShape)
			default:
				scaled.Draw(pos, ShapeFloor)
			}
		}
	}

	// flashflood
	for x := 0; x < scaled.Width; x++ {
		scaled.Fill(Vec2{x, 0}, '.', 'O')
		scaled.Fill(Vec2{x, scaled.Height - 1}, '.', 'O')
	}

	for y := 0; y < scaled.Height; y++ {
		scaled.Fill(Vec2{0, y}, '.', 'O')
		scaled.Fill(Vec2{scaled.Width - 1, y}, '.', 'O')
	}

	for y := 0; y < scaled.Height; y++ {
		for x := 0; x < scaled.Width; x++ {
			c := scaled.Get(Vec2{x, y})
			if c == 'O' || c == 'S' || c == '.' {
				continue
			}
			nbrs := scaled.Neighbours(Vec2{x, y})
			if bytes.ContainsAny(nbrs, "O") {
				continue
			}
			scaled.Draw(Vec2{x, y}, ShapeFloor)
		}
	}

	// downscale again
	for y := 0; y < input.Height; y++ {
		for x := 0; x < input.Width; x++ {
			input.Set(Vec2{x, y}, scaled.Get(Vec2{x*3 + 1, y*3 + 1}))
		}
	}

	fmt.Println(input.String())

	// return inside
	return fmt.Sprint(bytes.Count(input.Grid, []byte{'.'}))
}

var (
	Up    = Vec2{X: 0, Y: -1}
	Right = Vec2{X: 1, Y: 0}
	Down  = Vec2{X: 0, Y: 1}
	Left  = Vec2{X: -1, Y: 0}

	Directions = []Vec2{Up, Right, Down, Left}
)

func getConnectors(pos Vec2, input Input) []Vec2 {
	var connectors []Vec2
	for _, dir := range Directions {
		next, ok := input.Get(Vec2{pos.X + dir.X, pos.Y + dir.Y})
		if !ok {
			continue
		}

		switch dir {
		case Up:
			if bytes.ContainsAny([]byte{'|', 'F', '7'}, string(next)) {
				connectors = append(connectors, dir)
			}
		case Right:
			if bytes.ContainsAny([]byte{'-', 'J', '7'}, string(next)) {
				connectors = append(connectors, dir)
			}
		case Down:
			if bytes.ContainsAny([]byte{'|', 'J', 'L'}, string(next)) {
				connectors = append(connectors, dir)
			}
		case Left:
			if bytes.ContainsAny([]byte{'-', 'F', 'L'}, string(next)) {
				connectors = append(connectors, dir)
			}
		}
	}
	for i, c := range connectors {
		connectors[i] = Vec2{pos.X + c.X, pos.Y + c.Y}
	}
	return connectors
}

func getLoop(input Input) []Vec2 {
	start := input.GetStart()
	connectors := getConnectors(start, input)

	seen := make(map[Vec2]bool)
	seen[start] = true
	seen[connectors[0]] = true

	path := []Vec2{start, connectors[0]}
	for {
		links := getLinks(path[len(path)-1], input)
		if len(links) != 2 {
			panic("not 2 links")
		}
		before := len(path)
		for _, link := range links {
			if !seen[link] {
				seen[link] = true
				path = append(path, link)
				break
			}
		}
		if before == len(path) {
			break
		}
	}

	return path
}

func getLinks(curr Vec2, input Input) []Vec2 {
	b, ok := input.Get(curr)
	if !ok {
		panic("no byte")
	}

	var links []Vec2

	switch b {
	case 'F':
		links = append(links, Down, Right)
	case '7':
		links = append(links, Left, Down)
	case 'J':
		links = append(links, Left, Up)
	case 'L':
		links = append(links, Up, Right)
	case '|':
		links = append(links, Up, Down)
	case '-':
		links = append(links, Left, Right)
	}

	for i, link := range links {
		links[i] = Vec2{curr.X + link.X, curr.Y + link.Y}
	}

	return links
}

type Shape [][]byte

var (
	ShapeF = Shape{
		{'.', '.', '.'},
		{'.', 'F', 'F'},
		{'.', 'F', '.'},
	}
	Shape7 = Shape{
		{'.', '.', '.'},
		{'7', '7', '.'},
		{'.', '7', '.'},
	}
	ShapeJ = Shape{
		{'.', 'J', '.'},
		{'J', 'J', '.'},
		{'.', '.', '.'},
	}
	ShapeL = Shape{
		{'.', 'L', '.'},
		{'.', 'L', 'L'},
		{'.', '.', '.'},
	}
	ShapePipe = Shape{
		{'.', '|', '.'},
		{'.', '|', '.'},
		{'.', '|', '.'},
	}
	ShapeDash = Shape{
		{'.', '.', '.'},
		{'=', '=', '='},
		{'.', '.', '.'},
	}
	ShapeFloor = Shape{
		{'.', '.', '.'},
		{'.', '.', '.'},
		{'.', '.', '.'},
	}
)
