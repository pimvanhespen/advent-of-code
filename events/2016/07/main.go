package main

import (
	"bytes"
	"fmt"
	"github.com/pimvanhespen/advent-of-code/pkg/aoc"
	"io"
)

type Input []IP

type IP struct {
	Super [][]byte
	Hyper [][]byte
}

func (p *IP) SupportsTLS() bool {
	for _, b := range p.Hyper {
		if isABBA(b) {
			return false
		}
	}

	for _, b := range p.Super {
		if isABBA(b) {
			return true
		}
	}

	return false
}

func isABBA(data []byte) bool {
	for len(data) >= 4 {
		if abba(data[:4]) {
			return true
		}
		data = data[1:]
	}
	return false
}

func abba(sub []byte) bool {
	return sub[0] == sub[3] && sub[1] == sub[2] && sub[0] != sub[1]
}

func main() {
	event := aoc.New(2016, 7, parse)
	fmt.Println("1:", aoc.Must(event.Run(part1)))
	fmt.Println("2:", aoc.Must(event.Run(part2)))
}

func parse(r io.Reader) (Input, error) {
	return aoc.ParseLines(r, parseIP)
}

func parseIP(v string) (IP, error) {

	data := []byte(v)

	var (
		ip    IP
		index int
	)

	for len(data) > 0 {
		index = bytes.IndexByte(data, '[')
		if index == -1 {
			// no more sequences
			ip.Super = append(ip.Super, data)
			break
		}

		ip.Super = append(ip.Super, data[:index])
		data = data[index+1:]

		index = bytes.IndexByte(data, ']')
		if index == -1 {
			return IP{}, fmt.Errorf("bad data: %s", v)
		}

		ip.Hyper = append(ip.Hyper, data[:index])
		data = data[index+1:]
	}

	return ip, nil
}

func part1(input Input) string {
	var count int
	for _, ip := range input {
		if ip.SupportsTLS() {
			count++
		}
	}
	return aoc.Result(count)
}

type Finder func([]byte) [][]byte

func all(seqs [][]byte, fn Finder) map[string]struct{} {
	seen := make(map[string]struct{})
	for _, s := range seqs {
		for _, sub := range fn(s) {
			key := string(invert(sub))
			if _, ok := seen[key]; !ok {
				seen[key] = struct{}{}
			}
		}
	}
	return seen
}

func invert(seq []byte) []byte {
	if len(seq) != 3 {
		panic(fmt.Sprintf("len(%q) != 3", string(seq)))
	}

	return []byte{seq[1], seq[0], seq[1]}
}

func aba(seq []byte) [][]byte {
	var matches [][]byte
	for len(seq) >= 3 {
		if seq[0] != seq[1] && seq[0] == seq[2] {
			matches = append(matches, seq[:3])
		}
		seq = seq[1:]
	}
	return matches
}

func (p *IP) SSL() bool {
	found := all(p.Super, aba)

	for _, h := range p.Hyper {
		for len(h) >= 3 {
			_, ok := found[string(h[:3])]
			if ok {
				return true
			}
			h = h[1:]
		}
	}

	return false
}

func part2(input Input) string {

	var count int

	for _, a := range input {
		v := a.SSL()
		if v {
			count++
		}
	}

	return aoc.Result(count)
}
