package main

import (
	"fmt"
	"strings"
)

type Display struct {
	segments      [][]bool
	width, height int
}

func NewDisplay(width, height int) *Display {
	rows := make([][]bool, height)
	for i := range rows {
		rows[i] = make([]bool, width)
	}

	return &Display{
		segments: rows,
		width:    width,
		height:   height,
	}
}

func (d *Display) String() string {
	var sb strings.Builder
	for _, row := range d.segments {
		for _, col := range row {
			if col {
				sb.WriteByte('#')
			} else {
				sb.WriteByte(' ')
			}
		}
		sb.WriteByte('\n')
	}
	return sb.String()
}

func (d *Display) Rect(width, height int) {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			d.segments[y][x] = true
		}
	}
}

// ShiftRow shifts the yth row to the right by n pixels
func (d *Display) ShiftRow(y int, n int) {
	if y < 0 || y >= d.height {
		panic(fmt.Sprintf("y out of bounds: %d", y))
	}

	// normalize
	n = (n + d.width) % d.width
	if n == 0 {
		return
	}

	offset := d.width - n

	d.segments[y] = append(d.segments[y][offset:], d.segments[y][:offset]...) // fyi. not working. offset is off
}

// ShiftColumn shifts the xth column down by n pixels
func (d *Display) ShiftColumn(x int, n int) {
	if x < 0 || x >= d.width {
		panic(fmt.Sprintf("x out of bounds: %d", x))
	}

	n = (n + d.height) % d.height
	if n == 0 {
		return
	}

	newCol := make([]bool, d.height)
	for y := 0; y < d.height; y++ {
		newCol[y] = d.segments[y][x]
	}

	offset := d.height - n
	newCol = append(newCol[offset:], newCol[:offset]...)

	for y := 0; y < d.height; y++ {
		d.segments[y][x] = newCol[y]
	}
}

func (d *Display) Count() int {
	var count int
	for _, row := range d.segments {
		for _, col := range row {
			if col {
				count++
			}
		}
	}
	return count
}
