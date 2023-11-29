package main

import (
	"fmt"
	"github.com/pimvanhespen/advent-of-code/pkg/aoc"
	"io"
	"math/bits"
	"strconv"
)

type Input struct {
	Number int
}

func main() {
	event := aoc.New(2016, 19, parse)
	fmt.Println("1:", aoc.Must(event.Run(part1)))
	fmt.Println("2:", aoc.Must(event.Run(part2)))
}

func parse(r io.Reader) (Input, error) {
	n, err := aoc.ParseInput(r, strconv.Atoi)
	if err != nil {
		return Input{}, err
	}
	return Input{Number: n}, nil
}

func WhiteElephant(n int) int {
	if n < 1 {
		return 0
	} else {
		high := bits.Len32(uint32(n))
		remainder := n - (1 << (high - 1))
		return 1 + 2*remainder
	}
}

func part1(input Input) string {
	return strconv.Itoa(WhiteElephant(input.Number))
}

func opposite(size, offset int) int {
	o := offset + (size / 2)
	if o >= size {
		o -= size
	}
	return o
}

func naive2(n int) int {
	players := make([]int, n)
	for i := range players {
		players[i] = i + 1
	}

	var offset int

	for len(players) > 1 {
		if offset >= len(players) {
			offset = 0
		}
		target := opposite(len(players), offset)
		//log.Printf("player %d steals from player %d", players[offset], players[target])

		if target > offset {
			offset++ // only increment if we're not removing the player before us
		}

		players = append(players[:target], players[target+1:]...)
	}

	//log.Printf("player %d wins", players[0])

	return players[0]
}

func largestPow3(n int) int {
	sum := 1
	pow := 0
	for sum*3 <= n {
		pow++
		sum *= 3
	}

	return sum
}

func fast2(n int) int {

	l := largestPow3(n)

	if n == l {
		return n
	} else if n < 2*l {
		return n - l
	} else {
		return l + 2*(n-2*l)
	}
}

func part2(input Input) string {
	return strconv.Itoa(fast2(input.Number))
}
