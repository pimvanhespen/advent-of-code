package main

import (
	"github.com/pimvanhespen/aoc/2015/days/08/xstring"
	"github.com/pimvanhespen/aoc/2015/pkg/aoc"
	"io"
	"log"
	"strings"
)

func main() {
	r, err := aoc.Get(8)
	if err != nil {
		log.Fatal(err)
	}

	lines, err := parse(r)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Part 1:", part1(lines))
	log.Println("Part 2:", part2(lines))
}

func parse(r io.Reader) ([]string, error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	return strings.Split(string(b), "\n"), nil
}

func part1(lines []string) int {

	var base int
	for _, line := range lines {
		base += len(line)
	}

	total := 0
	for _, line := range lines {
		if line == "" {
			continue
		}

		self, err := xstring.Unquote(line)
		if err != nil {
			log.Fatal(err)
		}

		ln := len([]rune(self))
		total += ln
	}

	return base - total
}

func part2(lines []string) int {
	var initial int
	for _, line := range lines {
		initial += len(line)
	}

	var total int

	for _, line := range lines {
		if line == "" {
			continue
		}

		var size int

		size += 2

		for _, c := range line {
			switch c {
			case '"', '\\':
				size += 2
			default:
				size += 1
			}
		}

		total += size
	}

	return total - initial
}
