package main

import (
	"fmt"
	"runtime"
	"strconv"
	"testing"
)

func Test_part1(t *testing.T) {

	type solver func(int) int

	solvers := map[string]solver{
		"naive":   solveNaive,
		"quick":   solveQuick,
		"quickV2": WhiteElephant,
	}

	tests := [][2]int{
		{1, 1},
		{2, 1},
		{3, 3},
		{4, 1},
		{5, 3},
		{6, 5},
		{7, 7},
		{8, 1},
		{9, 3},
		{10, 5},
		{11, 7},
		{12, 9},
		{13, 11},
		{14, 13},
		{15, 15},
		{16, 1},
		{17, 3},
		{18, 5},
		{19, 7},
		{20, 9},
		{21, 11},
		{22, 13},
		{23, 15},
		{24, 17},
		{25, 19},
		{26, 21},
		{27, 23},
		{28, 25},
		{29, 27},
		{30, 29},
		{31, 31},
		{32, 1},
	}

	for name, slv := range solvers {
		for _, tt := range tests {
			t.Run(fmt.Sprintf("[%s]%d=%d", name, tt[0], tt[1]), func(t *testing.T) {
				if got := slv(tt[0]); got != tt[1] {
					t.Errorf("part1() = %v, want %v", got, tt[1])
				}
			})
		}
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d=%d", tt[0], tt[1]), func(t *testing.T) {
			if got := part1(Input{tt[0]}); got != strconv.Itoa(tt[1]) {
				t.Errorf("part1() = %v, want %v", got, tt[1])
			}
		})
	}
}

func Benchmark_part1(b *testing.B) {

	var n int
	b.Run("naive", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			n += solveNaive(i + 1)
		}
	})

	b.Run("quick", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			n += solveQuick(i + 1)
		}
	})

	b.Run("quickV2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			n += WhiteElephant(i + 1)
		}
	})

	runtime.KeepAlive(n)
}

func Test_part2(t *testing.T) {
	type solver func(int) int

	solvers := map[string]solver{
		"naive": naive2,
		"quick": fast2,
	}

	tests := [][2]int{
		{1, 1},
		{2, 1},
		{3, 3},
		{4, 1},
		{5, 2},
		{6, 3},
		{7, 5},
		{8, 7},
		{9, 9},
		{10, 1},
	}
	for name, fn := range solvers {
		for _, tt := range tests {
			t.Run(fmt.Sprintf("[%s]%d=%d", name, tt[0], tt[1]), func(t *testing.T) {
				if got := fn(tt[0]); got != tt[1] {
					t.Errorf("naive2() = %v, want %v", got, tt[1])
				}
			})
		}
	}
}

func TestOutput2(t *testing.T) {

	nums := []int{1, 3, 9, 27, 81, 243, 729}

	dump := func(offset, n int) {
		switch {
		case offset < 10:
			fmt.Printf("%1d ", n)
		case offset < 100:
			fmt.Printf("%2d ", n)
		case offset < 1000:
			fmt.Printf("%3d ", n)
		default:
			fmt.Printf("%4d ", n)
		}
	}

	for i := 1; i <= nums[len(nums)-1]; i++ {
		dump(i, i)
	}
	fmt.Println()

	for _, n := range nums {
		for i := 1; i <= n; i++ {
			dump(i, naive2(i))
		}
		fmt.Println()
	}

	//const (
	//	total  = 225
	//	height = 15
	//	width  = total / height
	//)
	//
	//for i := 1; i <= height; i++ {
	//	for j := 0; j < width; j++ {
	//		fmt.Printf("%3d: %3d    ", i+j*height, naive2(i+j*height))
	//	}
	//	fmt.Println()
	//}
}

func Test_opposite(t *testing.T) {
	// size, offset, expected
	tests := [][3]int{
		{1, 0, 0},
		{2, 0, 1},
		{2, 1, 0},
		{3, 0, 1},
		{3, 1, 2},
		{3, 2, 0},
		{4, 0, 2},
		{4, 1, 3},
		{4, 2, 0},
		{4, 3, 1},
		{5, 0, 2},
		{5, 1, 3},
		{5, 2, 4},
		{5, 3, 0},
		{5, 4, 1},
		{6, 0, 3},
		{6, 1, 4},
		{6, 2, 5},
		{6, 3, 0},
		{6, 4, 1},
		{6, 5, 2},
		{7, 0, 3},
		{7, 1, 4},
		{7, 2, 5},
		{7, 3, 6},
		{7, 4, 0},
		{7, 5, 1},
		{7, 6, 2},
	}
	for _, tt := range tests {
		name := fmt.Sprintf("size=%d_player=%d_opposite=%d", tt[0], tt[1], tt[2])
		t.Run(name, func(t *testing.T) {
			got := opposite(tt[0], tt[1])
			if got != tt[2] {
				t.Errorf("opposite() = %v, want %v", got, tt[2])
			}
		})
	}
}

func Test_largestPow3(t *testing.T) {

	tests := []struct {
		arg  int
		want int
	}{
		{0, 0},
		{1, 0},
		{2, 0},
		{3, 1},
		{4, 1},
		{5, 1},
		{6, 1},
		{7, 1},
		{8, 1},
		{9, 2},
		{27, 3},
		{81, 4},
		{243, 5},
		{729, 6},
		{2187, 7},
		{6561, 8},
		{19683, 9},
		{59049, 10},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("log3(%d)=%d", tt.arg, tt.want), func(t *testing.T) {
			if got := largestPow3(tt.arg); got != tt.want {
				t.Errorf("largestPow3() = %v, want %v", got, tt.want)
			}
		})
	}
}
