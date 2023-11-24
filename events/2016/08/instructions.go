package main

import "fmt"

type Instruction interface {
	fmt.Stringer
	Apply(d *Display)
}

type Rect struct {
	width, height int
}

func RectFromString(s string) (Rect, error) {
	var r Rect
	_, err := fmt.Sscanf(s, "rect %dx%d", &r.width, &r.height)
	return r, err
}

func (r Rect) String() string {
	return fmt.Sprintf("rect %dx%d", r.width, r.height)
}

func (r Rect) Apply(d *Display) {
	d.Rect(r.width, r.height)
}

type ShiftRow struct {
	y, n int
}

func (s ShiftRow) String() string {
	return fmt.Sprintf("rotate row y=%d by %d", s.y, s.n)
}

func (s ShiftRow) Apply(d *Display) {
	d.ShiftRow(s.y, s.n)
}

type ShiftColumn struct {
	x, n int
}

func (s ShiftColumn) String() string {
	return fmt.Sprintf("rotate column x=%d by %d", s.x, s.n)
}

func (s ShiftColumn) Apply(d *Display) {
	d.ShiftColumn(s.x, s.n)
}
