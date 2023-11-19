package main

import (
	"fmt"
	"io"
	"regexp"
	"sort"
	"strconv"

	"github.com/pimvanhespen/advent-of-code/pkg/aoc"
)

type Input []Room

type Room struct {
	SectorID int
	Name     []byte
	Checksum [5]byte
}

func (r Room) String() string {
	return fmt.Sprintf("%s-%d[%s]", r.Name, r.SectorID, r.Checksum)
}

func (r Room) Valid() bool {
	return r.Checksum == calculateChecksum(r.Name)
}

func (r Room) Decrypt() []byte {
	dec := make([]byte, len(r.Name))
	for i, c := range r.Name {
		if c == '-' {
			dec[i] = ' '
		} else {
			dec[i] = shift(c, r.SectorID)
		}
	}

	return dec
}

func shift(c byte, n int) byte {
	c = c - 'a'       // convert to number so we can apply modulo
	c += byte(n % 26) // mod 26 to keep the value within bounds of byte
	c = c % 26        // mod 26 to keep in bounds [0-25]
	c = c + 'a'       // convert back to char
	return c
}

func main() {
	event := aoc.New(2016, 4, parse)
	fmt.Println("1:", aoc.Must(event.Run(part1)))
	fmt.Println("2:", aoc.Must(event.Run(part2)))
}

func parse(r io.Reader) (Input, error) {

	re := regexp.MustCompile(`([a-z-]+)-(\d+)\[([a-z]+)\]`)

	return aoc.ParseLines(r, func(line string) (Room, error) {
		mathches := re.FindAllStringSubmatch(line, -1)
		if len(mathches) != 1 {
			return Room{}, fmt.Errorf("failed to parse line: %s", line)
		}
		results := mathches[0]
		if len(results) != 4 {
			return Room{}, fmt.Errorf("failed to parse line: %s", line)
		}

		var checksum [5]byte
		copy(checksum[:], results[3])

		sid, err := strconv.Atoi(results[2])
		if err != nil {
			return Room{}, fmt.Errorf("failed to parse sector id: %s", results[1][0])
		}

		return Room{
			SectorID: sid,
			Name:     []byte(results[1]),
			Checksum: checksum,
		}, nil
	})
}

func calculateChecksum(name []byte) [5]byte {
	found := make(map[byte]byte)
	for _, c := range name {
		if c == '-' {
			continue
		}
		found[c]++
	}

	type pair struct {
		char      byte
		frequency byte
	}

	pairs := make([]pair, 0, len(found))
	for char, frequency := range found {
		pairs = append(pairs, pair{char, frequency})
	}

	// sort by frequency, then by char
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].frequency == pairs[j].frequency {
			return pairs[i].char < pairs[j].char
		}
		return pairs[i].frequency > pairs[j].frequency
	})

	var checksum [5]byte
	for i := 0; i < 5; i++ {
		checksum[i] = pairs[i].char
	}

	return checksum
}

func part1(input Input) string {
	var sum int
	for _, room := range input {
		if room.Checksum == calculateChecksum(room.Name) {
			sum += room.SectorID
		}
	}
	return fmt.Sprint(sum)
}

func part2(input Input) string {

	re := regexp.MustCompile(`north\s?pole`)

	for _, room := range input {
		if !room.Valid() {
			continue
		}

		name := room.Decrypt()

		if re.Match(name) {
			return aoc.Result(room.SectorID)
		}
	}

	return "no matches"
}
