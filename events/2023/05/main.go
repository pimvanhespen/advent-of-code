package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/pimvanhespen/advent-of-code/pkg/aoc"
)

type Input struct {
	Seeds []int
	Maps  []Map
}

func (i Input) String() string {
	var sb strings.Builder
	sb.WriteString("seeds: ")
	sb.WriteString(strings.Join(intsToAs(i.Seeds), " "))
	sb.WriteByte('\n')
	for _, m := range i.Maps {
		_, _ = fmt.Fprintf(&sb, "%s-to-%s map:\n", m.From, m.To)
		for _, s := range m.Scales {
			_, _ = fmt.Fprintf(&sb, "%d %d %d\n", s.Dst, s.Src, s.Len)
		}
		sb.WriteByte('\n')
	}
	return sb.String()
}

type Map struct {
	From, To string
	Scales   []Scale
}

func (m Map) Next(n int) int {
	for _, s := range m.Scales {
		if n >= s.Src && n < s.Src+s.Len {
			return s.Dst + (n - s.Src)
		}
	}
	return n
}

type Scale struct {
	Dst, Src, Len int
}

func main() {
	event := aoc.New(2023, 5, parse)
	fmt.Println("1:", aoc.Must(event.Run(part1)))
	fmt.Println("2:", aoc.Must(event.Run(part2)))
}

func parse(r io.Reader) (Input, error) {
	scanner := bufio.NewScanner(r)

	var input Input

	scanner.Scan()
	seeds := Ints(scanner.Text())
	scanner.Scan() // Skips line

	for scanner.Scan() {
		aToBreg.FindAllStringSubmatch(scanner.Text(), -1)
		from, to := aToBreg.FindStringSubmatch(scanner.Text())[1], aToBreg.FindStringSubmatch(scanner.Text())[2]

		m := Map{
			From: from,
			To:   to,
		}

		for scanner.Scan() {
			if scanner.Text() == "" {
				break
			}
			var scale Scale
			_, err := fmt.Sscanf(scanner.Text(), "%d %d %d", &scale.Dst, &scale.Src, &scale.Len)
			if err != nil {
				return Input{}, fmt.Errorf("scanning scale: %w", err)
			}
			m.Scales = append(m.Scales, scale)
		}

		input.Maps = append(input.Maps, m)
	}

	input.Seeds = seeds
	return input, nil
}

func part1(input Input) string {

	lo := math.MaxInt

	m := make(map[string]Map)

	for _, ma := range input.Maps {
		m[ma.From] = ma
	}

	for _, seed := range input.Seeds {
		typ := "seed"

		for typ != "location" {
			ma := m[typ]
			seed = ma.Next(seed)
			typ = ma.To
		}

		if seed < lo {
			lo = seed
		}
	}

	return aoc.Result(lo)
}

func part2(input Input) string {
	lo := math.MaxInt

	m := make(map[string]Map)

	for _, ma := range input.Maps {
		m[ma.From] = ma
	}

	seeds := make([]int, 0, len(input.Seeds))

	for i := 0; i < len(input.Seeds); i += 2 {
		begin, ln := input.Seeds[i], input.Seeds[i+1]
		for j := begin; j < begin+ln; j++ {
			seeds = append(seeds, j)
		}
	}

	for _, seed := range seeds {
		typ := "seed"

		for typ != "location" {
			ma := m[typ]
			seed = ma.Next(seed)
			typ = ma.To
		}

		if seed < lo {
			lo = seed
		}
	}

	return aoc.Result(lo)
}

func intsToAs(s []int) []string {
	ret := make([]string, len(s))
	for i, v := range s {
		ret[i] = strconv.Itoa(v)
	}
	return ret
}

var aToBreg = regexp.MustCompile(`(\w+)-to-(\w+) map:`)

var intreg = regexp.MustCompile(`(-?\d+)`)

func Ints(s string) []int {
	matches := intreg.FindAllStringSubmatch(s, -1)
	ret := make([]int, 0, len(matches))
	for _, m := range matches {
		i, _ := strconv.Atoi(m[1])
		ret = append(ret, i)
	}
	return ret
}
