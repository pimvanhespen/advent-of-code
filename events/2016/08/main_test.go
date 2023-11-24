package main

import (
	"fmt"
	"testing"
)

func TestStuff(t *testing.T) {

	d := NewDisplay(7, 3)
	d.Rect(3, 2)
	fmt.Println(d)
	d.ShiftRow(0, 1)
	fmt.Println(d)
	d.ShiftColumn(1, 2)
	fmt.Println(d)
}
