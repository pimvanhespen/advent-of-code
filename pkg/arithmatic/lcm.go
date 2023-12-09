package arithmatic

// LCM returns the least common multiple of the given numbers.
func LCM(nums ...int) int {
	lcm := nums[0]
	for i := 1; i < len(nums); i++ {
		lcm = lcm * nums[i] / GCD(lcm, nums[i])
	}
	return lcm
}

// GCD returns the greatest common divisor of a and b.
func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
