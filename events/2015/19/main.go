package main

import (
	"bufio"
	"fmt"
	"github.com/pimvanhespen/advent-of-code/pkg/aoc"
	"io"
	"strings"
)

type Replacement struct {
	from string
	to   string
}

type Data struct {
	replacements []Replacement
	molecule     string
}

func main() {
	reader, err := aoc.NewChallenge(2015, 19).Input()
	if err != nil {
		panic(err)
	}

	data, err := parse(reader)
	if err != nil {
		panic(err)
	}

	fmt.Println("Part 1:", solve1(data))
	//fmt.Println("Part 2:", solve2(data))
}

func parse(reader io.Reader) (Data, error) {
	data := Data{}

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return Data{}, err
		}

		text := scanner.Text()
		if text == "" {
			break
		}
		parts := strings.Split(scanner.Text(), " => ")
		data.replacements = append(data.replacements, Replacement{parts[0], parts[1]})
	}

	scanner.Scan()
	data.molecule = scanner.Text()

	return data, scanner.Err()
}

// todo: learn how to solve this
func solve1(_ Data) int {
	return -1
}
