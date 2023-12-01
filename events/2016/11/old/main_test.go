package main

import "testing"

func Test_safeGroup(t *testing.T) {
	type args struct {
		component []Component
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "single chip",
			args: args{
				component: []Component{
					{Type: ChipType, Element: 1},
				},
			},
			want: true,
		},
		{
			name: "single rtg",
			args: args{
				component: []Component{
					{Type: RTGType, Element: 1},
				},
			},
			want: true,
		},
		{
			name: "single chip and rtg",
			args: args{
				component: []Component{
					{Type: ChipType, Element: 1},
					{Type: RTGType, Element: 1},
				},
			},
			want: true,
		},
		{
			name: "two chips",
			args: args{
				component: []Component{
					{Type: ChipType, Element: 1},
					{Type: ChipType, Element: 2},
				},
			},
			want: true,
		},
		{
			name: "two rtgs",
			args: args{
				component: []Component{
					{Type: RTGType, Element: 1},
					{Type: RTGType, Element: 2},
				},
			},
			want: true,
		},
		{
			name: "different chip+rtg",
			args: args{
				component: []Component{
					{Type: ChipType, Element: 1},
					{Type: RTGType, Element: 2},
				},
			},
			want: false,
		},
		{
			name: "different chip+rtg",
			args: args{
				component: []Component{
					{Type: ChipType, Element: 1},
					{Type: RTGType, Element: 1},
					{Type: RTGType, Element: 2},
				},
			},
			want: false,
		},
		{
			name: "different chip+rtg",
			args: args{
				component: []Component{
					{Type: ChipType, Element: 1},
					{Type: ChipType, Element: 2},
					{Type: RTGType, Element: 1},
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := safeGroup(tt.args.component); got != tt.want {
				t.Errorf("safeGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}
