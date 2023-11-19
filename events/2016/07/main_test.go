package main

import (
	"bytes"
	"testing"
)

func TestIP_SupportsTLS(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{
			input: "abba[mnop]qrst",
			want:  true,
		},
		{
			input: "abcd[bddb]xyyx",
			want:  false,
		},
		{
			input: "aaaa[qwer]tyui",
			want:  false,
		},
		{
			input: "ioxxoj[asdfgh]zxcvbn",
			want:  true,
		},
		{
			input: "xdsqxnovprgovwzkus[fmadbfsbqwzzrzrgdg]aeqornszgvbizdm",
			want:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			ip, err := parseIP(tt.input)
			if err != nil {
				t.Fatal(err)
			}

			t.Log(string(bytes.Join(ip.Hyper, []byte(", "))))
			t.Log(string(bytes.Join(ip.Super, []byte(", "))))

			got := ip.SupportsTLS()
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIP_SSL(t *testing.T) {

	tests := []struct {
		input string
		want  bool
	}{
		{
			input: "aba[bab]xyz",
			want:  true,
		},
		{
			input: "xyx[xyx]xyx",
			want:  false,
		},
		{
			input: "aaa[kek]eke",
			want:  true,
		},
		{
			input: "zazbz[bzb]cdb",
			want:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			ip, err := parseIP(tt.input)
			if err != nil {
				t.Fatal(err)
			}

			t.Log("hyper:", string(bytes.Join(ip.Hyper, []byte(", "))))
			t.Log("super:", string(bytes.Join(ip.Super, []byte(", "))))

			if got := ip.SSL(); got != tt.want {
				t.Errorf("SSL() = %v, want %v", got, tt.want)
			}
		})
	}
}
