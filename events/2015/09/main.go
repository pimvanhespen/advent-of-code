package main

import (
	"bytes"
	"fmt"
	"github.com/pimvanhespen/aoc/2015/pkg/aoc"
	"io"
	"log"
	"math"
)

func main() {
	r, err := aoc.Get(9)
	if err != nil {
		log.Fatal(err)
	}

	input, err := parse(r)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Part 1:", part1(input))
	log.Println("Part 2:", part2(input))
}

type City string

type Route struct {
	From     City
	To       City
	Distance int
}

func parse(r io.Reader) ([]Route, error) {

	b, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	lines := bytes.Split(bytes.TrimSpace(b), []byte("\n"))

	routes := make([]Route, len(lines))

	for i, line := range lines {
		if string(line) == "" {
			continue
		}
		var from, to City
		var distance int

		_, err = fmt.Sscanf(string(line), "%s to %s = %d", &from, &to, &distance)
		if err != nil {
			return nil, fmt.Errorf("could not parse line %d: %q: %w", i, line, err)
		}

		routes[i] = Route{From: from, To: to, Distance: distance}
	}

	return routes, nil
}

func part1(input []Route) int {

	distances := map[City]map[City]int{}

	for _, r := range input {
		if _, ok := distances[r.From]; !ok {
			distances[r.From] = map[City]int{}
		}
		if _, ok := distances[r.To]; !ok {
			distances[r.To] = map[City]int{}
		}

		distances[r.From][r.To] = r.Distance
		distances[r.To][r.From] = r.Distance
	}

	cities := Cities{
		c: make([]City, 0, len(distances)),
		i: make(map[City]int, len(distances)),
	}

	for c := range distances {
		cities.c = append(cities.c, c)
		cities.i[c] = len(cities.c) - 1
	}

	result := math.MaxInt

	for _, c := range cities.c {
		distance := path(c, distances, cities, Set(0), Min)
		result = min(result, distance)
	}

	return result

}

func part2(input []Route) int {

	distances := map[City]map[City]int{}

	for _, r := range input {
		if _, ok := distances[r.From]; !ok {
			distances[r.From] = map[City]int{}
		}
		if _, ok := distances[r.To]; !ok {
			distances[r.To] = map[City]int{}
		}

		distances[r.From][r.To] = r.Distance
		distances[r.To][r.From] = r.Distance
	}

	cities := Cities{
		c: make([]City, 0, len(distances)),
		i: make(map[City]int, len(distances)),
	}

	for c := range distances {
		cities.c = append(cities.c, c)
		cities.i[c] = len(cities.c) - 1
	}

	result := 0

	for _, c := range cities.c {
		distance := path(c, distances, cities, Set(0), Max)
		result = max(result, distance)
	}

	return result

}

type Distances map[City]map[City]int
type Cities struct {
	c []City
	i map[City]int
}

func (c Cities) Len() int {
	return len(c.c)
}

func (c Cities) Index(city City) int {
	return c.i[city]
}

type PickFn func(int, int) int

func Min(a, b int) int {
	return min(a, b)
}

func Max(a, b int) int {
	return max(a, b)
}

func path(start City, distances Distances, cities Cities, visited Set, pick PickFn) int {

	visited = visited.Add(cities.Index(start))

	if visited == (1<<cities.Len())-1 {
		return 0
	}

	var result int
	if pick(0, 1) == 0 {
		result = math.MaxInt
	} else {
		result = 0
	}

	for _, c := range cities.c {
		if visited.Contains(cities.Index(c)) {
			continue
		}

		distance := distances[start][c] + path(c, distances, cities, visited, pick)

		result = pick(result, distance)
	}

	return result
}

type Set uint64

func (s Set) Add(i int) Set {
	return s | (1 << i)
}

func (s Set) Remove(i int) Set {
	return s &^ (1 << i)
}

func (s Set) Contains(i int) bool {
	return s&(1<<i) != 0
}
