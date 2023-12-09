package arithmatic

import "testing"

func TestLCM(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1, 2, 3", args{[]int{1, 2, 3}}, 6},
		{"2, 3, 4", args{[]int{2, 3, 4}}, 12},
		{"3, 4, 5", args{[]int{3, 4, 5}}, 60},
		{"4, 5, 6", args{[]int{4, 5, 6}}, 60},
		{"5, 6, 7", args{[]int{5, 6, 7}}, 210},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LCM(tt.args.nums...); got != tt.want {
				t.Errorf("LCM() = %v, want %v", got, tt.want)
			}
		})
	}
}
