package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"slices"
	"strconv"
	"strings"

	"github.com/pimvanhespen/aoc/2015/pkg/aoc"
)

func main() {

	r, err := aoc.Get(02)
	if err != nil {
		panic(err)
	}

	boxes, err := parse(r)
	if err != nil {
		log.Fatal(err)
	}

	one := solve1(boxes)
	fmt.Println("one:", one)

	two := solve2(boxes)
	fmt.Println("two:", two)
}

func solve1(boxes []Box) int {
	total := 0
	for _, box := range boxes {
		total += box.GetArea() + box.SmallSide()
	}
	return total
}

func solve2(boxes []Box) int {
	var sum int
	for _, box := range boxes {
		sum += box.Volume() + box.Perimeter()
	}
	return sum
}

func Int(v string) int {
	i, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		panic(err)
	}
	return int(i)
}

func parse(r io.Reader) ([]Box, error) {

	b, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	rows := bytes.Split(b, []byte("\n"))

	boxes := make([]Box, len(rows))

	for i, row := range rows {
		if len(row) == 0 {
			break
		}
		fields := strings.Split(string(row), "x")
		boxes[i] = Box{
			Length: Int(fields[0]),
			Width:  Int(fields[1]),
			Height: Int(fields[2]),
		}
	}

	return boxes, nil
}

type Box struct {
	Length int
	Width  int
	Height int
}

func (b Box) GetArea() int {
	return 2*b.Length*b.Width + 2*b.Width*b.Height + 2*b.Height*b.Length
}

func (b Box) SmallSide() int {
	sides := []int{b.Length * b.Width, b.Width * b.Height, b.Height * b.Length}
	return slices.Min(sides)
}

func (b Box) Perimeter() int {
	sides := []int{b.Length, b.Width, b.Height}
	slices.Sort(sides)
	return sides[0]*2 + sides[1]*2
}

func (b Box) Volume() int {
	return b.Length * b.Width * b.Height
}
