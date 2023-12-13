package main

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
	"strings"
	"testing"

	"github.com/pimvanhespen/advent-of-code/pkg/aoc"
)

// exampleInput form the puzzle
const exampleInput = `???.### 1,1,3
.??..??...?##. 1,1,3
?#?#?#?#?#?#?#? 1,3,1,6
????.#...#... 4,1,1
????.######..#####. 1,6,5
?###???????? 3,2,1
`

func Test_parse(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name string
		args args
		want Input
	}{
		{
			name: "example",
			args: args{
				r: strings.NewReader(exampleInput),
			},
			want: Input{
				{Row: []byte("???.###"), Broken: []int{1, 1, 3}},
				{Row: []byte(".??..??...?##."), Broken: []int{1, 1, 3}},
				{Row: []byte("?#?#?#?#?#?#?#?"), Broken: []int{1, 3, 1, 6}},
				{Row: []byte("????.#...#..."), Broken: []int{4, 1, 1}},
				{Row: []byte("????.######..#####."), Broken: []int{1, 6, 5}},
				{Row: []byte("?###????????"), Broken: []int{3, 2, 1}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parse(tt.args.r)
			if err != nil {
				t.Fatalf("parse() error = %v", err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parse() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input Input
		want  string
	}{
		{
			name:  "example",
			input: aoc.Must(parse(strings.NewReader(exampleInput))),
			want:  "21",
		},
	}

	type tester struct {
		name string
		fn   func(Input) string
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("part1() = %s, want %s", got, tt.want)
			}
		})
	}
}

func Test_permute_nums(t *testing.T) {
	tests := []struct {
		line string
		want int
	}{
		{line: "???.### 1 1 3", want: 1},
		{line: ".??..??...?##. 1,1,3", want: 4},
		{line: "?#?#?#?#?#?#?#? 1,3,1,6", want: 1},
		{line: "????.#...#... 4,1,1", want: 1},
		{line: "????.######..#####. 1,6,5", want: 4},
		{line: "?###???????? 3,2,1", want: 10},
	}

	testers := []struct {
		name string
		fn   func(permutation) []permutation
	}{
		{name: "permute", fn: permute},
		{name: "permute2", fn: permute2},
		{name: "permute3", fn: permute3mask},
	}

	for _, tester := range testers {
		for _, tt := range tests {
			t.Run(fmt.Sprintf("%s-%s", tester.name, tt.line), func(t *testing.T) {
				n := bytes.IndexByte([]byte(tt.line), ' ')
				line := tt.line[:n]
				nums, err := aoc.Ints(tt.line[n+1:])
				if err != nil {
					t.Fatalf("parse broken: %v", err)
				}

				perm := permutation{
					locked: nil,
					rest:   []byte(line),
					broken: nums,
				}

				got := tester.fn(perm)

				if len(got) != tt.want {
					t.Fatalf("got %d permutations, want %d", len(got), tt.want)
				}
			})
		}
	}
}

func Test_part2(t *testing.T) {
	tests := []struct {
		name  string
		input Input
		want  string
	}{
		{
			name:  "example",
			input: aoc.Must(parse(strings.NewReader(exampleInput))),
			want:  "0",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("part2() = %s, want %s", got, tt.want)
			}
		})
	}
}

func Test_permute(t *testing.T) {
	type testcase struct {
		input permutation
		want  []permutation
	}

	tests := []testcase{
		{
			input: permutation{nil, []byte("???"), []int{1}},
			want: []permutation{
				{locked: []byte("#..")},
				{locked: []byte(".#.")},
				{locked: []byte("..#")},
			},
		},
		{
			input: permutation{rest: []byte("???.###"), broken: []int{1, 1, 3}},
			want: []permutation{
				{locked: []byte("#.#.###")},
			},
		},
		{
			input: permutation{rest: []byte("?#?#?#?#?#?#?#?"), broken: []int{1, 3, 1, 6}},
			want: []permutation{
				{locked: []byte(".#.###.#.######")},
			},
		},
		{
			input: permutation{rest: []byte("????.#...#..."), broken: []int{4, 1, 1}},
			want: []permutation{
				{locked: []byte("####.#...#...")},
			},
		},
		{
			input: permutation{rest: []byte("?.#..???"), broken: []int{1, 1}},
			want: []permutation{
				{locked: []byte("#.#.....")},
				{locked: []byte("..#..#..")},
				{locked: []byte("..#...#.")},
				{locked: []byte("..#....#")},
			},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.input), func(t *testing.T) {
			got := permute(tt.input)
			if len(got) != len(tt.want) {
				t.Fatalf("permute() = %v, want %v", got, tt.want)
			}

			for i, want := range tt.want {
				g := got[i]
				if !reflect.DeepEqual(g.String(), want.String()) {
					t.Errorf("permute() = %v, want %v", g, want)
				}
			}
		})
	}
}

func Test_Unfold_Permute(t *testing.T) {
	type testcase struct {
		input Record
		want  int
	}

	tests := []testcase{
		{
			input: Record{
				Row:    []byte("???.###"),
				Broken: []int{1, 1, 3},
			},
			want: 1,
		},
		{
			input: Record{
				Row:    []byte(".??..??...?##."),
				Broken: []int{1, 1, 3},
			},
			want: 16384,
		},
		{
			input: Record{
				Row:    []byte("?###????????"),
				Broken: []int{3, 2, 1},
			},
			want: 506250,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.input), func(t *testing.T) {
			in := unfold(tt.input)

			got := permute3(in.Row, in.Broken)
			if got != tt.want {
				t.Fatalf("permute() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fits(t *testing.T) {
	type args struct {
		row    []byte
		index  int
		broken int
	}
	tests := []struct {
		args args
		want bool
	}{
		{args: args{[]byte("???.###"), 0, 1}, want: true},
		{args: args{[]byte("???.###"), 0, 3}, want: true},
		{args: args{[]byte("???.###"), 1, 2}, want: true},
		{args: args{[]byte("???.###"), 2, 2}, want: false},
		{args: args{[]byte("???.###"), 3, 1}, want: false},
		{args: args{[]byte("......."), 0, 1}, want: false},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s-%d-%d", tt.args.row, tt.args.index, tt.args.broken), func(t *testing.T) {
			if got := fits(tt.args.row, tt.args.index, tt.args.broken); got != tt.want {
				t.Errorf("fits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_size(t *testing.T) {
	tests := []struct {
		ints []int
		want int
	}{
		{ints: []int{1, 2, 3}, want: 8},
		{ints: []int{1, 1, 1}, want: 5},
		{ints: []int{1, 2}, want: 4},
		{ints: []int{1}, want: 1},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("size(%v) = %d", tt.ints, tt.want), func(t *testing.T) {
			if got := size(tt.ints); got != tt.want {
				t.Errorf("size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_permute(b *testing.B) {

	permutations := []permutation{
		{nil, []byte("???"), []int{1}},
		{nil, []byte("???.###"), []int{1, 1, 3}},
		{nil, []byte("?#?#?#?#?#?#?#?"), []int{1, 3, 1, 6}},
		{nil, []byte("????.#...#..."), []int{4, 1, 1}},
		{nil, []byte("????.######..#####."), []int{1, 6, 5}},
		{nil, []byte("?###????????"), []int{3, 2, 1}},
		{nil, []byte("?.#..???"), []int{1, 1}},
	}

	b.Run("example", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			permute(permutations[i%len(permutations)])
		}
	})

	b.Run("permute-3", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			permute3(permutations[i%len(permutations)].rest, permutations[i%len(permutations)].broken)
		}
	})
}

func Benchmark_permute_unfold(b *testing.B) {

	permutations := []permutation{
		{nil, []byte("???"), []int{1}},
		{nil, []byte("???.###"), []int{1, 1, 3}},
		{nil, []byte("?#?#?#?#?#?#?#?"), []int{1, 3, 1, 6}},
		{nil, []byte("????.#...#..."), []int{4, 1, 1}},
		{nil, []byte("????.######..#####."), []int{1, 6, 5}},
		{nil, []byte("?###????????"), []int{3, 2, 1}},
		{nil, []byte("?.#..???"), []int{1, 1}},
	}

	b.Run("example", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			pm := unfold(Record{
				Row:    permutations[i%len(permutations)].rest,
				Broken: permutations[i%len(permutations)].broken,
			})
			permute(permutation{nil, pm.Row, pm.Broken})
		}
	})

	b.Run("permute-3", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {

			pm := unfold(Record{
				Row:    permutations[i%len(permutations)].rest,
				Broken: permutations[i%len(permutations)].broken,
			})
			permute3(pm.Row, pm.Broken)
		}
	})

	b.Run("permute-4", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {

			m := make(map[key]int)
			pm := unfold(Record{
				Row:    permutations[i%len(permutations)].rest,
				Broken: permutations[i%len(permutations)].broken,
			})

			permuteCached(m, pm.Row, pm.Broken)
		}
	})

}

func Test_Permute2(t *testing.T) {
	type testcase struct {
		in   string
		want []string
	}
	tcs := []testcase{
		{
			in: ".?.????.#?????#?? 1 1 1 1 6 ",
			want: []string{
				".#.#.#..#.######.",
				".#.#.#..#..######",
				".#.#..#.#.######.",
				".#.#..#.#..######",
				".#..#.#.#.######.",
				".#..#.#.#..######",
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.in, func(t *testing.T) {
			n := bytes.IndexByte([]byte(tc.in), ' ')
			line := tc.in[:n]
			nums, err := aoc.Ints(tc.in[n+1:])
			if err != nil {
				t.Fatalf("parse broken: %v", err)
			}

			perm := permutation{
				locked: nil,
				rest:   []byte(line),
				broken: nums,
			}

			got := permute(perm)

			if len(got) != len(tc.want) {
				t.Log(got)
				t.Fatalf("got %d permutations, want %d", len(got), len(tc.want))
			}

			for i, want := range tc.want {
				if string(got[i].locked) != want {
					t.Errorf("got %s, want %s", got[i], want)
				}
			}
		})
	}
}

type permutation struct {
	locked []byte
	rest   []byte
	broken []int
}

func (p permutation) String() string {
	return fmt.Sprintf("%s:%s:%v", p.locked, p.rest, p.broken)
}

// naive approach + optimizations

func permute(base permutation) []permutation {

	if len(base.broken) == 0 {
		if len(base.rest) > 0 {
			// Make sure that none of the remaining bytes are #
			if bytes.ContainsAny(base.rest, "#") {
				return nil
			}

			base.locked = append(base.locked, bytes.Repeat([]byte{'.'}, len(base.rest))...)
			base.rest = nil
		}
		return []permutation{base}
	}

	var permutations []permutation
	for i := 0; i < len(base.rest); i++ {
		if base.rest[i] == '.' {
			continue
		}

		// Does the allocation fit at this position?
		if !fits(base.rest, i, base.broken[0]) {
			if base.rest[i] == '#' || size(base.broken) >= len(base.rest[i:]) {
				break // permutation is invalid - stop
			}
			continue
		}

		// Does the remainder _POSSIBLY_ fit in the size of slice?
		if !(size(base.broken[1:]) < len(base.rest)) {
			break
		}

		// Lock bytes in place
		newSize := len(base.locked) + i + base.broken[0]

		// try to add a 'whitespace' unit
		if newSize < len(base.locked)+len(base.rest) {
			newSize++
		}

		locked := make([]byte, newSize)
		copy(locked, base.locked)
		for x := len(base.locked); x < newSize; x++ {
			if x >= len(base.locked)+i && x < len(base.locked)+i+base.broken[0] {
				locked[x] = '#'
			} else {
				locked[x] = '.'
			}
		}

		// adjust skipsize
		skip := newSize - len(base.locked)

		// Create new permutation
		next := permutation{
			locked: locked,
			rest:   base.rest[skip:],
			broken: base.broken[1:],
		}

		// Recurse
		permutations = append(permutations, permute(next)...)

		if base.rest[i] == '#' {
			break
		}
	}
	return permutations
}

func permute2(base permutation) []permutation {

	if len(base.broken) == 0 {
		if len(base.rest) > 0 {
			// Make sure that none of the remaining bytes are #
			if bytes.ContainsAny(base.rest, "#") {
				return nil
			}

			base.locked = append(base.locked, bytes.Repeat([]byte{'.'}, len(base.rest))...)
			base.rest = nil
		}
		return []permutation{base}
	}

	var permutations []permutation
	for i := 0; i < len(base.rest); i++ {
		if base.rest[i] == '.' {
			continue
		}

		// Does the allocation fit at this position?
		if !fits(base.rest, i, base.broken[0]) {
			if base.rest[i] == '#' || size(base.broken) >= len(base.rest[i:]) {
				break // permutation is invalid - stop
			}
			continue
		}

		// Does the remainder _POSSIBLY_ fit in the size of slice?
		if !(size(base.broken[1:]) < len(base.rest)) {
			break
		}

		// Lock bytes in place
		newSize := len(base.locked) + i + base.broken[0]

		// try to add a 'whitespace' unit
		if newSize < len(base.locked)+len(base.rest) {
			newSize++
		}

		// adjust skipsize
		skip := newSize - len(base.locked)

		// Create new permutation
		next := permutation{
			rest:   base.rest[skip:],
			broken: base.broken[1:],
		}

		// Recurse
		permutations = append(permutations, permute(next)...)

		if base.rest[i] == '#' {
			break
		}
	}
	return permutations
}

func permute3mask(base permutation) []permutation {
	n := permute3(base.rest, base.broken)
	return make([]permutation, n)
}

func permute3(row []byte, nums []int) int {

	if len(nums) == 0 {
		if len(row) > 0 {
			// Make sure that none of the remaining bytes are #
			if bytes.ContainsAny(row, "#") {
				return 0
			}
		}
		return 1
	}

	var permutations int
	for i := 0; i < len(row); i++ {
		if row[i] == '.' {
			continue
		}

		// Does the allocation fit at this position?
		if !fits(row[i:], 0, nums[0]) {
			if row[i] == '#' || size(nums) >= len(row[i:]) {
				break // permutation is invalid - stop
			}
			continue
		}

		// Does the remainder _POSSIBLY_ fit in the size of slice?
		if size(nums[1:]) >= len(row[i:]) {
			break
		}

		// Does the remainder _POSSIBLE_ align with the remainder of the ints?
		if !slightlyPossible(row[i+nums[0]:], nums[1:]) {
			break
		}

		// Lock bytes in place
		newSize := i + nums[0]

		// try to add a 'whitespace' unit
		if newSize < len(row) {
			newSize++
		}

		permutations += permute3(row[newSize:], nums[1:])
		// Recurse

		if row[i] == '#' {
			break
		}
	}
	return permutations
}

func sum(nums []int) int {
	var res int
	for _, n := range nums {
		res += n
	}
	return res
}

func slightlyPossible(row []byte, nums []int) bool {
	n := bytes.Count(row, []byte{'#'})

	lim := n + bytes.Count(row, []byte{'?'})

	if lim < sum(nums) {
		return false
	}

	return true
}
