package heap

import "github.com/pimvanhespen/advent-of-code/pkg/aoc"

// Node is a node in the heap.
// P is the priority type.
// T is the value type.
type Node[P aoc.Numeric, T any] struct {
	prio  P
	value T
}

// Min is a min-heap implementation of the heap data structure.
type Min[P aoc.Numeric, T any] struct {
	heap []*Node[P, T]
}

// NewMin returns a new Min heap.
func NewMin[P aoc.Numeric, T any]() *Min[P, T] {
	return &Min[P, T]{
		heap: make([]*Node[P, T], 0),
	}
}

// Push pushes a value onto the heap.
func (m *Min[P, T]) Push(t T, p P) {
	m.heap = append(m.heap, &Node[P, T]{value: t, prio: p})
	m.up(len(m.heap) - 1)
}

// Pop pops the minimum value from the heap.
func (m *Min[P, T]) Pop() T {
	if len(m.heap) == 0 {
		var zero T
		return zero
	}
	v := m.heap[0]
	m.heap[0] = m.heap[len(m.heap)-1]
	m.heap = m.heap[:len(m.heap)-1]
	m.down(0)
	return v.value
}

// Peek returns the minimum value from the heap without removing it.
func (m *Min[P, T]) Peek() T {
	if len(m.heap) == 0 {
		var zero T
		return zero
	}
	return m.heap[0].value
}

// Len returns the number of elements in the heap.
func (m *Min[P, T]) Len() int {
	return len(m.heap)
}

// up moves the element at index i up the heap.
func (m *Min[P, T]) up(i int) {
	for {
		p := (i - 1) / 2
		if p == i || m.heap[i].prio >= m.heap[p].prio {
			break
		}
		m.heap[i], m.heap[p] = m.heap[p], m.heap[i]
		i = p
	}
}

// down moves the element at index i down the heap.
func (m *Min[P, T]) down(i int) {
	for {
		c := 2*i + 1
		if c >= len(m.heap) {
			break
		}
		if c+1 < len(m.heap) && m.heap[c+1].prio < m.heap[c].prio {
			c++
		}
		if m.heap[i].prio <= m.heap[c].prio {
			break
		}
		m.heap[i], m.heap[c] = m.heap[c], m.heap[i]
		i = c
	}
}

func (m *Min[P, T]) Empty() bool {
	return len(m.heap) == 0
}
