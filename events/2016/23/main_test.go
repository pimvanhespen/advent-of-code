package main

import (
	"log"
	"strings"
	"testing"
)

const part1Example = `cpy 2 a
tgl a
tgl a
tgl a
cpy 1 a
dec a
dec a`

const day12part1 = `cpy 41 a
inc a
inc a
dec a
jnz a 2
dec a`

const demo = `cpy 1 a
cpy 1 b
cpy 26 d
jnz c 2
jnz 1 5
cpy 7 c
inc d
dec c
jnz c -2
cpy a c
inc a
dec b
jnz b -2
cpy c b
dec d
jnz d -6
cpy 17 c
cpy 18 d
inc a
dec d
jnz d -2
dec c
jnz c -5`

func Test_part1(t *testing.T) {
	tests := []struct {
		name string
		data string
		want string
	}{
		{"example", part1Example, "3"},
		{"example", day12part1, "42"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			in, err := parse(strings.NewReader(tt.data))
			if err != nil {
				t.Fatal(err)
			}

			if got := part1(in); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part1_ref(t *testing.T) {
	tests := []struct {
		name string
		data string
		want int
	}{
		{"example", part1Example, 3},
		{"example", day12part1, 42},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			in, err := parse(strings.NewReader(tt.data))
			if err != nil {
				t.Fatal(err)
			}

			r := NewComputer(in)
			r.Run()

			got := r.A

			if got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkSomething(b *testing.B) {
	in, err := parse(strings.NewReader(demo))
	if err != nil {
		b.Fatal(err)
	}

	var v int
	b.Run("Computer", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			r := NewComputer(in)
			r.Run()
			v = r.A
		}
	})

	b.Run("ComputerV1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			r := NewComputerV1(in)
			r.Run(nil)
			v = r.regs["a"]
		}
	})

	log.Println(v)
}
