package main

import (
	"fmt"
	"github.com/pimvanhespen/aoc/2015/pkg/aoc"
	"io"
	"slices"
	"strconv"
	"strings"
)

type Reindeer struct {
	name     string
	speed    int
	flyTime  int
	restTime int
	lap      int
}

func (r Reindeer) String() string {
	return fmt.Sprintf("%s: speed %d km/s, fly %d seconds, rest %d seconds.", r.name, r.speed, r.flyTime, r.restTime)
}

func (r Reindeer) Distance(raceTime int) int {

	var total int
	// How many full cycles
	fullCycles := raceTime / r.lap
	total += fullCycles * r.speed * r.flyTime

	// How many seconds in the last lap
	lastCycle := raceTime % r.lap
	if lastCycle > r.flyTime {
		lastCycle = r.flyTime
	}
	total += lastCycle * r.speed
	return total
}

func (r Reindeer) SpeedOn(time int) int {
	if time%r.lap < r.flyTime {
		return r.speed
	}
	return 0
}

func main() {
	reader, err := aoc.Get(14)
	if err != nil {
		panic(err)
	}

	reindeers, err := parse(reader)
	if err != nil {
		panic(err)
	}

	fmt.Println("Part 1:", part1(reindeers, 2503))
	fmt.Println("Part 2:", part2(reindeers, 2503))
}

func parse(reader io.Reader) ([]Reindeer, error) {

	b, err := io.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("could not read input: %w", err)
	}

	s := strings.TrimSpace(string(b))
	s = strings.NewReplacer(" can fly ", ",", " km/s for ", ",", " seconds, but then must rest for ", ",", " seconds.", "").Replace(s)

	reindeers := make([]Reindeer, 0)

	// Rudolph can fly 22 km/s for 8 seconds, but then must rest for 165 seconds.
	for _, line := range strings.Split(s, "\n") {
		if line == "" {
			continue
		}

		parts := strings.Split(line, ",")

		name := parts[0]
		speed, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, fmt.Errorf("could not parse speed: %w", err)
		}
		flyTime, err := strconv.Atoi(parts[2])
		if err != nil {
			return nil, fmt.Errorf("could not parse flyTime: %w", err)
		}
		restTime, err := strconv.Atoi(parts[3])
		if err != nil {
			return nil, fmt.Errorf("could not parse restTime: %w", err)
		}

		reindeers = append(reindeers, Reindeer{
			name:     name,
			speed:    speed,
			flyTime:  flyTime,
			restTime: restTime,
			lap:      flyTime + restTime,
		})
	}

	return reindeers, nil
}

func part1(reindeers []Reindeer, until int) int {
	var far int
	for _, r := range reindeers {
		far = max(far, r.Distance(until))
	}
	return far
}

// brute force, cuz it's fast enough (and 1am)
func part2(reindeers []Reindeer, until int) int {

	scores := make([]int, len(reindeers))
	position := make([]int, len(reindeers))

	for i := 0; i < until; i++ {
		var roundTop int
		for x, reindeer := range reindeers {
			position[x] += reindeer.SpeedOn(i)
			roundTop = max(roundTop, position[x])
		}

		for x, p := range position {
			if p < roundTop {
				continue
			}
			scores[x]++
		}
	}

	return slices.Max(scores)
}
