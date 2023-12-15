package main

import (
	"fmt"
	"io"
	"reflect"
	"strings"
	"testing"

	"github.com/pimvanhespen/advent-of-code/pkg/aoc"
)

// exampleInput form the puzzle
const exampleInput = `O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....
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
			want: Input{},
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
			want:  "136",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("part1() = %s, want %s", got, tt.want)
			}
		})
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
			want:  "64",
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

func Test_flip270(t *testing.T) {
	type args struct {
		grid Grid
	}
	tests := []struct {
		name string
		args args
		want Grid
	}{
		{
			name: "example",
			args: args{
				grid: Grid{
					[]byte("...."),
					[]byte("####"),
				},
			},
			want: Grid{
				[]byte(".#"),
				[]byte(".#"),
				[]byte(".#"),
				[]byte(".#"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := flip270(tt.args.grid); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("flip270() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_flip180(t *testing.T) {
	type args struct {
		grid Grid
	}
	tests := []struct {
		name string
		args args
		want Grid
	}{
		{
			name: "example",
			args: args{
				grid: Grid{
					[]byte("...."),
					[]byte("####"),
				},
			},
			want: Grid{
				[]byte("####"),
				[]byte("...."),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := flip180(tt.args.grid); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("flip180() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_flip90(t *testing.T) {
	type args struct {
		grid Grid
	}
	tests := []struct {
		name string
		args args
		want Grid
	}{
		{
			name: "example",
			args: args{
				grid: Grid{
					[]byte("...."),
					[]byte("####"),
				},
			},
			want: Grid{
				[]byte("#."),
				[]byte("#."),
				[]byte("#."),
				[]byte("#."),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := flip90(tt.args.grid); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("flip90() = %v, want %v", got, tt.want)
			}
		})
	}
}

var cycleExamples = []string{
	exampleInput,
	`.....#....
....#...O#
...OO##...
.OO#......
.....OOO#.
.O#...O#.#
....O#....
......OOOO
#...O###..
#..OO#....`,
	`.....#....
....#...O#
.....##...
..O#......
.....OOO#.
.O#...O#.#
....O#...O
.......OOO
#..OO###..
#.OOO#...O`,
	`.....#....
....#...O#
.....##...
..O#......
.....OOO#.
.O#...O#.#
....O#...O
.......OOO
#...O###.O
#.OOO#...O`,
}

func Test_cycle(t *testing.T) {

	start, _ := parse(strings.NewReader(cycleExamples[0]))

	grid := Grid(start)

	for i := 1; i < len(cycleExamples); i++ {
		grid = cycle(grid)
		w, _ := parse(strings.NewReader(cycleExamples[i]))
		want := Grid(w)
		if !grid.Equal(want) {
			t.Errorf("cycle() = mismatch:\n%s", gridComparison(grid, want))
		}
	}
}

func Test_slideNorth(t *testing.T) {
	for _, s := range cycleExamples {
		parsed, _ := parse(strings.NewReader(s))
		grid := Grid(parsed)
		got := slideNorth(grid)
		want := slideBoulders(grid, North)

		if !got.Equal(want) {
			t.Errorf("slideNorth() = mismatch:\n%s", gridComparison(got, want))
		}
	}
}

func gridComparison(got, want Grid) string {
	var sb strings.Builder
	format := fmt.Sprintf("%%-%ds %%-%ds\n", len(got[0]), len(want[0]))
	_, _ = fmt.Fprintf(&sb, format, "got", "want")

	if len(got) != len(want) {
		_, _ = fmt.Fprintf(&sb, "len(got)=%d len(want)=%d\n", len(got), len(want))
	}

	if len(got[0]) != len(want[0]) {
		_, _ = fmt.Fprintf(&sb, "len(got[0])=%d len(want[0])=%d\n", len(got[0]), len(want[0]))
	}

	for y := 0; y < len(got); y++ {

		for x := 0; x < len(got[y]); x++ {

		}

		if y < len(got) {
			sb.WriteString(string(got[y]))
		} else {
			sb.WriteString(strings.Repeat(" ", len(got[0])))
		}

		sb.WriteString(" ")

		if y < len(want) {
			sb.WriteString(string(want[y]))
		} else {
			sb.WriteString(strings.Repeat(" ", len(want[0])))
		}

		for x := 0; x < min(len(got[0]), len(want[0])); x++ {
			if got[y][x] != want[y][x] {
				_, _ = fmt.Fprintf(&sb, " %d:%c:%c", x, got[y][x], want[y][x])
			}
		}

		sb.WriteByte('\n')
	}
	return sb.String()
}
