package aoc

func popcnt(x uint8) int {
	// https://stackoverflow.com/questions/14009765/fastest-way-to-count-bits
	x = (x & 0x55) + ((x >> 1) & 0x55)
	x = (x & 0x33) + ((x >> 2) & 0x33)
	x = (x & 0x0f) + ((x >> 4) & 0x0f)
	return int(x)
}

func CountBits(n uint) int {
	var total int
	switch {
	case n >= 0xFFFFFFFF:
		total += popcnt(uint8(n>>32)) + popcnt(uint8(n>>40)) + popcnt(uint8(n>>48)) + popcnt(uint8(n>>56))
		fallthrough
	case n >= 0xFFFF:
		total += popcnt(uint8(n>>16)) + popcnt(uint8(n>>24))
		fallthrough
	case n >= 0xFF:
		total += popcnt(uint8(n >> 8))
		fallthrough
	case n >= 0:
		total += popcnt(uint8(n))
	}
	return total
}
