package main

import (
	"fmt"
	"github.com/pimvanhespen/advent-of-code/pkg/aoc"
	"io"
	"math"
)

type Input []Race

type Race struct {
	Time     int
	Distance int
}

func main() {
	event := aoc.New(2023, 6, parse)
	fmt.Println("1:", aoc.Must(event.Run(part1)))
	fmt.Println("2:", aoc.Must(event.Run(part2)))
}

func parse(r io.Reader) (Input, error) {
	lines, err := aoc.ReadLines(r)
	if err != nil {
		return nil, err
	}

	if len(lines) != 2 {
		return nil, fmt.Errorf("expected 2 lines, got %d", len(lines))
	}

	times, err := aoc.Ints(lines[0])
	if err != nil {
		return nil, fmt.Errorf("parse times: %w", err)
	}

	distances, err := aoc.Ints(lines[1])
	if err != nil {
		return nil, fmt.Errorf("parse distances: %w", err)
	}

	if len(times) != len(distances) {
		return nil, fmt.Errorf("expected %d times, got %d", len(distances), len(times))
	}

	races := make([]Race, len(times))
	for i := range times {
		races[i] = Race{
			Time:     times[i],
			Distance: distances[i],
		}
	}

	return races, nil
}

func part1(input Input) string {
	sum := 1
	for _, r := range input {
		n := options(r.Time, r.Distance)
		if n == 0 {
			continue
		}
		sum *= n
	}
	return aoc.Result(sum)
}

func part2(input Input) string {

	race := input[0]
	for _, r := range input[1:] {
		race.Distance = race.Distance*padding(r.Distance) + r.Distance
		race.Time = race.Time*padding(r.Time) + r.Time
	}

	res := options(race.Time, race.Distance)
	return aoc.Result(res)
}

// roots returns the roots of the quadratic equation Ax^2 + Bx + C.
func roots(A, B, C float64) []float64 {
	D := B*B - 4*A*C
	if D < 0 {
		// We always expect TWO solutions, not zero or one
		panic(fmt.Errorf("no intersection: denominator = %.1f", D))
	}

	root := math.Sqrt(D)

	//  - B ± √(B² - 4AC)
	// -------------------
	//         2A
	opts := []float64{
		(-B + root) / (2 * A),
		(-B - root) / (2 * A),
	}

	if D > 0 {
		return opts
	}

	// D == 0, there is only one solution, so check which one it is
	if opts[0] == 0 {
		return opts[1:]
	}
	return opts[:1]
}

func padding(n int) int {
	return int(math.Pow10(1 + int(math.Log10(float64(n)))))
}

func distance(time, release int) int {
	n := release * (time - release)
	return n
}

func options(time, record int) int {

	// -x(time-x) = distance
	// -x(x-y) = z
	// -x^2 + xy - z = 0
	// A = -1, B = time, C = -distance
	var (
		A = float64(-1)
		B = float64(time)
		C = float64(-record)
	)

	bounds := roots(A, B, C)
	if len(bounds) != 2 {
		return 0
	}

	lo := int(bounds[0])
	hi := int(bounds[1])

	// Check that distance(lo) > distance, to ensure we beat the record
	if distance(time, lo) <= record {
		lo++
	}

	// Check that distance(hi) > distance, to ensure we beat the record
	if distance(time, hi) <= record {
		hi--
	}

	return 1 + hi - lo // inclusive
}
