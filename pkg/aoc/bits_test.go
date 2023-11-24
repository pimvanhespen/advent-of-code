package aoc

import (
	"fmt"
	"io"
	"math"
	"testing"
)

func Test_popcnt(t *testing.T) {
	type args struct {
		x uint8
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"0", args{0}, 0},
		{"1", args{1}, 1},
		{"2", args{2}, 1},
		{"3", args{3}, 2},
		{"4", args{4}, 1},
		{"5", args{5}, 2},
		{"6", args{6}, 2},
		{"7", args{7}, 3},
		{"8", args{8}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := popcnt(tt.args.x); got != tt.want {
				t.Errorf("popcnt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountBits(t *testing.T) {
	type args struct {
		n uint
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"0", args{0}, 0},
		{"maxuint", args{math.MaxUint}, 64},
		{"maxuint8", args{math.MaxUint8}, 8},
		{"maxuint16", args{math.MaxUint16}, 16},
		{"maxuint32", args{math.MaxUint32}, 32},
		{"maxuint64", args{math.MaxUint64}, 64},
		{"half uint8", args{1 << 8}, 1},
	}

	type callee struct {
		name string
		fn   func(uint) int
	}

	callees := []callee{
		{"CountBits", CountBits},
		{"CountBits2", CountBits2},
		{"ContBits3", ContBits3},
	}

	for _, c := range callees {
		for _, tt := range tests {
			t.Run(c.name+" "+tt.name, func(t *testing.T) {
				if got := c.fn(tt.args.n); got != tt.want {
					t.Errorf("CountBits() = %v, want %v", got, tt.want)
				}
			})
		}
	}
}

func BenchmarkCountBits(b *testing.B) {

	type item struct {
		name string
		fn   func(uint) int
	}

	items := []item{
		{"CountBits", CountBits},
		{"CountBits1", CountBits1},
		{"CountBits2", CountBits2},
		{"ContBits3", ContBits3},
	}

	var total uint64
	for _, sut := range items {
		b.Run(sut.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				total += uint64(sut.fn(uint(i)))
			}
		})
	}
	_, _ = fmt.Fprintln(io.Discard, total)
}

func CountBits1(n uint) int {

	switch {
	case n == 0: // 0
		return 0
	case n <= 0xFF: // 8-bit
		return popcnt(uint8(n))
	case n <= 0xFFFF: // 16-bit
		return popcnt(uint8(n>>0)) + popcnt(uint8(n>>8))
	case n < 0xFFFFFFFF: // 32-bit
		return popcnt(uint8(n>>0)) + popcnt(uint8(n>>8)) + popcnt(uint8(n>>16)) + popcnt(uint8(n>>24))
	default: // 64-bit
		return popcnt(uint8(n>>0)) +
			popcnt(uint8(n>>8)) +
			popcnt(uint8(n>>16)) +
			popcnt(uint8(n>>24)) +
			popcnt(uint8(n>>32)) +
			popcnt(uint8(n>>40)) +
			popcnt(uint8(n>>48)) +
			popcnt(uint8(n>>56))
	}
}

func CountBits2(n uint) int {
	var c int
	for n > 0 {
		c++
		n &= n - 1
	}
	return c
}

func ContBits3(n uint) int {
	var c int
	for i := 0; i < 64; i++ {
		if n&(1<<i) != 0 {
			c++
		}
	}
	return c
}
