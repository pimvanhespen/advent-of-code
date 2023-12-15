package main

import (
	"bytes"
	"fmt"
	"github.com/pimvanhespen/advent-of-code/pkg/aoc"
	"io"
	"log"
)

type Input [][]byte

func main() {
	event := aoc.New(2023, 15, parse)
	fmt.Println("1:", aoc.Must(event.Run(part1)))
	fmt.Println("2:", aoc.Must(event.Run(part2)))
}

func parse(r io.Reader) (Input, error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return Input{}, err
	}

	parts := bytes.Split(bytes.TrimSpace(b), []byte(","))
	return parts, nil
}

func hash(b []byte) byte {
	var h byte
	for _, c := range b {
		h += c
		h *= 17
	}
	return h
}

func part1(input Input) string {

	var sum uint64
	for _, b := range input {
		sum += uint64(hash(b))
	}
	return fmt.Sprintf("%d", sum)
}

type Lens struct {
	Label       []byte
	FocalLength uint8
}

type Box struct {
	Lenses []Lens
}

func (b *Box) Remove(label []byte) {
	for i := len(b.Lenses) - 1; i >= 0; i-- {
		if bytes.Equal(b.Lenses[i].Label, label) {
			b.Lenses = append(b.Lenses[:i], b.Lenses[i+1:]...)
		}
	}
}

func (b *Box) Add(l Lens) {
	for i, lens := range b.Lenses {
		if bytes.Equal(lens.Label, l.Label) {
			b.Lenses[i] = l
			return
		}
	}
	b.Lenses = append(b.Lenses, l)
}

func part2(input Input) string {

	var boxes [256]Box

	for _, b := range input {
		last := b[len(b)-1]
		switch last {
		case '-':
			label := b[:len(b)-1]
			boxes[hash(label)].Remove(label)

		case '1', '2', '3', '4', '5', '6', '7', '8', '9':
			label := b[:len(b)-2]
			boxes[hash(label)].Add(Lens{label, last - '0'})

		default:
			log.Fatalf("bad input: %s", b)
		}
	}

	var sum uint64
	for i, box := range boxes {
		for slot, lens := range box.Lenses {
			n := uint64(1+i) * uint64(1+slot) * uint64(lens.FocalLength)
			log.Printf("%s: %d (box %d) * %d (slot %d) * %d (focal length) = %d", lens.Label, 1+i, i, 1+slot, slot, lens.FocalLength, n)
			sum += n
		}
	}

	return aoc.Result(sum)
}
