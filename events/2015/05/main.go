package main

import (
	"fmt"
	"github.com/pimvanhespen/aoc/2015/pkg/aoc"
	"io"
	"log"
	"regexp"
	"strings"
)

var (
	exempt = regexp.MustCompile(`ab|cd|pq|xy`)
	vowels = regexp.MustCompile(`a|e|i|o|u`)
)

func main() {

	input, err := aoc.Get(5)
	if err != nil {
		log.Fatal(err)
	}

	strings, err := parse(input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part 1:", solve1(strings))
	fmt.Println("Part 2:", solve2(strings))

}

func parse(input io.Reader) ([]string, error) {
	b, err := io.ReadAll(input)
	if err != nil {
		return nil, err
	}

	return strings.Split(string(b), "\n"), nil
}

func isDouble(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			return true
		}
	}
	return false
}

func isDoubleWithSep(s string) bool {
	for i := 0; i < len(s)-2; i++ {
		if s[i] == s[i+2] {
			return true
		}
	}
	return false
}

func containsRepeatingPair(s string) bool {
	for i := 0; i < len(s)-3; i++ {
		if strings.Contains(s[i+2:], s[i:i+2]) {
			return true
		}
	}
	return false
}

func isNice(s string) bool {
	return isDouble(s) && !exempt.MatchString(s) && len(vowels.FindAllString(s, -1)) >= 3
}

func isNice2(s string) bool {
	return isDoubleWithSep(s) && containsRepeatingPair(s)
}

func solve1(strings []string) int {
	var nice int
	for _, s := range strings {
		if isNice(s) {
			nice++
		}
	}
	return nice
}

func solve2(strings []string) int {
	var nice int
	for _, s := range strings {
		if isNice2(s) {
			nice++
		}
	}
	return nice
}
