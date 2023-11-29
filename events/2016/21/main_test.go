package main

import (
	"reflect"
	"testing"
)

func TestRotateLeft_Execute(t *testing.T) {
	type fields struct {
		Steps int
	}
	type args struct {
		input []byte
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		{
			name: "rotateLeft 1",
			fields: fields{
				Steps: 1,
			},
			want: []byte("bcdea"),
		},
		{
			name: "rotateLeft 2",
			fields: fields{
				Steps: 2,
			},
			want: []byte("cdeab"),
		},
		{
			name: "rotateLeft -1",
			fields: fields{
				Steps: -1,
			},
			want: []byte("eabcd"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := RotateLeft{
				Steps: tt.fields.Steps,
			}
			data := []byte("abcde")
			r.Apply(data)
			if !reflect.DeepEqual(data, tt.want) {
				t.Errorf("Apply() = %s, want %s", data, tt.want)
			}
		})
	}
}

func TestMovePosition_Execute(t *testing.T) {
	type fields struct {
		From int
		To   int
	}
	type args struct {
		input []byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []byte
	}{
		{
			name: "move 1 to 4",
			fields: fields{
				From: 1,
				To:   4,
			},
			want: []byte("acdeb"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := MovePosition{
				From: tt.fields.From,
				To:   tt.fields.To,
			}
			data := []byte("abcde")
			m.Apply(data)
			if !reflect.DeepEqual(data, tt.want) {
				t.Errorf("Apply() = %s, want %s", data, tt.want)
			}
		})
	}
}

func TestMovePosition_Execute1(t *testing.T) {
	type fields struct {
		From int
		To   int
	}
	type args struct {
		input []byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []byte
	}{
		{
			name: "move 1 to 4",
			fields: fields{
				From: 1,
				To:   4,
			},
			want: []byte("acdeb"),
		},
		{
			name: "move 4 to 1",
			fields: fields{
				From: 4,
				To:   1,
			},
			want: []byte("aebcd"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := MovePosition{
				From: tt.fields.From,
				To:   tt.fields.To,
			}
			data := []byte("abcde")
			m.Apply(data)

			if !reflect.DeepEqual(data, tt.want) {
				t.Errorf("Apply() = %s, want %s", data, tt.want)
			}
		})
	}
}

func TestReversePositions_Execute(t *testing.T) {
	type fields struct {
		From int
		To   int
	}
	type args struct {
		input []byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []byte
	}{
		{
			name: "reverse 1 to 4",
			fields: fields{
				From: 1,
				To:   4,
			},
			want: []byte("aedcb"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := ReversePositions{
				From: tt.fields.From,
				To:   tt.fields.To,
			}
			data := []byte("abcde")
			r.Apply(data)
			if !reflect.DeepEqual(data, tt.want) {
				t.Errorf("Apply() = %s, want %s", data, tt.want)
			}
		})
	}
}

func TestRotateBasedOnPosition_Execute(t *testing.T) {
	type fields struct {
		Letter byte
	}
	type args struct {
		input []byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []byte
	}{
		{
			name: "rotate based on position",
			fields: fields{
				Letter: 'a',
			},
			want: []byte("eabcd"),
		},
		{
			name: "rotate based on position",
			fields: fields{
				Letter: 'e',
			},
			want: []byte("eabcd"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := RotateBasedOnPosition{
				Letter: tt.fields.Letter,
			}

			data := []byte("abcde")
			r.Apply(data)
			if !reflect.DeepEqual(data, tt.want) {
				t.Errorf("Apply() = %s, want %s", data, tt.want)
			}
		})
	}
}

func TestUndo(t *testing.T) {

	actions := map[string]Instruction{
		"RotateLeft":              RotateLeft{Steps: 1},
		"RotateRight":             RotateRight{Steps: 1},
		"RotateBasedOnPosition_a": RotateBasedOnPosition{Letter: 'a'},
		"RotateBasedOnPosition_b": RotateBasedOnPosition{Letter: 'b'},
		"RotateBasedOnPosition_c": RotateBasedOnPosition{Letter: 'c'},
		"RotateBasedOnPosition_d": RotateBasedOnPosition{Letter: 'd'},
		"RotateBasedOnPosition_e": RotateBasedOnPosition{Letter: 'e'},
		"SwapLetter":              SwapLetter{From: 'a', To: 'b'},
		"SwapPosition":            SwapPosition{From: 1, To: 4},
		"MovePosition":            MovePosition{From: 1, To: 4},
		"ReversePositions":        ReversePositions{From: 1, To: 4},
	}

	for name, action := range actions {
		t.Run(name, func(t *testing.T) {
			data := []byte("abcde")
			t.Run("Apply", func(t *testing.T) {
				action.Apply(data)
				if string(data) == "abcde" {
					t.Errorf("%s.Apply() = abcde, want something else", name)
				}
			})

			t.Run("Undo", func(t *testing.T) {
				action.Undo(data)
				if string(data) != "abcde" {
					t.Errorf("%s.Undo() = %s, want abcde", name, data)
				}
			})
		})
	}
}
