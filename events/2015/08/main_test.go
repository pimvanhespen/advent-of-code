package main

import (
	"github.com/pimvanhespen/aoc/2015/days/08/xstring"
	"testing"
)

func Test_part1(t *testing.T) {
	sss := []string{
		`""`,
		`"abc"`,
		`"aaa\"aaa"`,
		`"\x27"`,
	}
	n := part1(sss)

	if n != 12 {
		t.Errorf("part1() = %v, want %v", n, 12)
	}
}

func Test_part2(t *testing.T) {
	lines := []string{
		`""`,
		`"abc"`,
		`"aaa\"aaa"`,
		`"\x27"`,
	}
	n := part2(lines)

	if n != 19 {
		t.Errorf("part1() = %v, want %v", n, 19)
	}
}

func Test_unescape(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "empty string",
			args: args{
				s: `""`,
			},
			want: "",
		},
		{
			name: "abc",
			args: args{
				s: `"abc"`,
			},
			want: "abc",
		},
		{
			name: `"aaa\"aaa"`,
			args: args{
				s: `"aaa\"aaa"`,
			},
			want: `aaa"aaa`,
		},
		{
			name: "hex",
			args: args{
				s: `"\x27"`,
			},
			want: "'",
		},
		{
			name: "multi escape",
			args: args{
				s: `"\\\\"`,
			},
			want: `\\`,
		},
		{
			name: "escape slash, hex",
			args: args{
				s: `"\\\x27"`,
			},
			want: `\'`,
		},
		{
			name: "escape slash, hex, escape",
			args: args{
				s: `"\\\x27\\"`,
			},
			want: `\'\`,
		},
		{
			name: "escape slash, hex, escape, hex",
			args: args{
				s: `"\\\x27\\\x27"`,
			},
			want: `\'\'`,
		},
		{
			name: "escape slash, hex, escape, hex, escape",
			args: args{
				s: `"\\\x27\\\x27\\"`,
			},
			want: `\'\'\`,
		},
		{
			name: "demo",
			args: args{
				s: `"rjjkfh\x78cf\x2brgceg\"jmdyas\"\\xlv\xb6p"`,
			},
			want: `rjjkfhxcf+rgceg"jmdyas"\xlv¶p`,
		},
		{
			name: "demo",
			args: args{
				s: `"aaa\x272\"xa"`,
			},
			want: `aaa'2"xa`,
		},
		{
			name: "demo",
			args: args{
				s: `"\""`,
			},
			want: `"`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := xstring.Unquote(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("unescape() error = %v, wantErr %v", err, tt.wantErr)
			}

			if got != tt.want {
				t.Errorf("unescape() = %v, want %v", got, tt.want)
			}
		})
	}
}
