package aoc

type Numeric interface {
	int8 | int16 | int32 | int64 | int | uint8 | uint16 | uint32 | uint64 | uint | float32 | float64
}

func Sum[V Numeric](values ...V) V {
	total := V(0)
	for _, v := range values {
		total += v
	}
	return total
}
