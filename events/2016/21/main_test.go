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
		args   args
		want   []byte
	}{
		{
			name: "rotateLeft 1",
			fields: fields{
				Steps: 1,
			},
			args: args{
				input: []byte("abcde"),
			},
			want: []byte("bcdea"),
		},
		{
			name: "rotateLeft 2",
			fields: fields{
				Steps: 2,
			},
			args: args{
				input: []byte("abcde"),
			},
			want: []byte("cdeab"),
		},
		{
			name: "rotateLeft -1",
			fields: fields{
				Steps: -1,
			},
			args: args{
				input: []byte("abcde"),
			},
			want: []byte("eabcd"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := RotateLeft{
				Steps: tt.fields.Steps,
			}
			if got := r.Execute(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Execute() = %s, want %s", got, tt.want)
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := MovePosition{
				From: tt.fields.From,
				To:   tt.fields.To,
			}
			if got := m.Execute(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Execute() = %v, want %v", got, tt.want)
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
			args: args{
				input: []byte("abcde"),
			},
			want: []byte("acdeb"),
		},
		{
			name: "move 4 to 1",
			fields: fields{
				From: 4,
				To:   1,
			},
			args: args{
				input: []byte("abcde"),
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
			if got := m.Execute(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Execute() = %v, want %v", got, tt.want)
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
			args: args{
				input: []byte("abcde"),
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
			if got := r.Execute(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Execute() = %s, want %s", got, tt.want)
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
				Letter: 'b',
			},
			args: args{
				input: []byte("abcdef"),
			},
			want: []byte("efabcd"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := RotateBasedOnPosition{
				Letter: tt.fields.Letter,
			}
			if got := r.Execute(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Execute() = %s, want %s", got, tt.want)
			}
		})
	}
}
